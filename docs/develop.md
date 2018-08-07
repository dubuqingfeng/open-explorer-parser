## socket5 转 http 代理

利用 privoxy, cow, polipo

以下方式是用 privoxy

```
brew install privoxy
vim /usr/local/etc/privoxy/config

修改为以下内容：
forward-socks5t   /               127.0.0.1:1080 .
listen-address  127.0.0.1:8118

启动：
/usr/local/sbin/privoxy /usr/local/etc/privoxy/config

使用代理：
export http_proxy=127.0.0.1:8118
export https_proxy=127.0.0.1:8118
```

## 本地启动 redis 服务器

```
docker run --name example-redis -p 6379:6379 -d redis
```