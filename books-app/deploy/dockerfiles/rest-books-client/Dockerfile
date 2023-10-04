FROM public.ecr.aws/bitnami/golang:1.20.4 AS builder

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make compile-rest-client

FROM alpine:3.14.2
COPY --from=builder ./service/books-app/build/out/rest-books-client rest-books-client
COPY --from=builder ./service/books-app/configs/rest-books-client.yaml rest-books-client.yaml
ENTRYPOINT ["./rest-books-client", "-configFile", "rest-books-client.yaml"]