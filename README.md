# caddy 

## build caddy v2
```bash
docker build --no-cache -t dukangxu/caddy:v2.x.x https://github.com/dukangxu/dockerfile.git\#caddy
```

* 默认添加插件[exec](https://github.com/abiosoft/caddy-exec),[cloudflare](https://github.com/caddy-dns/cloudflare),[realip](https://github.com/kirsch33/realip)

### build caddy v2 with hugo
```bash
curl https://raw.githubusercontent.com/dukangxu/dockerfile/caddy/Dockerfile_hugo|docker build --no-cache -t dukangxu/caddy-hugo:v2.x.x -
```

若出现以下错误，检查复制时是否自动添加了斜杠 ```\```
```
curl: option --no-cache: is unknown
curl: try 'curl --help' or 'curl --manual' for more information
```