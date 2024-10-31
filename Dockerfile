FROM php:8.0-apache

WORKDIR /var/www/html

COPY public_html/ /var/www/html/
COPY cmd/generate_pdf/generate_pdf /usr/local/bin/generate_pdf

EXPOSE 80
