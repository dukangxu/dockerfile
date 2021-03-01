#
# Build stage
#
FROM golang:latest as builder

# args
ARG version="1.0.5"
# add plugin import paths here separated by commas
ARG plugins=""
ARG telemetry="false"

# build root
WORKDIR /build

# plugins
COPY plugins.go ./

# build & test
RUN apt-get update && apt-get install -y xz-utils upx git ca-certificates \
    && echo "module caddy\nrequire github.com/caddyserver/caddy v${version}" > go.mod \
    && go run plugins.go -plugins="${plugins}" -telemetry="${telemetry}" \
    && CGO_ENABLED=0 GOOS=linux go build \
    && upx caddy \
    && ./caddy -version \
    && go get -v github.com/abiosoft/parent \
    && rm -rf /var/lib/apt/lists/*

#
# Final image
#
FROM alpine:latest
LABEL maintainer "Kangxu Du <official@dukangxu.com>"

ARG version="1.0.5"
LABEL caddy_version="$version"

# Let's Encrypt Agreement
ENV ACME_AGREE="false"

# Telemetry Stats
ENV ENABLE_TELEMETRY="$enable_telemetry"

RUN apk add --no-cache \
    ca-certificates \
    git \
    mailcap \
    openssh-client \
    tzdata

# install caddy
COPY --from=builder /build/caddy /usr/bin/caddy

# validate install
RUN /usr/bin/caddy -version
RUN /usr/bin/caddy -plugins

EXPOSE 80 443 2015
VOLUME /root/.caddy /srv
WORKDIR /srv

COPY Caddyfile /etc/Caddyfile
COPY index.html /srv/index.html

# install process wrapper
COPY --from=builder /go/bin/parent /bin/parent

ENTRYPOINT ["/bin/parent", "caddy"]
CMD ["--conf", "/etc/Caddyfile", "--log", "stdout", "--agree=$ACME_AGREE"]
