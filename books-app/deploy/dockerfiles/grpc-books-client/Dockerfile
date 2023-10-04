FROM public.ecr.aws/bitnami/golang:1.20.4 AS builder

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make compile-grpc-client

FROM alpine:3.14.2
COPY --from=builder ./service/books-app/build/out/grpc-books-client grpc-books-client
COPY --from=builder ./service/books-app/configs/grpc-books-client.yaml grpc-books-client.yaml
ENTRYPOINT ["./grpc-books-client", "-configFile", "grpc-books-client.yaml"]