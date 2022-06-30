FROM golang:1.17-alpine AS builder
WORKDIR /app

#build logging settings.
ENV LOG_LEVEL DEBUG
ENV SERVICE sum-build
ENV VERSION v1

#copying all source.
COPY . /app
#installing req libs.
RUN apk update \
    && apk add --no-cache git \
    && apk --no-cache add build-base
#go build binary.
RUN cd /app/ \
    && go mod download \
    && go test ./... -v \
    && cd /app/cmd/iban \
    && go build -o iban
#building final image.
FROM alpine
WORKDIR /app


RUN apk update
#copying only binary from builder.
COPY --from=builder /app/cmd/iban /app

EXPOSE 8080
ENTRYPOINT ["/app/iban"]