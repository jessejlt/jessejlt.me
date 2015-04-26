.PHONY: build, run

default: run

build:
	docker build -t jessejlt.me .

run: build
	docker run -d -p 8080:80 jessejlt.me