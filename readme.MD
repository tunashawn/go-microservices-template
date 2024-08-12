### Simple golang template for microservices using go 1.22 net/http package

This is a basic layout for microservices projects using net/http package of **go 1.22** based on what I had at my past
company. We used gin gonic back then, but with the new upgrade of net/http package in go 1.22, I want to make a new
template for it as
it is faster way more simpler compares to gin-gonic for a simple web server usage.

This can be used with other http packages, you just have to change code in cmd and internal/response to match with your
package.

### Project structure

This project structure is created based on suggestion from go team `https://go.dev/doc/modules/layout` and my experience
in writing go microservices.

```|-- .
|-- cmd
|   |-- service_A
|-- internal
|   |-- config
|   |-- db
|   |-- response
|-- pkg
|   |-- service_A
|   |   |-- controllers
|   |   |-- models
|   |   |-- repository
|   |   |-- services
|-- go.mod
|-- go.sum
```

### Explanation

1. **cmd**: main file of each microservices
2. **internal**: common packages between services
    1. **config**: manage all environment variables
    2. **db**: simple code to connect to your databases. It will automatically get env config so in main file.
    3. **response**: http responses. Main functionality is to return http responses and print all logs along with them
       so others packages don't have to deal with logging. As in the end, a http request will be either success or
       fail (bad request, internal server error, unauthorized, ...), then all logging in the lower levels should also be
       divided in to these kind of status. This makes logging simpler and easier to track. I also use only log/slog as
       all logging is in one place and it's fast and simple.
        * example of a succeed request log:
       ```json
         "time": "2024-08-11T18:51:04.946485+07:00",
         "level": "INFO",
         "msg": "ok",
         "request": {
             "method": "GET",
             "uri": "/ping",
             "request_body": ""
         },
         "response": {
             "meta": {
                 "code": 200,
                 "message": "ok"
             },
             "data": "pong"
         }
       ```
        * example of a failed request log:
       ```json
       "time": "2024-08-11T19:28:16.54059+07:00",
       "level": "ERROR",
       "msg": "something happened",
       "request": {
           "method": "GET",
           "uri": "/ping",
           "request_body": ""
       },
       "response": {
           "code": 500,
           "message": "something happened",
           "data": null
       }
       ```

3. **pkg**: main packages of each microservice

### Additional Packages

* `github.com/kelseyhightower/envconfig` simplified reading environment variables task
* `github.com/pkg/errors` I highly recommend using this package to wrap errors through layers of code to make all middle
  functions cleaner and tracing bugs easier. For example, if querying database has an errors, it then will
  be passed through /repository and /services to /controller, then print logs and returns http response at the
  package /response. This helps reducing bunch of logging codes and giving a traceable log. All errors should be wrapped
  for a better reading and tracing ability and printing logs should be only on the outermost layer such as controller.
* `github.com/stretchr/testify` is used to mock lower layers of code. This would reduce many efforts in writing tests.
