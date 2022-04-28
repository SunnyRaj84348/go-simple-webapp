# go-simple-webapp
### A simple web application written using HTML, CSS, JavaScript and Go
----

While learning Go Language, i created this webapp to demonstrate how stuffs are done using Go,
mainly for serving static files, handling GET/POST requests, handling data and running the server.

## Prerequisites
* [Go](//go.dev/dl/)
* [MySQL](//www.mysql.com/downloads/)

## Clone the repository

```
git clone https://github.com/SunnyRaj84348/go-simple-webapp
cd go-simple-webapp
```

## Run the app

```
cd api
MYSQL_DATA="user:password@tcp(127.0.0.1:3306)/database" go run main.go
```
----

Head over your browser and open [localhost:8080](http://localhost:8080/) to access the running app.
