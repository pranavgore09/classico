FROM node:7.5.0

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app
COPY . /usr/src/app

RUN cd /usr/src/app
RUN npm install


EXPOSE 50051
CMD ["npm", "start"]
