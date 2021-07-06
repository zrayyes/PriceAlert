FROM golang:latest
WORKDIR /price-alert
COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

# Do not leave -polling=true on Production
ENTRYPOINT CompileDaemon -log-prefix=false -polling=true -build="go build ./main.go" -command="./main"
EXPOSE 8080