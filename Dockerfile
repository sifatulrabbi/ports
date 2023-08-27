FROM golang:1.21.0 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN GOOS=linux go build -o ./ports-app ./main.go


FROM alpine:latest

WORKDIR /app

ENV GO_ENV='production'
ENV GIN_MODE='release'
ENV PGUSER='admin'
ENV PGPASSWORD='p@$$word'
ENV PGHOST='locahost'
ENV PGDBNAME='ports'
ENV PGSSL='disable'
ENV PORT='8000'

COPY --from=builder /app/ports-app ./

EXPOSE 8000

CMD [ "./ports-app" ]
