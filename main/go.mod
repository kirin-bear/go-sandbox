module main

go 1.25.5

replace fizzbuzz => ../fizzbuzz

replace defers/defers => ../defers

require (
	defers/defers v0.0.0-00010101000000-000000000000 // indirect
	fizzbuzz v0.0.0-00010101000000-000000000000 // indirect
    github.com/stretchr/testify v1.7.0 
)
