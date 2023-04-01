FROM golang:1.20-alpine
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main main.go

RUN cd /app
CMD ./main