# syntax=docker/dockerfile:experimental
FROM alpine AS deps

RUN apk add jq

COPY package*.json /tmp

RUN jq '.version="0.0.0"' < /tmp/package.json > /tmp/deps.json
RUN jq '.version="0.0.0"' < /tmp/package-lock.json > /tmp/deps-lock.json

FROM node:12.14.0-alpine

WORKDIR /var/app/src

COPY --from=deps /tmp/deps.json ./package.json
COPY --from=deps /tmp/deps-lock.json ./package-lock.json

RUN --mount=type=secret,id=npm_token,dst=/root/.npmrc npm ci --strict-ssl false

COPY . .

RUN npm run lint
RUN npm test
RUN npm version patch
RUN --mount=type=secret,id=npm_token,dst=/root/.npmrc npm publish --registry https://npm.ponglehub.co.uk --strict-ssl false