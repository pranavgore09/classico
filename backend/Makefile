build:
	GOOS=linux go build -o app
	docker build -t pranav/saymath .

run:
	docker run -p 8080:8080 -p 9090:9090 pranav/saymath

clean:
	rm -f app