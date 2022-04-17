protobuf:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./idl/goodguy-web.proto

cli:
	wget https://mirror.ghproxy.com/https://raw.githubusercontent.com/goodguy-project/goodguy-crawl/main/crawl_service/crawl_service.proto -O idl/crawl_service.proto
	protoc -I. --go_out=. --go_opt=Midl/crawl_service.proto=./idl --go-grpc_out=. --go-grpc_opt=Midl/crawl_service.proto=./idl ./idl/crawl_service.proto

build:
	docker build -t goodguy-core .

run:
	docker run  -it --name="goodguy-core" --network=container:goodguy_mysql goodguy-core

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
