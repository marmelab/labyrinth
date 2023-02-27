FROM nginx:alpine

RUN apk update

RUN mkdir -p /var/www/labyrinth/public

COPY proxy/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
EXPOSE 443
