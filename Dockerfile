FROM alpine:edge
MAINTAINER Jim Alateras <jima@comware.com.au>

ENV GOPATH /golang
ENV APP_HOME /app
ENV APP_NAME version
ENV APP_SRC $GOPATH/src/github.com/jalateras/$APP_NAME
ENV APP_USER worker
ENV APP_GROUP worker
ENV PATH $APP_HOME:$GOPATH/bin:$PATH

ADD . $APP_SRC

RUN \
  addgroup -S $APP_USER && \
  adduser -S -s /bin/bash -G $APP_GROUP $APP_USER

RUN \
  apk add --update bash wget git mercurial bzr go make && \
  cd $APP_SRC && \
  make bootstrap build && \
  mkdir -p $APP_HOME && \
  cp ./$APP_NAME $APP_HOME/   && \
  chown -R $APP_USER:$APP_GROUP $APP_HOME && \
  apk del --purge wget git mercurial bzr go  && \
  rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $GOPATH

# Expose the http api port
EXPOSE 3000

USER $APP_USER
WORKDIR $APP_HOME
CMD ["./version"]



