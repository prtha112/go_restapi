# Build
FROM golang:1.18.3-buster AS build

WORKDIR /app

COPY . ./

RUN make build


# Deploy
FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/app /app

EXPOSE 8081

# USER nonroot:nonroot
USER root:root

CMD ["/app"]

