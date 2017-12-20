FROM golang:alpine3.6 AS binary
ADD . /app
WORKDIR /app
RUN go build -o http

FROM alpine:3.6
WORKDIR /app
RUN apk update
RUN apk add docker
ENV PORT 8000
EXPOSE 8000
COPY --from=binary /app/http /app
CMD ["/app/http"]
