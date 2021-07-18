FROM golang:1.16 AS runner

WORKDIR /go/src/app
COPY . .

RUN go mod vendor
RUN go build -o server main.go

EXPOSE 8080

ENTRYPOINT [ "./server"]
