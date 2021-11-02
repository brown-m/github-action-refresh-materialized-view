FROM golang:1.16-buster AS build

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o refresh-view cmd/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /app/refresh-view /refresh-view

USER nonroot:nonroot

ENTRYPOINT ["/refresh-view"]