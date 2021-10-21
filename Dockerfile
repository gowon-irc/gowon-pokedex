FROM golang:alpine as build-env
COPY . /src
WORKDIR /src
RUN go build -o gowon-pokedex

FROM alpine:3.14.2
WORKDIR /app
COPY --from=build-env /src/gowon-pokedex /app/
ENTRYPOINT ["./gowon-pokedex"]
