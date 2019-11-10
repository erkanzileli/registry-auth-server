FROM golang:1.13-alpine AS builder
WORKDIR /go/src/github.com/erkanzileli/registry-auth-server
COPY . .
RUN go install
FROM alpine
ENV GIN_MODE release
USER nobody:nobody
COPY --from=builder /go/bin/registry-auth-server /registry-auth-server
CMD [ "/registry-auth-server" ]
