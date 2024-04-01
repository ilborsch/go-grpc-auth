1. `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
2. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
3. `export PATH="$PATH:$(go env GOPATH)/bin"`
4. `protoc -I proto proto/auth/auth.proto --go_out=./generated/go --go_opt=paths=source_relative --go-grpc_out=./generated/go/ --go-grpc_opt=paths=source_relative`
5. `go get github.com/GolangLessons/protos`

`go run cmd/sso/main.go --config=config/local.yaml`