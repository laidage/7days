module main

go 1.19

require gee v0.0.0

replace gee => ./gee

require geecache v0.0.0

require (
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace geecache => ./geecache
