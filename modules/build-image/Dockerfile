FROM node:12.14.0-alpine

RUN apk add expect

WORKDIR /var/app/src

COPY ./ /usr/local/bin

RUN npm config set strict-ssl false

ENTRYPOINT [ "/usr/local/bin/entrypoint.sh" ]