# This image is used for development purposes
FROM node:18-bullseye

RUN apt update && apt upgrade -y

RUN mkdir -p /var/www/labyrinth

WORKDIR /var/www/labyrinth
ADD mobile .

RUN npm install

EXPOSE 5173

ENTRYPOINT ["npm", "run", "dev"]