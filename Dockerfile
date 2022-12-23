# build
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
# CGO for sqlite
RUN CGO_ENABLED=1 GOOS=linux go build -o app cmd/server/main.go

# copy binary only
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/app /app
EXPOSE 8080

# because security
USER nonroot:nonroot

ENTRYPOINT ["/app"]
