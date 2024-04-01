FROM node:latest

WORKDIR /app

COPY ./react-app /app

RUN echo "BACKEND_HOST=backend\nBACKEND_PORT=9000" > /app/.env.local

RUN npm install

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]