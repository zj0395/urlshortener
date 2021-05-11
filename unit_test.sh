go test $(go list ./... | grep -v meta) -count=1 -v -bench=.
