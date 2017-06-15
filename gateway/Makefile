build:
	GOOS=linux go build -o gateway
	docker build -t pranav/gateway .

run:
	docker run -p 1111:1111 -p 8080:8080 pranav/gateway

clean:
	rm -f gateway