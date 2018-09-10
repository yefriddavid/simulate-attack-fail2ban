FROM nginx:1.15.2


COPY ./config/default.conf /etc/nginx/conf.d/default.conf

sed -i -e "s/server-ip/80/g" /etc/nginx/conf.d/default.conf
