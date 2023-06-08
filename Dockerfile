FROM golang:1.19

WORKDIR /app
COPY . ./
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o /gophkeeper cmd/gophkeeper/main.go

EXPOSE 3000

CMD ["/gophkeeper"]