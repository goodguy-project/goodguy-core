protobuf:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./idl/goodguy-web.proto

cli:
	wget https://mirror.ghproxy.com/https://raw.githubusercontent.com/goodguy-project/goodguy-crawl/main/crawl_service/crawl_service.proto -O idl/crawl_service.proto
	protoc -I. --go_out=. --go_opt=Midl/crawl_service.proto=./idl --go-grpc_out=. --go-grpc_opt=Midl/crawl_service.proto=./idl ./idl/crawl_service.proto

build:
	docker build -t goodguy-core .

run:
	-docker network create goodguy-net
	docker run -dit --name="goodguy-core" -p 9853:9887 --restart=always --network goodguy-net --network-alias goodguy-core goodguy-core

clean:
	-docker stop $$(docker ps -a -q --filter="name=goodguy-core")
	-docker rm $$(docker ps -a -q --filter="name=goodguy-core")
	-FOR /f "usebackq tokens=*" %%i IN (`docker ps -q -a --filter="name=goodguy-core"`) DO docker stop %%i
	-FOR /f "usebackq tokens=*" %%i IN (`docker ps -q -a --filter="name=goodguy-core"`) DO docker rm %%i

deploy:
	make build
	make run

restart:
	make clean
	make deploy

protobuf-tool:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
