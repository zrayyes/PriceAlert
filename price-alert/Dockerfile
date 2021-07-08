FROM golang:latest
WORKDIR /price-alert
COPY . .

RUN ["go", "build", "./main.go"]

ENTRYPOINT ./main
EXPOSE 8080