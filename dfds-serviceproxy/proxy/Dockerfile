FROM nginx

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./run.sh /run.sh
RUN chmod +x run.sh

ENTRYPOINT [ "/run.sh" ]