#
# Build
#
FROM golang:1.18 AS build

ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /soja
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make install

#
# Main
#
FROM alpine

EXPOSE 1317
EXPOSE 26656
EXPOSE 26657

COPY --from=build /go/bin/sojanoded /usr/bin/sojanoded
COPY --from=build /go/bin/sojagen /usr/bin/sojagen
