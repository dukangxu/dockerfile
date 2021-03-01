FROM python:3.7-alpine
LABEL maintainer="Kangxu Du<im@dukangxu.com>"

RUN set -ex \
	&& apk add --no-cache tar libsodium-dev openssl \
	&& wget -O /tmp/shadowsocksr-3.2.2.tar.gz https://github.com/dukangxu/shadowsocksr/archive/3.2.2.tar.gz \
	&& tar zxf /tmp/shadowsocksr-3.2.2.tar.gz -C /tmp \
	&& mv /tmp/shadowsocksr-3.2.2/shadowsocks /usr/local/ \
	&& rm -fr /tmp/shadowsocksr-3.2.2 \
	&& rm -f /tmp/shadowsocksr-3.2.2.tar.gz

ADD config_sample.json /etc/shadowsocks-r/config.json
VOLUME /etc/shadowsocks-r

CMD [ "/usr/local/shadowsocks/server.py", "-c", "/etc/shadowsocks-r/config.json" ]