FROM golang:latest
WORKDIR /producer
COPY . .

RUN ["go", "build", "./main.go"]

ENTRYPOINT ./main