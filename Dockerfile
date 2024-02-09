FROM golang:1.20-bullseye as build

RUN apt-get update -y
RUN apt-get install -y pkg-config \
     libvips-tools \
     libvips-dev

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download
RUN go mod tidy

COPY . /app/

RUN go build -o /app/main

# # --------

FROM ubuntu:22.04

WORKDIR /app

# Web service
EXPOSE 8082

RUN apt-get update -y
RUN apt-get install -y pkg-config \
     libvips-tools \
     libvips-dev

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Jakarta
RUN apt-get install -y tzdata
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime

COPY --from=build /app/conf /app/conf
COPY --from=build /app/main /app/main

CMD ["./main"]