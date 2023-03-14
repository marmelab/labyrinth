# This image is used for development purposes
FROM node:18-bullseye

RUN apt update && apt upgrade -y

RUN mkdir -p /var/www/admin

WORKDIR /var/www/admin
ADD admin .

RUN npm install

EXPOSE 5173

ENTRYPOINT ["npm", "run", "dev"]