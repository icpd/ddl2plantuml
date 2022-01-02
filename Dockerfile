FROM golang:latest AS builder

WORKDIR /src
COPY . .
RUN go build -o main main.go

FROM scratch

COPY --from=builder /src/main .

ENTRYPOINT ["/main"]
