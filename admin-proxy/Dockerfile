FROM node:18-alpine

RUN apk update

RUN mkdir -p /var/www/labyrinth/admin-proxy

WORKDIR /var/www/labyrinth/admin-proxy
ADD admin-proxy .

RUN npm install

EXPOSE 3000

ENTRYPOINT ["node", "app.js"]