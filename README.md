# Dating App Backend for Upwork

This is a backend service for a dating app implemented in Go using the Gin framework and GORM for ORM for a test round for an Upwork job. The service includes endpoints for user management and matchmaking, with data stored in a PostgreSQL database. Matchmaking is optimized with mutual interests filtering and location-based ranking.

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
  - [Test Data](#test-data)
  - [Using `seed.py` for Testing](#using-seedpy-for-testing)
- [Docker](#docker)
- [Makefile](#makefile)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

This backend service handles user operations and matchmaking functionality for a dating app. The service includes the following features:

- Create and manage user profiles.
- Delete user profiles.
- Recommend potential matches based on user preferences, location, and mutual interests, with performance optimizations for large datasets.

## Requirements

- `Go` 1.18 or higher
- `PostgreSQL` with PostGIS extension (for location-based queries)
- `Docker` (for containerization)
- `migrate` tool for database migrations
- `make` tool for facilitating the development and testing process
- `python` (optional for testing and seeding the user database)

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

   Rename the `.ev.example` file to `.env` and replace the variables with those specific to your local machine.

5. **Run Database Migrations:**

   Update the content of the Makefile to accommodate the local setup of your machine, then run the following commands:

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

   You can use one of the following options to run the app:

   1. Dry run

      ```bash
      make run
      ```

   2. Using `air` for hot reloading

      ```bash
      make air
      ```

   3. Using Docker Compose

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
    "password": "the_most_secure_password_on_earth",
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

### Sign In

- **Endpoint:** `POST /api/v1/sign-in`
- **Description:** Login as a user.
- **Request Body:**

  ```json
  {
    "username": "John Doe",
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

- **Endpoint:** `DELETE /users/delete/:user_id`
- **Description:** Deletes a user profile based on the user ID.
- **Parameters:**
  - `user_id` (path parameter): The UUID of the user to be deleted.
  - `Authorization` (header): The JWT token provided during creation.
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
      "name": "Margaret Lua",
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

Migrations are managed using the `migrate` tool. the migration files are in the `migrations` directory. Apply migrations with:

```bash
migrate -path ./migrations -database "$DATABASE_URL" up
```

## Testing

To test the application, use the `Postman collection` in the `docs` folder. along side the `test` folder contents which are explained below.

### Test Data

You can use the `test.json` file to seed your database. The data includes a variety of users with different interests and preferences for thorough testing.

### Using `seed.py` for Testing

To quickly populate the database with test data, you can use the `seed.py` script. This script reads from a JSON file and inserts the data into your database.

1. **Ensure Python is Installed:**

   Make sure you have Python 3.x installed on your system.

2. **Install Required Python Packages:**

   ```bash
   pip install psycopg2-binary requests
   ```

3. **Prepare Test Data File:**

   Save the test data provided above into a file named `test_data.json`.

4. **Run `seed.py`:**

   ```bash
   python seed.py --file test_data.json
   ```

   The script will read the data from `test.json` and insert it into the PostgreSQL database. Make sure to update the database connection details in `seed.py` if needed.

## Docker

To run the application in Docker containers using Docker Compose, follow these steps:

1. **Build the Docker Image and Start Containers:**

   To build the Docker image and start the application with Docker Compose, run:

   ```bash
   make start
   ```

   This command will:
   - Build the Docker image using `docker-compose`.
   - Apply database migrations using `migrate-up`.
   - Start the application and related services defined in `docker-compose.yml`.

2. **Build the Docker Image Only:**

   If you want to build the Docker image without starting the containers, use:

   ```bash
   make docker-compose-build
   ```

   This will build the Docker image without caching.

3. **Start Docker Containers:**

   To start the Docker containers defined in `docker-compose.yml` without rebuilding the image, run:

   ```bash
   make docker-compose-up
   ```

4. **Stop Docker Containers:**

   To stop and remove the running Docker containers, use:

   ```bash
   make docker-compose-down
   ```

## Makefile

The Makefile provides convenient commands for managing the application. You can update it to fit your needs.

## Contributing

Feel free to submit issues or pull requests. Please ensure that your code follows the project's coding standards and includes relevant tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
