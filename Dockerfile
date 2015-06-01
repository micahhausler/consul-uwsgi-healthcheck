# Script for registering a linked container with consul
#
# To build:
# $ docker run --rm -v $(pwd):/go/src/github.com/micahhausler/consul-registration -w /go/src/github.com/micahhausler/consul-registration golang:1.4.2 go build -v -a -tags netgo -installsuffix netgo -ldflags '-w'
# $ docker build -t micahhausler/consul-registration .
#
# To run:
# $ docker run --link <linked-container-name>:<container-alias>  micahhausler/consul-registration -h

FROM busybox

MAINTAINER Micah Hausler, <micah.hausler@ambition.com>

COPY uwsgi-healthcheck/uwsgi-healtcheck /bin/uwsgi-healthcheck
RUN chmod 755 /bin/uwsgi-healthcheck

ENTRYPOINT ["/bin/uwsgi-healthcheck"]
