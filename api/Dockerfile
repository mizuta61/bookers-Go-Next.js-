FROM golang:1.19.1
WORKDIR /app
COPY . . 
RUN go mod tidy
CMD [ "go", "run", "example.go" ]