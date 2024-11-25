install:
	go install github.com/favadi/protoc-go-inject-tag@latest
	go install github.com/cloudzenith/DouTok/backend/gopkgs/tools/protoc-gen-go-http@latest
	go install github.com/cloudzenith/DouTok/backend/gopkgs/tools/protoc-gen-openapi@latest

gen-proto:
	protoc \
		--proto_path=./api \
		--proto_path=./third_party \
		--go_out=paths=source_relative:./api \
		--go-http_out=paths=source_relative:./api \
		--openapi_out=fq_schema_naming=true,default_response=false:. \
		./api/msapi/*.proto
	protoc-go-inject-tag -input="./api/msapi/*.pb.go"

build:
	docker build -t manage-api-service .

build1:
	 docker build --build-arg proxy="https://goproxy.cn" -t base-service .
