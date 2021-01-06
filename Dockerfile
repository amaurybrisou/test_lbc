FROM golang:1.15-alpine as base
WORKDIR /go/app
ENV CGO_ENABLED=0
COPY . .
RUN go mod download

FROM base AS unit-test
RUN --mount=type=cache,target=/root/.cache/go-build \
    go test -v ./...

FROM base as build
RUN go build -o main .

FROM scratch as bin
WORKDIR /
COPY --from=build /go/app/main .
EXPOSE 8000

CMD ["./main"]
