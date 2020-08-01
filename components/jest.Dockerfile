FROM docker.io/node:10.15.3-alpine

WORKDIR /var/usr/app

COPY package.json ./

RUN npm install --registry https://npm.ponglehub.co.uk --strict-ssl false

COPY ./ ./

RUN npm run lint

ENTRYPOINT ["npm"]
CMD [ "start", "--silent" ]