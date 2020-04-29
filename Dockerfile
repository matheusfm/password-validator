FROM golang:1.13-stretch as builder
ENV GO111MODULE=on
COPY . /password-validator
WORKDIR /password-validator
RUN go mod download && \
    go test ./... && \
    CGO_ENABLED=0 GOOS=linux go build -o bin/application

FROM scratch
COPY --from=builder /password-validator/bin/application /password-validator/bin
EXPOSE 8000
ENTRYPOINT ["./password-validator/bin"]