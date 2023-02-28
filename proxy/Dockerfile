FROM nginx:alpine

RUN apk update

RUN mkdir -p /var/www/labyrinth/public

COPY proxy/nginx.conf /etc/nginx/nginx.conf

COPY webapp/public/images /var/www/labyrinth/public/images
COPY webapp/public/styles /var/www/labyrinth/public/styles

EXPOSE 80
EXPOSE 443
