FROM node:16.0.0-alpine3.13 AS builder

WORKDIR /app

RUN apk add --no-cache git bash

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:1.20.0-alpine AS produciton

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/build/ ./

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY nginx/default.conf /etc/nginx/conf.d/
