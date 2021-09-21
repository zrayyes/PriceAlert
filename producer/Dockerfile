FROM golang:latest AS build
RUN mkdir /app
ADD . /app
WORKDIR /app

FROM build as dev
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -log-prefix=false -polling=true -build="go build ./main.go" -command="./main"

FROM build AS prod
RUN go build -o main .
CMD ["/app/main"]