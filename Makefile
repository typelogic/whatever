app: app.go
	go build app.go

test: app
	./app sample.json	

clean:
	@rm -f app

.PHONY: clean
