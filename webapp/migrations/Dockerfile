# Base
FROM php:8.1-cli as base

RUN apt update && apt upgrade -y

RUN apt install -y zlib1g-dev g++ git libicu-dev zip libpq-dev libzip-dev zip netcat \
    && docker-php-ext-install intl opcache pdo pdo_pgsql \
    && pecl install apcu \
    && docker-php-ext-enable apcu \
    && docker-php-ext-configure zip \
    && docker-php-ext-install zip

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN mkdir -p /var/www/labyrinth

WORKDIR /var/www/labyrinth

ADD webapp  .

RUN /usr/local/bin/composer install

ENTRYPOINT [ "./run-migrations.sh" ]
