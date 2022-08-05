# Build
FROM golang:1.18.3-buster AS build

WORKDIR /app

COPY . ./

RUN go mod download
RUN GOARCH=amd64
RUN go build -o /go/bin/app

# Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /go/bin/app .

EXPOSE 8081

CMD ["./app"]