FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY listing ./listing

RUN CGO_ENABLED=0 GOOS=linux go build -o pokemon-server ./cmd

EXPOSE 8080

# Run
CMD ["./pokemon-server"]
