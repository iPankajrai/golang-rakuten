# User Registration Application

This is a User Registration application built with Go, Gin, GORM, and MySQL. The application is designed to handle CRUD operations for user data and is capable of handling 5000 QPS for read operations using a worker pool. The project follows the standard Go project layout and is Dockerized for easy deployment.

## Project Structure

user-registration/
├── api/
│   └── v1/
│       └── user.go
├── cmd/
│   └── user-registration/
│       └── main.go
├── configs/
│   └── config.go
├── internal/
│   ├── app/
│   │   ├── controllers/
│   │   │   └── user_controller.go
│   │   ├── repositories/
│   │   │   └── user_repository.go
│   │   ├── services/
│   │   │   └── user_service.go
│   ├── db/
│   │   └── db.go
│   ├── worker/
│   │   └── worker.go
├── pkg/
│   └── models/
│       └── user.go
├── go.mod
├── go.sum
├── wait-for-it.sh
├── Dockerfile
└── docker-compose.yml
├── README.md


## Prerequisites

- Docker
- Docker Compose

## Setup

1. **Clone the repository**:

    ```sh
    git clone https://github.com/your-username/user-registration.git
    cd user-registration
    ```

2. **Set up environment variables**:

    Create a `.env` file in the root directory with the following content:

    ```env
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=user_registration
    DB_HOST=db
    DB_PORT=3306
    ```

3. **Build and run the application using Docker Compose**:

    ```sh
    docker-compose up --build
    ```

4. **Verify the application**:

    Open your browser and navigate to `http://localhost:8080`.

## API Endpoints

### Create a new user

- **URL**: `/users`
- **Method**: `POST`
- **Request Body**:

    ```json
    {
        "first_name": "Pankaj",
        "last_name": "Kumar",
        "gender": "Male",
        "age": 27,
        "address": "123 Main St",
        "contact_no": "1234567890",
        "contact_email": "pankaj.a.kumar@rakuten.com",
        "photo": "http://example.com/photo.jpg"
    }
    ```

- **Response**:

    ```json
    {
        "id": 1,
        "first_name": "Pankaj",
        "last_name": "Kumar",
        "gender": "Male",
        "age": 27,
        "address": "123 Main St",
        "contact_no": "1234567890",
        "contact_email": "pankaj.a.kumar@rakuten.com",
        "photo": "http://example.com/photo.jpg"
    }
    ```

### Get a user by ID

- **URL**: `/users/:id`
- **Method**: `GET`
- **Response**:

    ```json
    {
        "id": 1,
        "first_name": "Pankaj",
        "last_name": "Kumar",
        "gender": "Male",
        "age": 27,
        "address": "123 Main St",
        "contact_no": "1234567890",
        "contact_email": "pankaj.a.kumar@rakuten.com",
        "photo": "http://example.com/photo.jpg"
    }
    ```

### Update a user

- **URL**: `/users/:id`
- **Method**: `PUT`
- **Request Body**:

    ```json
    {
        "first_name": "Pankaj",
        "last_name": "Kumar",
        "gender": "Male",
        "age": 28,
        "address": "123 Main St",
        "contact_no": "1234567890",
        "contact_email": "pankaj.a.kumar@rakuten.com",
        "photo": "http://example.com/photo.jpg"
    }
    ```

- **Response**:

    ```json
    {
        "id": 1,
        "first_name": "Pankaj",
        "last_name": "Kumar",
        "gender": "Male",
        "age": 28,
        "address": "123 Main St",
        "contact_no": "1234567890",
        "contact_email": "pankaj.a.kumar@rakuten.com",
        "photo": "http://example.com/photo.jpg"
    }
    ```

### Delete a user

- **URL**: `/users/:id`
- **Method**: `DELETE`
- **Response**: `204 No Content`

### Get all users

- **URL**: `/users`
- **Method**: `GET`
- **Response**:

    ```json
    [
        {
            "id": 1,
            "first_name": "Pankaj",
            "last_name": "Kumar",
            "gender": "Male",
            "age": 27,
            "address": "123 Main St",
            "contact_no": "1234567890",
            "contact_email": "pankaj.a.kumar@rakuten.com",
            "photo": "http://example.com/photo.jpg"
        }
    ]
    ```