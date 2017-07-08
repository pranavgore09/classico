build:
	npm install -g grpc-tools
	grpc_tools_node_protoc --proto_path ./proto/ --js_out=import_style=commonjs,binary:./proto --grpc_out=./proto/ --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` ./proto/starwars.proto

run:
	npm start