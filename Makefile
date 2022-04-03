protobuf:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./idl/goodguy-web.proto

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
