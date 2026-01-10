module main

go 1.25.5

replace fizzbuzz => ../examples/fizzbuzz

replace defers/defers => ../examples/defers

require (
	defers/defers v0.0.0-00010101000000-000000000000 // indirect
	fizzbuzz v0.0.0-00010101000000-000000000000 // indirect
)
