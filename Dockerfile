FROM golang

ENV CGO_ENABLED=0
ADD . /tmp/app
WORKDIR /tmp/app
RUN go build main.go

FROM scratch

COPY --from=0 /tmp/app/main /usr/local/bin/quickstart-go-heartbeat

ENTRYPOINT ["/usr/local/bin/quickstart-go-heartbeat"]