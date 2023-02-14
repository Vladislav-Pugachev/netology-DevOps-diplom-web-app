FROM golang:alpine3.17
WORKDIR /go/src/web-app
COPY . .
RUN go mod download && go mod verify
RUN ln -s /go/src/web-app /usr/local/go/src
RUN ln -s /go/src/web-app/.aws /root
RUN go build
EXPOSE 8080
ENTRYPOINT ["./web"]