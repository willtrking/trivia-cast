FROM node:latest

WORKDIR /src
COPY package*.json /src/
RUN npm install

COPY pages/ /src/pages

CMD npm run dev
