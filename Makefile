build: kromatique.go test_runner.go
	go build kromatique.go
	go build test_runner.go

test: test_runner
	./test_runner

clear:
	rm ./test_results/* kromatique test_runner
