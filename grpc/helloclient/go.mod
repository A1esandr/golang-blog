module helloclient

go 1.16

require (
	github.com/A1esandr/golang-blog/grpc/hello v0.0.0-20210406181329-3322ba9ee095
	google.golang.org/grpc v1.36.1
)

replace github.com/A1esandr/golang-blog/grpc/hello => ../hello
