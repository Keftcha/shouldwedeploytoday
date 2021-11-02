.PHONY: binary clean run

binary:
	@CGO_ENABLED=0 go build -o ./out/server

clean:
	@rm -r ./out/*

run: binary
	@./out/server

image: binary
	@docker build --no-cache -t shouldwedeploytoday .

ctn-run: image
	@docker run -p 5461:5461 shouldwedeploytoday
