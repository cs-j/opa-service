# Guide

To run this project, compile the `main.go` file with `go build main.go`
and then run the compiled executable with `./main`, or do both at once with `go run main.go`.

TODO to verify app is working

### Get a list of all users
```
curl http://localhost:8080/users
```

### Get a particular user
```
curl http://localhost:8080/users/Camille
```

### Delete a user
```
curl -XDELETE http://localhost:8080/users/Camille
```

### Create a user
```
curl -H 'content-type: application/json' -d '{"name": "Camille"}' http://localhost:8080/users
```