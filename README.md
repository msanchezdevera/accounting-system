# accounting-system
Money accounting system

This is a web application that simulates an accounting system.
For simplicity, only one user with one account exists.

# How yo run locally

Clone the project and follow these instructions:

## Requisits
- Install Go 1.15

## Execution

Execute the command:

`./accounting`

You should see the following message:

```
2020-10-18T12:07:59.372 info [mariano-Inspiron-5423]: Launching accounting system
2020-10-18T12:07:59.373 info [mariano-Inspiron-5423]: Starting Server (HTTP) on :8080
```

This indicates that the server starting runnig at localhost, on port 8080

# Tests

## How to execute tests

On the root folder execute the command:
```
go test ./...
```

## Coverage

To test code coverage, execute the following commands:
```
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

# Examples

### Create a transaction
```
curl -X POST \
  http://localhost:8080/transactions \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "type": "credit",
    "amount": 222
}'
```

### Retrieve a transaction
```
curl -X GET \
  http://localhost:8080/transactions/fe38a9b5-249d-4dec-b10d-1115c6f30db6
```

### Retrieve all transactions

```
curl -X GET \
  http://localhost:8080/transactions
```

### Retrieve current balance

The response includes the current account balance.

```
curl -X GET \
  http://localhost:8080/balance
```

### Postman

If you use Postman, you can import the [project collection](postman/Acccount-system.postman_collection.json) and test all requests.