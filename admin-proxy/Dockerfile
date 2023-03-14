FROM node:18-bullseye

RUN apt update && apt upgrade -y

RUN mkdir -p /var/www/labyrinth/admin-proxy

WORKDIR /var/www/labyrinth/admin-proxy
ADD admin-proxy .

RUN npm install

EXPOSE 3000

ENTRYPOINT ["node", "app.js"]