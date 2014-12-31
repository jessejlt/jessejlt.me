FROM dockerfile/nginx

RUN rm /etc/nginx/sites-available/default
ADD home /etc/nginx/sites-available/default

ADD site/ /var/www/home
RUN chown -R www-data /var/www/home

RUN ln -s /dev/stdout /var/log/nginx/access.log
RUN ln -s /dev/stderr /var/log/nginx/error.log
