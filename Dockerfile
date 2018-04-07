FROM golang

ENV APP_HOME /go-docker-sample
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . $APP_HOME
EXPOSE 8090
CMD ["/usr/local/go/bin/go", "run", "/go-docker-sample/server.go"]
