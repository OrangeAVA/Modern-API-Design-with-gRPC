FROM public.ecr.aws/bitnami/golang:1.20.4 AS builder

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make compile-grpc-server

FROM alpine:3.14.2
COPY --from=builder ./service/books-app/build/out/grpc-books-server grpc-books-server
COPY --from=builder ./service/books-app/configs/grpc-books-server.yaml grpc-books-server.yaml
ENTRYPOINT ["./grpc-books-server", "-configFile", "grpc-books-server.yaml"]