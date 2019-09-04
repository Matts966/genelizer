FROM golang:alpine
COPY ./ /work
RUN apk update && apk add --no-cache make git
WORKDIR /work
RUN make dep && CGO_ENABLED=0 make install
ENTRYPOINT [ "/bin/sh" ]
