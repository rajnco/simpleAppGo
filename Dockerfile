FROM golang:1.23.2-bookworm

WORKDIR /app

COPY . .

RUN go get

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
