

# Routine Management MS

This is a Go-based CRUD MS application for managing exercises and routines. It uses MongoDB as the database and the `chi` router for routing. The application exposes RESTful API endpoints for managing exercises and routines.

---

## Docker

To build and run the application using Docker:

```bash
sudo docker build -t routineManagement_ms .
sudo docker run --rm -p 3000:3000 --env-file .env routineManagement_ms
```

---



## Installation and Setup without docker

### Prerequisites
- Go 1.20 or later installed on your system
- MongoDB instance running (Atlas free tier in test environment of this code)

### Clone the Repository
```bash
git clone https://github.com/your-username/exercise_api.git
cd exercise_api
```

### Install Dependencies
Ensure all required Go modules are installed:
```bash
go mod tidy
```

### Configure the Database
Update the `DB` connection in the application code to point to your MongoDB instance. Look for the database configuration (e.g., `DB.Client.Database("exercise_app")`) and ensure it aligns with your setup, in our home we are using a .env file like this one:
```bash
MONGO_URI = connection_string_given_by_atlas
```

### Run the Application
Start the server:
```bash
go run main.go
```
By default, the application will run on `http://localhost:8080`.

---

## API Endpoints

### Base URL
`http://localhost:8080`

### Exercise Routes
| Method | Endpoint       | Description                |
|--------|----------------|----------------------------|
| POST   | `/exercises/`  | Create a new exercise      |
| GET    | `/exercises/`  | List all exercises         |
| GET    | `/exercises/{id}` | Get a specific exercise by ID |
| PUT    | `/exercises/{id}` | Update an exercise by ID   |
| DELETE | `/exercises/{id}` | Delete an exercise by ID   |

#### Example Requests
1. **Create an Exercise**
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{
       "exercise_name": "Push Ups",
       "muscular_group": [
           { "muscle_id": 1, "muscle_name": "Chest" },
           { "muscle_id": 2, "muscle_name": "Triceps" }
       ]
   }' http://localhost:8080/exercises/
   ```

2. **List All Exercises**
   ```bash
   curl -X GET http://localhost:8080/exercises/
   ```

3. **Get an Exercise by ID**
   ```bash
   curl -X GET http://localhost:8080/exercises/{id}
   ```

4. **Update an Exercise**
   ```bash
   curl -X PUT -H "Content-Type: application/json" -d '{
       "exercise_name": "Pull Ups",
       "muscular_group": [
           { "muscle_id": 3, "muscle_name": "Back" }
       ]
   }' http://localhost:8080/exercises/{id}
   ```

5. **Delete an Exercise**
   ```bash
   curl -X DELETE http://localhost:8080/exercises/{id}
   ```

### Routine Routes
| Method | Endpoint       | Description                |
|--------|----------------|----------------------------|
| POST   | `/routines/`   | Create a new routine       |
| GET    | `/routines/`   | List all routines          |
| GET    | `/routines/{id}` | Get a specific routine by ID |
| PUT    | `/routines/{id}` | Update a routine by ID    |
| DELETE | `/routines/{id}` | Delete a routine by ID    |

---

## Updating the Application
1. Pull the latest changes from the repository:
   ```bash
   git pull origin main
   ```
2. Update dependencies if needed:
   ```bash
   go mod tidy
   ```
3. Restart the application:
   ```bash
   go run main.go
   ```

