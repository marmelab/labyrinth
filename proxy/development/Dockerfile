FROM nginx:alpine

RUN apk update

RUN mkdir -p /var/www/labyrinth/{admin,mobile,webapp}

COPY proxy/development/nginx.conf   /etc/nginx/nginx.conf

COPY mobile/public                  /var/www/labyrinth/mobile
COPY webapp/public/images           /var/www/labyrinth/webapp/images
COPY webapp/public/styles           /var/www/labyrinth/webapp/styles

EXPOSE 80
EXPOSE 443
