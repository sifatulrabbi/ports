FROM golang:1.21.0 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN GOOS=linux go build -o ./ports-app ./main.go


FROM alpine:latest

WORKDIR /app

RUN echo -e "PORT=8000\nPGHOST=localhost\nPGDBNAME=postgres\nPGSSL=disable" > .env

COPY --from=builder /app/ports-app ./

EXPOSE 8000

CMD [ "./ports-app" ]
