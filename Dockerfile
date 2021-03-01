FROM caddy:builder AS builder

RUN xcaddy build \
    --with github.com/abiosoft/caddy-exec \
    --with github.com/caddy-dns/cloudflare \
    --with github.com/kirsch33/realip \
    && apk add upx \
    && upx /usr/bin/caddy

FROM caddy:latest

RUN apk add --no-cache git bash

COPY --from=builder /usr/bin/caddy /usr/bin/caddy