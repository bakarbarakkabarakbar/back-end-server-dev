FROM golang:alpine
EXPOSE 8081:8081
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o back-end-server-dev
CMD ["./back-end-server-dev"]
