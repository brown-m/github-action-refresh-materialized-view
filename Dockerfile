FROM golang:1.16-buster AS build

WORKDIR /app

COPY cmd/main.go ./
RUN go get github.com/lib/pq
RUN go build -o /run

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /run /run

USER nonroot:nonroot

ENTRYPOINT ["/run"]