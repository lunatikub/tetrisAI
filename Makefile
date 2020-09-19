all: tetrisBot

tetrisBot: 
	go build -o $@

test:
	go test -v ./...

clean:
	rm tetrisBot

.PHONY: unit-test tetrisBot clean
