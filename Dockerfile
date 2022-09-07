# syntax=docker/dockerfile:1

FROM golang:1.17 AS build
WORKDIR /project/src
COPY . /project/src
COPY main /project/src
RUN go mod tidy -compat=1.17
RUN go mod download


RUN go build -o /almanac ./main

###

FROM debian:buster-slim
COPY --from=build /almanac /almanac
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 80/tcp
EXPOSE 443/tcp
CMD [ "/almanac" ]