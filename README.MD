## Quickstart
1. Install Go 1.24+
2. `cd grade-calculator`
3. Run tests:
   ```bash
   go test ./... -coverprofile="coverage.out"
   go tool cover -func="coverage.out"