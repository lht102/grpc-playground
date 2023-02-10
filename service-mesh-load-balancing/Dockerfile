FROM golang:1.19-alpine AS builder
WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o bin/adder-api cmd/adder-api/*
RUN CGO_ENABLED=0 go build -o bin/subtractor-api cmd/subtractor-api/*
RUN CGO_ENABLED=0 go build -o bin/transformer-api cmd/transformer-api/*
RUN CGO_ENABLED=0 go build -o bin/transformer-client cmd/transformer-client/*

FROM alpine:3.17 AS adder-api
RUN addgroup -S app \
    && adduser -S -G app app
WORKDIR /home/app
COPY --from=builder /src/bin/adder-api .
RUN chown -R app:app ./
USER app
CMD ["./adder-api"]

FROM alpine:3.17 AS subtractor-api
RUN addgroup -S app \
    && adduser -S -G app app
WORKDIR /home/app
COPY --from=builder /src/bin/subtractor-api .
RUN chown -R app:app ./
USER app
CMD ["./subtractor-api"]

FROM alpine:3.17 AS transformer-api
RUN addgroup -S app \
    && adduser -S -G app app
WORKDIR /home/app
COPY --from=builder /src/bin/transformer-api .
RUN chown -R app:app ./
USER app
CMD ["./transformer-api"]

FROM alpine:3.17 AS transformer-client
RUN addgroup -S app \
    && adduser -S -G app app
WORKDIR /home/app
COPY --from=builder /src/bin/transformer-client .
RUN chown -R app:app ./
USER app
CMD ["./transformer-client"]
