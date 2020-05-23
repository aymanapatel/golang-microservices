module github.com/AP/Ch3-GOMS

replace github.com/AP/Ch3-GOMS/handlers => ./handlers

replace github.com/AP/Ch3-GOMS/data => ./data

go 1.14

require (
	github.com/AP/Ch3-GOMS/data v0.0.0-00010101000000-000000000000 // indirect
	github.com/AP/Ch3-GOMS/handlers v0.0.0-00010101000000-000000000000
)
