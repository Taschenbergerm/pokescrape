# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-pokescraper"
LABEL REPO="https://github.com/taschenbergerm/pokescraper"

ENV PROJPATH=/go/src/github.com/taschenbergerm/pokescraper

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/taschenbergerm/pokescraper
WORKDIR /go/src/github.com/taschenbergerm/pokescraper

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/taschenbergerm/pokescraper"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/pokescraper/bin

WORKDIR /opt/pokescraper/bin

COPY --from=build-stage /go/src/github.com/taschenbergerm/pokescraper/bin/pokescraper /opt/pokescraper/bin/
RUN chmod +x /opt/pokescraper/bin/pokescraper

# Create appuser
RUN adduser -D -g '' pokescraper
USER pokescraper

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/pokescraper/bin/pokescraper"]
