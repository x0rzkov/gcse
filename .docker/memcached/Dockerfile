FROM alpine:3.10
MAINTAINER x0rzkov@protonmail.com

# Configurations.
ENV MEMCACHED_MEMORY 128
ENV MEMCACHED_MAX_CONNECTIONS 1024
ENV MEMCACHED_MAX_ITEM_SIZE 4M

RUN apk --no-cache --update add memcached supervisor && \
    rm -rf /var/cache/apk/*

ADD memcached.conf /etc/supervisor/conf.d/memcached.conf

# Expose port 11211.
EXPOSE 11211

# Set start script as default command.
CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/conf.d/memcached.conf"]
