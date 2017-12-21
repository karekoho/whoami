FROM golang AS binary
ADD . /go/src/github.com/karekoho/whoami
WORKDIR /go/src/github.com/karekoho/whoami
RUN go-wrapper download docker.io/go-docker
RUN go-wrapper install

FROM golang
WORKDIR /go/src/github.com/karekoho/whoami
ENV PORT 8000
EXPOSE 8000
COPY --from=binary /go/bin/whoami /go/bin
CMD ["whoami"]
