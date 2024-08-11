This is a basic layout for microservices projects using net/http package of **go 1.22** based on what I had at my past
company. We used gin gonic, but with new release of net/http package of go1.22, I want to make a new template for it as
it is faster way more simpler to use than gin gonic for simple web server.

This can be used with other http packages, you just have to change code in cmd and internal/response to match with your
package.

### Project structure

```|-- .
|-- cmd
|   |-- cmd/service_A
|   |   |-- service_A.go
|-- internal
|   |-- config
|   |   |-- config.go
|   |-- db
|   |   |-- mongo.go
|   |   |-- mysql.go
|   |-- response
|   |   |-- response.go
|-- pkg
|   |-- pkg/service_A
|   |   |-- controllers
|   |   |   |-- controller.go
|   |   |-- models
|   |   |-- repository
|   |   |   |-- mongo.go
|   |   |   |-- mysql.go
|   |   |-- services
|   |   |   |-- service.go
|-- go.mod
|-- go.sum
```

### Explanation

1. **cmd**: main file of each microservices
2. **internal**: common packages
    1. **config**: using package `github.com/kelseyhightower/envconfig` for to manage all environment variables. It is
       pretty simple to use and make our codes cleaner.
    2. **db**: simple code to connect to your databases. It will automatically get env config so in main file.
    3. **response**: http responses. Using basic net/http package in 1.22
3. **pkg**: main packages of each microservice

### Notes
* I personally recommend using package `github.com/pkg/errors` to simply wrap errors through layers of code and only
  print logs on the outermost layer such as controllers to keep codes simpler and easier to read.
* If a microservice has multiple repositories connected to a same database, then initializing db should be changed and placed in the main file
* Whatever changes, go codes should remain as simple as possible.