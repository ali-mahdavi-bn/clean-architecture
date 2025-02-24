run:
	go run cmd/main.go
test:
	go test ./tests/... -v
test-integration:
	TEST_TYPE=integration go test ./tests/integration/... -v