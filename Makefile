run_local:
	go run ./cmd/sso/main.go --config=./config/local.yaml
test_local:
	go test ./tests

