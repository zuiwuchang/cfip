FROM alpine:3.17.1

RUN set -eux;  \
    apk add --no-cache tzdata bash

Copy root /

RUN set -eux;  \
    chown root.root /usr/local/bin/cfip; \
    chmod a+x /usr/local/bin/cfip; \
    chown root.root  /usr/local/bin/docker-entrypoint.sh; \
    chmod a+x /usr/local/bin/docker-entrypoint.sh;

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["default-command"]