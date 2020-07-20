# Go Microservice Project


<p align="center">
  Repository to learn Golang Microservices by building an end-to-end project. <br>
   🔨 Libraries used: 🔨
</p>



<p align="center">
  <a href="https://miro.medium.com/max/750/1*xLKFtlXiTPif_hTAIlXRjw.jpeg">
    <img alt="" src="https://miro.medium.com/max/750/1*xLKFtlXiTPif_hTAIlXRjw.jpeg" height="70" width="50" />
  </a>
  <a href="https://miro.medium.com/max/750/1*xLKFtlXiTPif_hTAIlXRjw.jpeg">
    <img alt="" src="https://upload.wikimedia.org/wikipedia/commons/a/ab/Swagger-logo.png" height="70" width="70" />
  </a>    
  <a href="https://miro.medium.com/max/750/1*xLKFtlXiTPif_hTAIlXRjw.jpeg">
    <img alt="" src="https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQGphDpO2--RtpVgIv_UAy9NObzHXWwbSM49wB0q1U_H0sbxQm80ZGrZNnyujl826dQ_XM8NGH-r7nm1Q&usqp=CAU" height="70" width="120" />
  </a>    
  
</p>


![](https://github.com/AymanArif/golang-microservices/workflows/Go/badge.svg)

# Feature 0: Setup Go

# Table of Contents

- [Links](#links)
- [Setup Go Environment](#setup-go-environment)
- [Serve Mux and handler](#servermux-and-handler)


# Links
1. Set up Go Environment and Go modules in VSCode
[Medium](https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a)

2. [Go Modules outside #GO_PATH in Docker](https://devandchill.com/posts/2019/03/go-modules-working-outside-gopath/)

# Setup Go environment

## Demo with `go mod`

    mkdir my-demo // Create a project folder, or clone a github repository    
    cd my-demo
    touch main.go // Create a file if we are starting the project
    go mod init my-demo // This will initiate the module
    go build // This will build
`
Now if you add an import, if VSCODE is configured properly (user.settings), then auto-complete for used imports

# Working
[Stack Overflow](https://stackoverflow.com/questions/60680470/could-not-import-local-modules-in-golang/60681078#60681078)
- Change `hello.go` to `handlers.go`


# ServeMux and handler

[Reference](https://www.alexedwards.net/blog/a-recap-of-request-handling)
[Go documentaton]()
1. DefaultServeMux
Do not use `DefaultServeMux`, as it is a global variable. Hence, it is not secure.

2.


# Refer Golang doc.

