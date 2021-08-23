FROM golang:1.16

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify

RUN go build -o /app

EXPOSE 8080

CMD [ "/app" ]
