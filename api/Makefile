build:
	protoc -I ./say ./say/say.proto --go_out=plugins=grpc:./say
	protoc -I ./math ./math/math.proto --go_out=plugins=grpc:./math