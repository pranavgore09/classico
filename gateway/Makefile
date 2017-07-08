build:
	npm install -g grpc-tools
	grpc_tools_node_protoc --proto_path ../services/*/proto/ --js_out=import_style=commonjs,binary:./services_pb --grpc_out=./services_pb/ --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` ../services/*/proto/*.proto

run:
	npm start