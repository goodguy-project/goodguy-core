FROM envoyproxy/envoy:v1.20-latest

FROM kubeless/unzip

FROM golang:1.18-bullseye

COPY --from=0 /usr/local/bin/envoy /usr/local/bin
COPY --from=1 /usr/bin/unzip /usr/local/bin

RUN mkdir /usr/local/protoc
WORKDIR /usr/local/protoc
RUN wget https://mirror.ghproxy.com/https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-linux-x86_64.zip -O protoc.zip \
    && unzip protoc.zip && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="/usr/local/protoc/bin:${PATH}"

COPY ./ /home
WORKDIR /home
RUN make protobuf && make cli && go build

CMD ./goodguy-core & envoy -c envoy.yaml
