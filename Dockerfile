FROM golang:1.23.4 AS builder
LABEL maintainer="Stepan K. <xamust@gmail.com>"
WORKDIR /app
COPY . .
RUN make migrup
RUN make build

FROM golang:1.23.4
WORKDIR /app
COPY --from=builder /app .
ENTRYPOINT ["./bin/app"]
EXPOSE 8082
