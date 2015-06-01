# Consul uwsgi Healthcheck
This project is for consul to run a health check against a uwsgi endpoint

## Build

### Compile inside the Docker container

If you just want to build the app, but not run it in a docker container, run:

```bash
docker run --rm -v $(pwd):/go/src/github.com/micahhausler/consul-uwsgi-healthcheck -w /go/src/github.com/micahhausler/consul-uwsgi-healthcheck/uwsgi-healthcheck golang:1.4.2 go build -v

```
If you want to build for busybox and have a mini-container:

```bash
docker run --rm -v $(pwd):/go/src/github.com/micahhausler/consul-uwsgi-healthcheck -w /go/src/github.com/micahhausler/consul-uwsgi-healthcheck/uwsgi-healthcheck golang:1.4.2 go build -v -a -tags netgo -installsuffix netgo -ldflags '-w'
```

### Cross-compile inside the Docker container

If you need to compile for a platform other than linux/amd64 (such as windows/386), this can be easily accomplished with the provided cross tags:

```bash
docker run --rm -v $(pwd):/go/src/github.com/micahhausler/consul-uwsgi-healthcheck -w /go/src/github.com/micahhausler/consul-uwsgi-healthcheck -e GOOS=windows -e GOARCH=386 golang:1.4.2-cross go build -v
```

Alternatively, you can build for multiple platforms at once:

```bash
docker run --rm -it -v $(pwd):/go/src/github.com/micahhausler/consul-uwsgi-healthcheck -w /go/src/github.com/micahhausler/consul-uwsgi-healthcheck golang:1.4.2-cross bash
$ for GOOS in darwin linux; do
>   for GOARCH in 386 amd64; do
>     go build -v -o myapp-$GOOS-$GOARCH
>   done
> done
```

## License
See [License](/LICENSE)
