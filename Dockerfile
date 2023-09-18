FROM golang:1.21.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /liapi

FROM build AS test
RUN go test -v ./...

FROM scratch as release

WORKDIR /
COPY --from=build /liapi /liapi

EXPOSE 8080

ENTRYPOINT ["/liapi"]