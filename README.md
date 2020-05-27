
# Feature 1: RESTful GO

# Table of Contents

- [API Specification](#api-specification)
- [Links](#links)
- [Setup Go Environment](#setup-go-environment)
- [Serve Mux and handler](#servermux-and-handler)

# API Specification

|API|Verb|curl|
|--|--|--|
|/||`curl http://localhost:9090`|
|/hello||`curl http://localhost:9090/hello -d 'Ayman'`|
|/goodbye||`curl http://localhost:9090/goodbye`|
|/products|GET|`curl http://localhost:9090/products | jq`|
|/products|POST|`curl -v http://localhost:9090/products -d '{}'| jq`|
|/products|PUT|`curl -v http://localhost:9090/products -d '{}'| jq`|
|/products|POST|`curl -v  localhost:9090/1 -XPUT -d '{"name":"asfsdsyman"}'| jq`|


# Links

- [Recap of Request Handling in Go](https://www.alexedwards.net/blog/a-recap-of-request-handling)

- [Go documentaton](https://golang.org/doc/)

- [Golangs' `json/encoding`](https://golang.org/pkg/encoding/json/)

- [RESTful best practices](https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design)

# Setup Go environment

## Demo with `go mod`

    mkdir my-demo // Create a project folder, or clone a github repository    
    cd my-demo
    touch main.go // Create a file if we are starting the project
    go mod init my-demo // This will initiate the module
    go build // This will build

Now if you add an import, if VSCODE is configured properly (user.settings), then auto-complete for used imports

## Issues

- [Stack Overflow](https://stackoverflow.com/questions/60680470/could-not-import-local-modules-in-golang/60681078#60681078)
- Set up Go Environment and Go modules in VSCode
  - [Medium](https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a)

- [Go Modules outside $GO_PATH in Docker](https://devandchill.com/posts/2019/03/go-modules-working-outside-gopath/)

- Change `hello.go` to `handlers.go`

# ServeMux and handler

1. DefaultServeMux
Do not use `DefaultServeMux`, as it is a global variable. Hence, it is not secure.

2.
