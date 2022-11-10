module example

go 1.19

require (
	geecache v0.0.0
	github.com/golang/protobuf v1.5.2
)

require google.golang.org/protobuf v1.26.0 // indirect

replace geecache => ./geecache
