FROM node:24.14.1-alpine

WORKDIR /app

COPY package*.json ./
RUN npm i

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
