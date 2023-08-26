FROM golang:1.21.0 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY .env ./

COPY . ./

RUN GOOS=linux go build -o ./ports-app ./main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ports-app ./

COPY .env ./

EXPOSE 8000

CMD [ "./ports-app" ]
