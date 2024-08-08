# Go Monorepo

## Running Tests
```bash
go test ./pkg/...  -coverprofile=coverage.out
```

To view the tests visually run
```bash
go tool cover -html=coverage.out
```
