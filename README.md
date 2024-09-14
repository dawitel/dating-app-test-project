# Dating App Backend for Upwork

This is a backend service for a dating app implemented in Go using the Gin framework and GORM for ORM for a test round for upwork job. The service includes endpoints for user management and matchmaking, with data stored in a PostgreSQL database. Matchmaking is optimized with mutual interests filtering and location-based ranking.

## Table of Contents

- [Project Overview](#project-overview)
- [Requirements](#requirements)
- [Setup and Installation](#setup-and-installation)
- [API Endpoints](#api-endpoints)
  - [Sign Up](#sign-up)
  - [Sign In](#sign-in)
  - [Delete User](#delete-user)
  - [Matchmaking Recommendations](#matchmaking-recommendations)
- [Database Migrations](#database-migrations)
- [Testing](#testing)
- [Docker](#docker)
- [Makefile](#makefile)
- [License](#license)

## Project Overview

This backend service handles user operations and matchmaking functionality for a dating app. The service includes the following features:

- Create and manage user profiles.
- Delete user profiles.
- Recommend potential matches based on user preferences, location, and mutual interests, with performance optimizations for large datasets.

## Requirements

- Go 1.18 or higher
- PostgreSQL with PostGIS extension (for location-based queries)
- Docker (for containerization)
- `migrate` tool for database migrations
- `make` tool for facilitating the development and testing process

## Setup and Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/dawitel/dating-app-test-project.git
   cd dating-app-test-project
   ```

2. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

3. **Setup PostgreSQL Database:**

   Ensure you have PostgreSQL installed and running, with PostGIS enabled. Create a database for the application.

   ```sql
   CREATE DATABASE dating_app;
   CREATE EXTENSION postgis;
   ```

4. **Configure Environment Variables:**

   rename the .ev.example file to `.env` and replace them with the variables of your local machine

5. **Run Database Migrations:**

  update the content of the make file to accomodate the local setup of you machine then run the following commands

   ```bash
   make migrate-create
   ```

   ```bash
   make migrate-up
   ```

   ```bash
   make migrate-down  ## this is for down migration
   ```

6. **Start the Application:**

  you can use one of the followig options to ru the app
  
  1. Dry run
  
   ```bash
   make run
   ```
  
  2. using air for hot reloading
  
   ```bash
   make air
   ```
  
  3. using docker-compose
  
   ```bash
   make start
   ```

## API Endpoints

### Sign Up

- **Endpoint:** `POST /api/v1/sign-up`
- **Description:** Creates a new user profile.
- **Request Body:**

  ```json
  {
    "name": "John Doe",
    "passord":"the_most_secure_password_on_earth",
    "age": 30,
    "gender": "male",
    "location": {
      "latitude": 37.7749,
      "longitude": -122.4194
    },
    "interests": ["hiking", "cooking", "music", "coding"],
    "preferences": {
      "min_age": 25,
      "max_age": 35,
      "preferred_gender": "female",
      "max_distance": 50
    }
  }
  ```

- **Response:**

  ```json
  {
    "message": "User created successfully",
    "user_id": "some-uuid",
    "token": "jwt-token"
  }
  ```

## Sign In

- **Endpoint:** `POST /api/v1/sign-in`
- **Description:** Login as a user.
- **Request Body:**

  ```json
  {
    "name": "John Doe",
    "password": "your_password"
  }
  ```

- **Response:**

  ```json
  {
    "message": "you are Logged in",
    "token": "your.jwt.token"
  }

  ```

### Delete User

- **Endpoint:** `DELETE /api/v1/delete/:user_id`
- **Description:** Deletes a user profile based on the user ID.
- **Parameters:**
  - `user_id` (path parameter): The UUID of the user to be deleted.
  - `Authorization`(header): The jwt token provided durig creation.
- **Response:**

  ```json
  {
    "message": "User deleted successfully"
  }
  ```

### Matchmaking Recommendations

- **Endpoint:** `GET /api/v1/match/recommendations/:user_id`
- **Description:** Provides a list of potential matches for a user based on preferences, mutual interests, location proximity, and activity status.
- **Parameters:**
  - `user_id` (path parameter): The UUID of the current user.
- **Response:**

  ```json
  [
    {
      "user_id": "another-uuid",
      "name": "Marry Jane",
      "age": 28,
      "gender": "female",
      "location": {
        "latitude": 37.7749,
        "longitude": -122.4194
      },
      "interests": ["hiking", "music"],
      "score": 85
    },
    {
      "user_id": "another-uuid",
      "name": "Margaret lua",
      "age": 28,
      "gender": "female",
      "location": {
        "latitude": 37.7749,
        "longitude": -122.4194
      },
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

Make sure your migrations handle fields for user location, preferences, interests, and timestamps.

- **Indexes:**
  - Ensure `GIN` indexing on `interests` for faster mutual interest querying:

    ```sql
    CREATE INDEX idx_users_interests ON users USING gin (interests);
    ```

  - Create `GIST` index for location-based queries using PostGIS:

    ```sql
    CREATE INDEX idx_users_location ON users USING gist (ST_MakePoint(longitude, latitude));
    ```

## Testing

To test the application, use tools like Postman or curl. Ensure your server is running and use the provided endpoints to perform various CRUD operations.

## Docker

To run the application in a Docker container, build and run the Docker image:

1. **Build the Docker Image:**

   ```bash
   make build 
   ```

2. **Run the Docker Container:**

   ```bash
   docker run -p 8080:8080 --env-file .env dating-app-backend
   ```

## Makefile

The Makefile provides convenient commands for managing the application. you can update it to make it go for your needs

## Contributing

Feel free to submit issues or pull requests. Please ensure that your code follows the project's coding standards and includes relevant tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
