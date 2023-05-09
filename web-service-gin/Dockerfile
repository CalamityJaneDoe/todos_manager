FROM golang:1.19.0

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o api .
