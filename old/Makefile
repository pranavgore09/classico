run:
	protoc -I ./api/say ./api/say/say.proto --go_out=plugins=grpc:./api/say
	protoc -I ./api/math ./api/math/math.proto --go_out=plugins=grpc:./api/math
	cd gateway && go build -o gateway
	cd backend && go build -o app
	cd daily-backend && npm install -g grpc-tools
	cd daily-backend && grpc_tools_node_protoc --proto_path ../api/daily/ --js_out=import_style=commonjs,binary:. --grpc_out=./ --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` ../api/daily/daily.proto
	docker-compose up -d --build

stop:
	docker-compose down

clean:
	rm -f gateway/gateway backend/app