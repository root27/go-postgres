# Go server with PostgreSQL

This is a simple Go server with a PostgreSQL database.


## Running the server

```bash

go run .

```

## Running the PostgreSQL database

```bash

docker run --name postgres  -e POSTGRES_USER=test -e POSTGRES_PASSWORD=123 -p 5432:5432 -d postgres

```

### API

#### Create a user

```bash
 

curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe", "age": 12}' http://localhost:3000/create

```

#### Get all users

```bash

curl -X GET http://localhost:3000/getAll

```

#### Get user by id

```bash

curl -X GET http://localhost:3000/get/:id


```

#### Update user by id

```bash

curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe", "age": 12}' http://localhost:3000/update/:id

```

#### Delete user by id

```bash

curl -X DELETE http://localhost:3000/delete/:id

```


