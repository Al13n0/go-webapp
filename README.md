# go-webapp
example web app in go written using gomux
https://github.com/gorilla/mux

## Setup

Save Db password inside .env file

```
go init go-webapp
go run main.go
```

## Endpoints

```
POST /books Add new book

GET /books return all books

GET /books/{id} return specific book based on provided id

PUT /books update a book

DELET /books delete a book
```
Use postman collection provided in the repo for testing the endpoint


