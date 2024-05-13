FROM golang:alpine3.19

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN apk add curl

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY . .

# RUN go build -o main .

# RUN swag init

RUN go test unit_test.go

EXPOSE 9000

# CMD ["./main", "run"]
CMD ["go", "run", "main.go"]
