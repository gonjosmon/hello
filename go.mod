module example/user/hello

go 1.17

require (
	example/user/greetings v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.7
)

replace example/user/greetings => ../greetings
