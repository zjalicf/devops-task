FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /app/bin/app .

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /app/bin/app /app
ENV PORT=11000
USER nonroot:nonroot
ENTRYPOINT ["/app"]