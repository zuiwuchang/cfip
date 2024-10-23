# cfip

測試以指定 ip 連接 cloudflare 的速度，從中找到連接最快的 ip

# docker

```
services:
  cfip:
    image: king011/cfip:v0.0.1
    restart: always
    volumes:
      - ./conf/cfip:/data:ro
    command: ["cfip", "-conf", "/data/v4.jsonnet"]
```