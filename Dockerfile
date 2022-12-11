FROM golang:1.19.1-bullseye as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -trimpath -ldflags="-w -s" -o "jwt-checker"

FROM gcr.io/distroless/base-debian11 as dev
COPY --from=builder /app/jwt-checker /jwt-checker
CMD ["/jwt-checker"]
