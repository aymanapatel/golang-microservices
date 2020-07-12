
# Feature 3: Documenting APIs using Swagger/OAS

# Table of Contents

- [Links](#links)



# Links
1. [Go Swagger Website](https://goswagger.io/)
2. [Generate Swagger doc from Go source code](https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9)


# Run

1. `go build`
2. `go run main.go`


## API 

1. Works `curl localhost:9091/products -X POST -d '{"name": "Tea", "Price": 1.0, "sku":"aada-sdd"}'`

2. Does not work `curl localhost:9091/products -X POST -d '{"name": "Tea", "Price": 1.0, "sku":"aada-sdd-ddf"}'`



