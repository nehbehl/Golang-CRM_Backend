# CRM Backend Project Solution

This repo contains implementation of CRM Backend using Go language.

The application handles the following 5 operations for customers in mock database:

* Getting a single customer through a /customers/{id} path
* Getting all customers through a /customers path
* Creating a customer through a /customers path
* Updating a customer through a /customers/{id} path
* Deleting a customer through a /customers/{id} path

## Setup

Use `go run main.go` to run any of the scripts in any IDE like Visual Studio Code. You would need Go language downloaded and extension for Go enabled in VS code to get it working.

If you encounter missing package issues, run `go get`.

## Dependencies

Run `go get github.com/google/uuid` for UUID libraries installation.


## Testing

Users can interact with the application (i.e., make API requests) by simply using Postman or cURL.

Also, main_test.go method can be run using `go test` to check each function implementation.

## Details


The application uses a router - gorilla/mux that supports HTTP method-based routing and variables in URL paths.
The Handler interface is used to handle HTTP requests sent to defined paths.
There are five routes that return a JSON response, and each is registered to a dedicated handler:    

      - getCustomers()
      - getCustomer()
      - addCustomer()
      - updateCustomer()
      - deleteCustomer()
      
      
