package worker

import (
	"log"
	"time"
	"user-registration/internal/db"
	"user-registration/pkg/models"
)

type ReadRequest struct {
	ID       uint
	Response chan ReadResponse
}

type ReadResponse struct {
	User  *models.User
	Error error
}

var readRequests = make(chan ReadRequest, 10000)

// func StartReadWorkerPool(numWorkers int, getUserFunc func(uint) (*models.User, error)) {
// 	for i := 0; i < numWorkers; i++ {
// 		go readWorker(getUserFunc)
// 	}
// 	log.Printf("Started %d read workers", numWorkers)
// }

func StartReadWorkerPool(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go readWorker()
	}
	log.Printf("Started %d read workers", numWorkers)
}

// func readWorker(getUserFunc func(uint) (*models.User, error)) {
// 	for req := range readRequests {
// 		log.Printf("Worker received request for ID: %d", req.ID)
// 		user, err := getUserFunc(req.ID)
// 		// req.Response <- ReadResponse{User: user, Error: err}
// 		select {
// 		case req.Response <- ReadResponse{User: user, Error: err}:
// 			log.Printf("Worker processed request for ID: %d", req.ID)
// 		case <-time.After(5 * time.Second):
// 			log.Printf("Worker timed out processing request for ID: %d", req.ID)
// 		}
// 	}
// }

func readWorker() {
	for req := range readRequests {
		log.Printf("Worker received request for ID: %d", req.ID)
		var user models.User
		err := db.DB.First(&user, req.ID).Error
		if err != nil {
			log.Printf("Error fetching user with ID: %d, error: %v", req.ID, err)
		} else {
			log.Printf("Fetched user with ID: %d, user: %+v", req.ID, user)
		}
		select {
		case req.Response <- ReadResponse{User: &user, Error: err}:
			log.Printf("Worker processed request for ID: %d", req.ID)
		case <-time.After(5 * time.Second):
			log.Printf("Worker timed out processing request for ID: %d", req.ID)
		}
	}
}

func QueueReadRequest(id uint) ReadResponse {
	log.Printf("Queueing read request for ID: %d", id)
	response := make(chan ReadResponse)
	readRequests <- ReadRequest{ID: id, Response: response}
	// return <-response
	select {
	case res := <-response:
		log.Printf("Received response for ID: %d, Response: %v ", id, res)
		return res
	case <-time.After(5 * time.Second):
		log.Printf("Timed out waiting for response for ID: %d", id)
		return ReadResponse{User: nil, Error: nil}
	}
}
