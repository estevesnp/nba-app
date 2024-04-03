FROM node:latest

WORKDIR /app

COPY ./react-app /app

COPY ./docker.env /app/.env.local

RUN npm install

RUN npm run build

EXPOSE 3000

CMD ["npm", "start", "--", "--port", "3000"]