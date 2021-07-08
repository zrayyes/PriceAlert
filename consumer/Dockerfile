FROM golang:latest
WORKDIR /consumer
COPY . .

RUN ["go", "build", "./main.go"]

ENTRYPOINT ./main