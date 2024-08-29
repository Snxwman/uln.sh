run:
	@templ generate
	go build -o ./tmp/main src/main.go && air
