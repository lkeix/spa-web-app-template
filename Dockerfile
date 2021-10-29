FROM node:16-alpine
MAINTAINER lkeix

WORKDIR /app

ADD . /app

RUN npm install -g npm@7.18.1

CMD ["sh", "./npm_start.sh"]
