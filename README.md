![](https://github.com/AymanArif/golang-microservices/workflows/Feature%20Gorilla%20Framework%20workflow/badge.svg)
# Feature 1: Restful Go using  Gorilla framework

# Table of Contents

- [Links](#links)



# APIs


curl localhost:9090/ -X POST -d '{"name": "Tfdea"}'

curl localhost:9090/

curl localhost:9090/{id} -X POST -d '{"name": "Tea"}'




# Links
1. Gorilla Framework
    - https://www.gorillatoolkit.org/pkg/mux
2. [Go Validator](https://github.com/go-playground/validator)    

3. Blogs



# JSON Validation


	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku"  validate:"required,sku"`


Custom SKU validation:

```


func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)

}

func validateSKU(fl validator.FieldLevel) bool {

	// sku is of format aqz-sws-qaf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)

	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true

}
```
