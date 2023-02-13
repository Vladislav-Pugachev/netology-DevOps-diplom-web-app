FROM golang:alpine3.17
WORKDIR /go/src/web-app
COPY . .
RUN go mod download
RUN ln -s /go/src/web-app /usr/local/go/src
EXPOSE 80 443 8080
ENTRYPOINT ["go", "run", "web]