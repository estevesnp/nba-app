# NBA Web-App

A web-app of a trivia game involving nba players

## Dependencies

- Go 1.22.1
- PostgreSQL
- NPM

## Install

First, make sure you have PostgreSQL running on port 5432.
Then, create a database with the name "nbaappdb" and create a table inside it with the schema from `backend/schema.sql`.
Afterwards, cd into `backend/go-api` and run `go build -o ../../backend-api.exe`.
You also need to create a `.env` file in the same directory as the executable, with a user and password credentials for your database, like this:

```bash
USER=username
PASSWORD=password
```

Now, cd into `frontend/react-app` and run `npm install` and `npm run build`

## Run the app

To run the app, cd into the project root and run `backend-api.exe`. It runs on port 8080.
Afterwards, cd into frontend/react-app and run `npm run start -- --port 8081`. You can change the port to whatever you want.

Afterwards, you can access the app on your browser by going to localhost:8081 (or whatever port you set).

## TO-DO

### Backend (Go)

- Protobuffs?
- Store user score?

### Frontend (React)

- Improve visuals

### Misc

- Create Docker file
