# Dating App Backend for upwork

This is a backend service for a dating app implemented in Go using the Gin framework and GORM for ORM. The service includes endpoints for user management and matchmaking, with data stored in a PostgreSQL database.

## Table of Contents

- [Project Overview](#project-overview)
- [Requirements](#requirements)
- [Setup and Installation](#setup-and-installation)
- [API Endpoints](#api-endpoints)
  - [Create User](#create-user)
  - [Delete User](#delete-user)
  - [Matchmaking Recommendations](#matchmaking-recommendations)
- [Database Migrations](#database-migrations)
- [Testing](#testing)
- [Docker](#docker)
- [Makefile](#makefile)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

This backend service handles user operations and matchmaking functionality for a dating app. The service includes the following features:

- Create and manage user profiles.
- Delete user profiles.
- Recommend potential matches based on user preferences and interests.

## Requirements

- Go 1.18 or higher
- PostgreSQL
- Docker (for containerization)
- `migrate` tool for database migrations
- `make` tool for facilitating the development and testing process

## Setup and Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/dating-app-backend.git
   cd dating-app-backend
   ```

2. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

3. **Setup PostgreSQL Database:**

   Ensure you have PostgreSQL installed and running. Create a database for the application.

4. **Configure Environment Variables:**

   Create a `.env` file in the root directory with the following variables:

   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/yourdatabase
   ```

5. **Run Database Migrations:**

   ```bash
   migrate -path ./migrations -database "$DATABASE_URL" up
   ```

6. **Start the Application:**

   ```bash
   go run main.go
   ```

## API Endpoints

### Create User

- **Endpoint:** `POST /api/create-user`
- **Description:** Creates a new user profile.
- **Request Body:**

  ```json
  {
    "name": "John Doe",
    "age": 30,
    "gender": "male",
    "location": "37.7749,-122.4194",
    "interests": ["hiking", "cooking", "music"],
    "preferences": {
      "min_age": 25,
      "max_age": 35,
      "preferred_gender": "female",
      "max_distance": 50
    },
    "last_active": "2024-09-12T12:00:00Z"
  }
  ```

- **Response:**

  ```json
  {
    "message": "User created successfully",
    "user_id": "some-uuid"
  }
  ```

### Delete User

- **Endpoint:** `DELETE /api/delete/:user_id`
- **Description:** Deletes a user profile based on the user ID.
- **Parameters:**
  - `user_id` (path parameter): The UUID of the user to be deleted.
- **Response:**

  ```json
  {
    "message": "User deleted successfully"
  }
  ```

### Matchmaking Recommendations

- **Endpoint:** `GET /api/match/recommendations/:user_id`
- **Description:** Provides a list of potential matches for a user based on preferences, mutual interests, and activity status.
- **Parameters:**
  - `user_id` (path parameter): The UUID of the current user.
- **Response:**

  ```json
  [
    {
      "user_id": "another-uuid",
      "name": "Jane Smith",
      "age": 28,
      "gender": "female",
      "location": "37.7749,-122.4194",
      "interests": ["hiking", "music"],
      "score": 85
    }
  ]
  ```

## Database Migrations

Migrations are managed using the `migrate` tool. Place migration files in the `migrations` directory. Apply migrations with:

```bash
migrate -path ./migrations -database "$DATABASE_URL" up
```

## Testing

To test the application, use tools like Postman or curl. Ensure your server is running and use the provided endpoints to perform various CRUD operations.

## Docker

To run the application in a Docker container, build and run the Docker image:

1. **Build the Docker Image:**

   ```bash
   docker build -t dating-app-backend .
   ```

2. **Run the Docker Container:**

   ```bash
   docker run -p 8080:8080 --env-file .env dating-app-backend
   ```

## Makefile

The Makefile provides convenient commands for managing the application. Here are some common commands:

- **Build Docker Image:**

  ```bash
  make build
  ```

- **Run Migrations:**

  ```bash
  make migrate
  ```

- **Start Application:**

  ```bash
  make run
  ```

## Contributing

Feel free to submit issues or pull requests. Please ensure that your code follows the project's coding standards and includes relevant tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
