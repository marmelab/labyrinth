# This image is used for development purposes
FROM node:18-bullseye as builder

RUN apt update && apt upgrade -y

RUN mkdir -p /var/www/labyrinth/{admin,mobile}

WORKDIR /var/www/labyrinth/admin
ADD admin .
RUN npm install
RUN npm run build

WORKDIR /var/www/labyrinth/mobile
ADD mobile .
RUN npm install
RUN npm run build

FROM nginx:alpine

RUN apk update

RUN mkdir -p /var/www/labyrinth/{admin,mobile,webapp}

COPY                    proxy/production/nginx.conf         /etc/nginx/nginx.conf

COPY --from=builder     /var/www/labyrinth/admin/dist       /var/www/labyrinth/admin
COPY --from=builder     /var/www/labyrinth/mobile/dist      /var/www/labyrinth/mobile

COPY                    webapp/public/images                /var/www/labyrinth/webapp/images
COPY                    webapp/public/styles                /var/www/labyrinth/webapp/styles

EXPOSE 80
EXPOSE 443
