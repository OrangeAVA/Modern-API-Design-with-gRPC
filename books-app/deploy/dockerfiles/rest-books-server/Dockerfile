FROM public.ecr.aws/bitnami/golang:1.20.4 AS builder

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make compile-rest-server

FROM alpine:3.14.2
COPY --from=builder ./service/books-app/build/out/rest-books-server rest-books-server
COPY --from=builder ./service/books-app/configs/rest-books-server.yaml rest-books-server.yaml
ENTRYPOINT ["./rest-books-server", "-configFile", "rest-books-server.yaml"]