FROM golang:1.18-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  small-user

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY dist ./dist

RUN go mod download
RUN CGO_ENABLED=0 go build -o /dog-tracker

FROM alpine:3.16

RUN apk add --no-cache \
    ca-certificates \
    zlib-dev

COPY --from=base /dog-tracker /dog-tracker

EXPOSE 8080

CMD [ "/dog-tracker", "serve", "--http", "0.0.0.0:8080" ]
