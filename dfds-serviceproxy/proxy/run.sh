#!/bin/sh
sed -i -e "s|REPLACE_PROXY_HOST|"$ADDR"|g" /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh nginx -g "daemon off;"