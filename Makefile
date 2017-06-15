run:
	protoc -I ./api/say ./api/say/say.proto --go_out=plugins=grpc:./api/say
	protoc -I ./api/math ./api/math/math.proto --go_out=plugins=grpc:./api/math
	cd gateway && go build -o gateway
	cd backend && go build -o app
	docker-compose up -d --build
	rm -f gateway/gateway backend/app

stop:
	docker-compose down