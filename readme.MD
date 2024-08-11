This is a basic layout for microservices projects using net/http package of **go 1.22** based on what I had at my past
company. We used gin gonic back then, but with the new upgrade of net/http package in go 1.22, I want to make a new template for it as
it is faster way more simpler compares to gin-gonic for a simple web server usage.

This can be used with other http packages, you just have to change code in cmd and internal/response to match with your
package.

### Project structure
This project structure is created based on suggestion from go team `https://go.dev/doc/modules/layout` and my experience in writing go microservices.
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