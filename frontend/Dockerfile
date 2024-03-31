FROM node:latest

WORKDIR /app

COPY ./react-app /app

RUN npm install

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]