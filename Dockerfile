FROM envoyproxy/envoy:v1.20-latest

FROM golang:1.18.1-bullseye

COPY --from=0 /usr/local/bin/envoy /usr/local/bin

COPY ./ /home

WORKDIR /home

RUN go build

CMD ./goodguy-core & envoy -c envoy.yaml
