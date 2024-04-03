# NBA Web-App

A web-app of a trivia game involving nba players

## Installing and Running

### Using Docker

To install the app, simply cd into the project root and run `docker compose build`.

Then, to run the app, run `docker compose up`.

The app will be running on localhost, port 3000.

### Manual Installation

Dependencies:

- postgreSQL
- Go 1.22.1
- Node 18+

#### Setting up the Database

You will need to have a database setup with a table following the schema of `database/schema.sql`.

Take note of the host, port, user, password and database name to setup the backend.

#### Setting up the Backend

Now, cd into backend/go-api and run `go build -o path/to/executable/main`, where `path/to/executable` is your desired path.

Then, create a .env file in the same directory as your executable with the following format:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=nbaappdb
BACKEND_PORT=9000
```

Change the values to whatever fits your database, and take note of the backend port you chose.

#### Setting up the Frontend

Afterwards, cd into frontend/react-app and run the following commands:

```bash
npm install
npm run build
```

In the same directory, create a .env.local file with the following format:

```bash
BACKEND_HOST=localhost
BACKEND_PORT=9000
```

Choose the same port you setup earlier for your backend.

#### Running the app

Firstly, make sure your database is up and running.

Then, start up the backend by running the executable you setup earlier.

Finally, cd once again into frontend/react-app and run `npm start -- --port 3000`. You can change the port to whatever you want.

The app will, by default, be running on localhost to whatever port you set.

## TO-DO

### Backend (Go)

- Protobuffs?
- Store user score?

### Frontend (React)

- Improve visuals
- Crop image to hide jerseys
