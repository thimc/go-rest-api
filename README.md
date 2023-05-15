# REST API using Go

* main.go sets up the handlers and listening on the TCP network address.
* handlers.go contains all the CRUD (create, read, update and delete) HTTP handlers.
* utils.go contains a wrapper function that sets the response headers and HTTP
  code. It also marshals the response struct as a JSON and writes it to the
  body.

The "database" is initialized in the main.go file and the data is not
persistent.

## Note
This API was created for educational purposes only and should only be treated as
such, it is not production ready.
