FROM node:lts-alpine

# Enough for demo project but in case of prod env, prefer Nginx or Apache
RUN npm install -g http-server
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

CMD [ "http-server", "dist" ]