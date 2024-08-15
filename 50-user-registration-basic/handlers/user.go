// handlers/user.go
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"

	"user-registration-basic/models"

	"github.com/gorilla/mux"
)

var (
	users  = make(map[int]models.User)
	nextID = 1
	mu     sync.Mutex
)

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// To have 5000 Query per second (QPS), we can implement worker pool using channels and goroutines
// START >>>> Worker pool
type request struct {
	id  int
	res chan<- models.User
}

var (
	requests = make(chan request, 5000)
)

func init() {
	for i := 0; i < 100; i++ {
		// create a new goroutine
		go worker()
	}
}

func worker() {
	for req := range requests {
		mu.Lock()
		user, exists := users[req.id]
		mu.Unlock()
		if exists {
			req.res <- user
		} else {
			req.res <- models.User{}
		}
	}
}

// END <<<<< Worker pool

// validateUser validates the given user object and checks if all the required fields are present.
// It returns an error if any of the required fields are missing or empty.
func validateUser(user models.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Gender == "" || user.Age <= 0 || user.Address == "" || user.ContactNo == "" || user.ContactEmail == "" || user.Photo == "" {
		return errors.New("all fields are required")
	}
	return nil
}

/* CreateUser Method:
* CreateUser creates a new user based on the provided request body and adds it to the users map.
* It expects a JSON-encoded user model in the request body and returns the created user as the response.
* If there is any error in decoding the request body or adding the user to the map, it returns a bad request error.
 */
func CreateUser(w http.ResponseWriter, r *http.Request) {

	setCORSHeaders(w)
	if r.Method == "OPTIONS" {
		return
	}

	// 1. Create new user model
	var user models.User

	// 2. check if any error in requested user model body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Lock the mutex >> set the user model to global users map variable >> Unlock the mutex
	mu.Lock()
	user.ID = nextID
	nextID++
	users[user.ID] = user
	mu.Unlock()

	// 4. send the http status to response and send the actual http response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser retrieves a user by their ID and writes the user information as JSON to the response writer.
func GetUser(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	// 1. Get the id as integer from http request >> and return the error if the id is invalid
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// // 2. Lock the mutex >> get the user  from the globally defined users map array >> unlock the mutex
	// mu.Lock()
	// user, exists := users[id]
	// mu.Unlock()

	// // 3. Throw the error if any error comes while fetching user form users map
	// if !exists {
	// 	http.Error(w, "User not found", http.StatusNotFound)
	// 	return
	// }

	res := make(chan models.User)
	requests <- request{id: id, res: res}
	user := <-res

	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// 4. return the encoded user to response writer
	json.NewEncoder(w).Encode(user)
}

/* UpdateUser updates the user with the specified ID.
* It expects a PUT request with a JSON payload containing the updated user information.
* The user ID is extracted from the request URL parameters.
* If the user ID is invalid or the request body is not a valid JSON, it returns an appropriate error response.
* If the user with the specified ID does not exist, it returns a "User not found" error response.
* Otherwise, it updates the user information and returns the updated user as a JSON response.
 */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	// 1. Get the user id to be updated from request variables
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// 2. Create new updatedUser variable of type User model
	var updatedUser models.User

	// 3. return the http error while updating any invalid user details
	// json.NewDecoder(r.Body) creates a new Decoder object that reads from the r.Body which is an io.Reader interface representing the request body.
	//  .Decode(&updatedUser) decodes the JSON data from the Decoder into the updatedUser variable. The & symbol is used to pass the address of the updatedUser variable so that the Decode method can modify its value directly.
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4. locks the mutex for other go routines to not to use the resources >> Fetches the User model from the users map
	mu.Lock()
	user, exists := users[id]
	if !exists {
		mu.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// 5. fetched user id is getting updated to newly created updatedUser model >> users map is appended/updated with new user Id >> unlock of mutex
	updatedUser.ID = user.ID
	users[id] = updatedUser
	mu.Unlock()

	json.NewEncoder(w).Encode(updatedUser)
}

/* DeleteUser deletes a user with the given ID from the users map.
* It takes an http.ResponseWriter and an http.Request as parameters.
* If the user ID is invalid, it returns a Bad Request error.
* If the user is found and deleted successfully, it returns a No Content status.
* If the user is not found, it returns a Not Found error.
 */
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	_, exists := users[id]
	if exists {
		delete(users, id)
	}
	mu.Unlock()

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
