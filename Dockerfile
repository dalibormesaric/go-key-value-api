FROM golang

WORKDIR ./src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["src"]