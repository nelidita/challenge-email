FROM node:18.8.0-alpine

WORKDIR "/email-client"

COPY ./package*.json ./

RUN npm install

COPY . .

ARG VITE_API_URL=http://localhost:8080/
ENV VITE_API_URL=${VITE_API_URL}

RUN npm run build

CMD ["npx","serve","dist"]