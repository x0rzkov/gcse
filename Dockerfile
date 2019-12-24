FROM golang:1.13-alpine AS builder

ARG TWINT_GENERATOR_VERSION

RUN apk add --no-cache make

COPY .  /go/src/github.com/x0rzkov/gcse
WORKDIR /go/src/github.com/x0rzkov/gcse

RUN cd /go/src/github.com/x0rzkov/gcse/pipelines/crawler \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/pipelines/indexer \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/pipelines/mergedocs \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/pipelines/spider \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/pipelines/tocrawl \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/tools/countdocs \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/tools/dump \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/tools/exps \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/tools/fillfound \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/tools/fixcrawldb \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/service/stored \
    && go build -v \
    && ls -l

RUN cd /go/src/github.com/x0rzkov/gcse/service/web \
    && go build -o web-server -v \
    && ls -l

FROM alpine:3.10 AS runtime

# Build argument
ARG VERSION
ARG BUILD
ARG NOW

# Install runtime dependencies & create runtime user
RUN apk --no-cache --no-progress add ca-certificates \
 && mkdir -p /opt \
 && adduser -D gcse -h /opt/gcse -s /bin/sh \
 && su gcse -c 'cd /opt/gcse; mkdir -p bin config data'

# Switch to user context
USER gcse
WORKDIR /opt/gcse/data

# Copy gcse binary to /opt/gcse/bin
COPY --from=builder /go/src/github.com/x0rzkov/gcse/pipelines/crawler/crawler /opt/gcse/bin/crawler
COPY --from=builder /go/src/github.com/x0rzkov/gcse/pipelines/indexer/indexer /opt/gcse/bin/indexer
COPY --from=builder /go/src/github.com/x0rzkov/gcse/pipelines/mergedocs/mergedocs /opt/gcse/bin/mergedocs
COPY --from=builder /go/src/github.com/x0rzkov/gcse/pipelines/spider/spider /opt/gcse/bin/spider
COPY --from=builder /go/src/github.com/x0rzkov/gcse/pipelines/tocrawl/tocrawl /opt/gcse/bin/tocrawl

COPY --from=builder /go/src/github.com/x0rzkov/gcse/service/stored/stored /opt/gcse/bin/stored
COPY --from=builder /go/src/github.com/x0rzkov/gcse/service/web/web-server  /opt/gcse/bin/web-server 
COPY --from=builder /go/src/github.com/x0rzkov/gcse/service/web/web  /opt/gcse/data/service/web/web
COPY --from=builder /go/src/github.com/x0rzkov/gcse/service/web/static  /opt/gcse/data/service/web/static
COPY --from=builder /go/src/github.com/x0rzkov/gcse/service/web/resource  /opt/gcse/data/service/web/resource
COPY --from=builder /go/src/github.com/x0rzkov/gcse/service/web/images  /opt/gcse/data/service/web/images
COPY --from=builder /go/src/github.com/x0rzkov/gcse/service/web/css  /opt/gcse/data/service/web/css

COPY --from=builder /go/src/github.com/x0rzkov/gcse/tools/countdocs/countdocs /opt/gcse/bin/countdocs
COPY --from=builder /go/src/github.com/x0rzkov/gcse/tools/dump/dump /opt/gcse/bin/dump
COPY --from=builder /go/src/github.com/x0rzkov/gcse/tools/exps/exps /opt/gcse/bin/exps
COPY --from=builder /go/src/github.com/x0rzkov/gcse/tools/fillfound/fillfound /opt/gcse/bin/fillfound
COPY --from=builder /go/src/github.com/x0rzkov/gcse/tools/fixcrawldb/fixcrawldb /opt/gcse/bin/fixcrawldb

ENV PATH $PATH:/opt/gcse/bin

RUN cd /opt/gcse/data && mkdir -p data/docs-updated data/docs data/newdocs data/person data/crawler data/tocrawl data/package data/store

# Container metadata
LABEL name="gcse" \
      version="$VERSION" \
      build="$BUILD" \
      architecture="x86_64" \
      build_date="$NOW" \
      vendor="x0rzkov" \
      maintainer="x0rzkov <x0rzkov@protonmail.com>" \
      url="https://github.com/x0rzkov/gcse" \
      summary="Dockerized gcse project" \
      description="Dockerized gcse project" \
      vcs-type="git" \
      vcs-url="https://github.com/x0rzkov/gcse" \
      vcs-ref="$VERSION" \
      distribution-scope="public"

# Container configuration
VOLUME ["/opt/gcse/data"]
CMD ["web-server"]




