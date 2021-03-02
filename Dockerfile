FROM alpine:latest
MAINTAINER Kangxu Du <im@dukangxu.com>

WORKDIR /

RUN apk add --no-cache curl && \
    FRP_VERSION=$(curl -sX GET "https://api.github.com/repos/fatedier/frp/releases/latest" | awk '/tag_name/{print $4}' FS='[""]' | sed 's/^.//g') && \
    wget --no-check-certificate https://github.com/fatedier/frp/releases/download/v${FRP_VERSION}/frp_${FRP_VERSION}_linux_amd64.tar.gz && \
    tar xzf frp_${FRP_VERSION}_linux_amd64.tar.gz && \
    mv frp_${FRP_VERSION}_linux_amd64 /frp && \
    rm -rf *.tar.gz frp/LICENSE frp/*_full.ini frp/systemd

VOLUME /frp

CMD ["/frp/frpc", "-c", "/frp/frpc.ini"]
