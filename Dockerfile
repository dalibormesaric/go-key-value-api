FROM golang:1.22-alpine

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/server

RUN go test -v ./...

EXPOSE 9000

CMD ["app"]
