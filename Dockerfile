
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /code-finder



FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /code-finder /code-finder

EXPOSE 5000

USER nonroot:nonroot

ENTRYPOINT ["/code-finder"]