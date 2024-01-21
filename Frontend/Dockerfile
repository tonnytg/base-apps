FROM node:lts-slim as core
WORKDIR /app
COPY ./webapp/ .

RUN apt-get update && apt-get install -y curl wget telnet nginx certbot

RUN npm install

RUN npm install -g firebase-tools

RUN npm run build

#COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 3000:3000
EXPOSE 80:3080
EXPOSE 443:3443
EXPOSE 9005:9005

ENV PORT=3000

CMD ["npm", "start"]