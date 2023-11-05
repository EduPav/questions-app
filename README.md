# QuestionsApp
## Author:
Eduardo Martín Pavéz Fabriani
https://www.linkedin.com/in/eduardo-pavez/

## Description

This Go application provides a platform where users can discover and share insightful questions to ask when they meet someone new. Its goal is to enhance dialogues by inspiring users with a wide range of questions.

In its current initial state, it features a RESTful API for managing users and questions, with full CRUD operations and other specific functionalities, such as random question retrieval.

## Prerequisites

- Go 1.21.3
- PostgreSQL or Docker
- Any REST client (e.g., Postman, Thunder Client) to test the endpoints.


## Running the Application

To run the server, go to the root directory and run:

```sh
go run .
```

This will start the server on `http://localhost:8080`.

## API Endpoints

The application exposes several API endpoints:

- `/api/users`: POST to create a new user, GET to retrieve all users.
    -  JSON body: `{"username": "username", "password": "password"}`
- `/api/users/{id}`: GET to retrieve a specific user, PATCH to update, DELETE to remove.
- `/api/questions`: POST to create a new question, GET to retrieve all questions.
- `/api/questions/{id}`: GET to retrieve a specific question, PATCH to update, DELETE to remove.
    - JSON body: `{"description": "description", "rating": "rating", "creator_id": "creator_id"}`
- `/api/rquestions/{id}`: GET to retrieve a random question created by the user, enhancing the experience of the app. 

## Database Configuration

Update the `connection.go` DSN variable with your PostgreSQL credentials.

## Instructions to run postgresql from docker and create the database
I assume you have docker installed on your machine, and you want to create an isolated container for the database.
1. If you don't have PostgreSQL installed jump to step 2. As we are using port 5432 for PostgreSQL, we need to stop its service running on the host machine. To do so, follow this instructions on windows
    - Press Winkey+R
    - Type services.msc
    - Search Postgres service
    - Right click and stop the service
2. Run the following commands to download the image and start the container
- ```docker pull postgres```
- ```docker run --name qapp-postgres -e POSTGRES_USER=eduyio -e POSTGRES_PASSWORD=admin -p 5432:5432 -d postgres``` 
- ```Docker exec -it qapp-postgres bash```
- ```psql -U eduyio --password```
- Type your password and Enter. For my example it's ```admin```
- ```CREATE DATABASE qappdb;```




## Models

- `User`: Represents a user with a username, password, and associated questions.
- `Question`: Represents a question with a description, rating, and creator.




## Future Development
- Option to explore other user's questions.
- Front-end interface to interact with the API.
- User authentication and authorization.
