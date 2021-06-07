FROM nginx:1.20.0-alpine

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY default.conf /etc/nginx/conf.d/

RUN mkdir /etc/nginx/virtual-hosts

COPY virtual-hosts/ /etc/nginx/virtual-hosts/
