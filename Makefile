run:
	go run cmd/main.go

tidy:
	export GOSUMDB=off ; go mod tidy

run-stubby:
	stubby -d stubby/service.yaml -s 59040 --verbose