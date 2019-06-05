FROM golang:stretch as build
ADD . /src
RUN cd /src && CGO_ENABLED=0 go build -o prom-fake

FROM alpine
WORKDIR /app
COPY --from=build /src/prom-fake /app
EXPOSE 8080
ENTRYPOINT ["./prom-fake"]
