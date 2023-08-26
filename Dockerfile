FROM golang:1.21-alpine AS build-stage

RUN apk add upx
WORKDIR /stations-src
COPY . .
RUN go build -o /stations .
RUN upx /stations

FROM alpine:latest AS release-stage

COPY --from=build-stage /stations /stations
# -- Environment variables
ENV MEMPHIS_STATION     "alerts"
ENV MEMPHIS_PRODUCER    "alerts-producer"
ENV MEMPHIS_HOST        "memphis"
ENV MEMPHIS_USERNAME    "root"
ENV MEMPHIS_PASSWORD    "memphis"
ENV MEMPHIS_CONN_TOKEN  ""
ENV POSTGRES_DSN        "host=postgres user=sulcud password=sulcud dbname=sulcud port=5432 sslmode=disable"
# -- Environment variables
ENTRYPOINT [ "sh", "-c", "/stations :5000" ]