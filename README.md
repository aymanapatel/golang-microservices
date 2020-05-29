
# Feature 5: gRPC

# Table of Contents

- [Links](#links)



# Links

1. [gRPC docs](https://grpc.io/docs/)
2. [Intro to gRPC](https://www.youtube.com/watch?v=QyxCX2GYHxk)


# gRPC Introduction

A binary HTTP/2 protocol which is faster and more efficient due to HTTP/2 and compressing of request using Protocol buffers.

![gRPC overiew](https://grpc.io/img/landing-2.svg)
Support languages/runtimes:
- Android
- Java, Kotin/JVM
- Go
- .NET/C#
- Node
- Web
- Ruby
- PHP, Obejective-C, Dart

# gRPC Concepts

# gRPC Miscellaneous

## Error handling


### General Errors

|Status Code|Case|
|--|--|
|`GRPC_STATUS_CANCELLED`||
|`GRPC_STATUS_DEADLINE_EXCEEDED`||
|`GRPC_STATUS_UNIMPLEMENTED`||
|`GRPC_STATUS_UNAVAILABLE`||
|`GRPC_STATUS_UNKNOWN`||


### Network failures


|Status Code|Case|
|--|--|
|`GRPC_STATUS_DEADLINE_EXTENSION`||
|`GRPC_STATUS_UNAVAILABLE`||

### Protocol Errors

|Status Code|Case|
|--|--|
|`GRPC_STATUS_INTERNAL`||
|`GRPC_STATUS_UNIMPLEMENTED`||
|`GRPC_STATUS_RESOURCE_EXHAUSTED`||
|`GRPC_STATUS_INTERNAL`||
|`GRPC_STATUS_UNKNOWN`||
|`GRPC_STATUS_UNAUTHENTICATED`||
|`GRPC_STATUS_UNAUTHENTICATED`||
|`GRPC_STATUS_INTERNAL`||
|`GRPC_STATUS_INTERNAL`||

### Richer Error Model
If using Protocol Buffers, you might want to use a richer mdoel developed by (Google)[https://cloud.google.com/apis/design/errors#error_model]


## AuthN

## Benchmarking


# gRPC vs RESTful vs GraphQL


# Refer Golang doc.

