.PHONY: build run clean

build:
	docker build -t car-rent-go .

run: build
	docker run --name car-rent-go -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=car-rent-go -d postgres
	docker run --link car-rent-go:postgres car-rent-go

clean:
	docker stop car-rent-go
	docker rm car-rent-go
