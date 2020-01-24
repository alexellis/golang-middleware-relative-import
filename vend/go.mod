module github.com/alexellis/vend/vend

go 1.13

replace github.com/alexellis/vend/vend/handlers => ./handlers

// go test added this, not me.
require github.com/alexellis/vend/vend/handlers v0.0.0-00010101000000-000000000000
