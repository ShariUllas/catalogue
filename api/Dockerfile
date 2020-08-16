FROM golang:1.12.7-alpine AS builder

WORKDIR /catalogue/api

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh build-base 

COPY go.mod go.sum ./

RUN  go mod download 

COPY internal internal

# # Install binary
RUN ["go", "install", "internal/cmd/main.go"]

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /opt/
ENV PATH /opt:$PATH

COPY --from=builder /go/bin/ .

COPY internal/core/data/migrations internal/core/data/migrations

CMD ["./main", "-serve"]  
