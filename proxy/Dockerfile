FROM nginx:1.20.0-alpine

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY virtual-hosts/* /etc/nginx/conf.d/
