build: clean
	@go build -o output/portscanner

clean:
	@rm -f output/portscanner


run: build
	@output/portscanner