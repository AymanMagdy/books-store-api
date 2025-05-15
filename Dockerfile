FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN go build -o book-store

RUN chmod +x ./book-store

EXPOSE 8080

CMD ["./book-store"]
