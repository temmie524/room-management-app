FROM --platform=linux/x86_64 node:18.12.1-alpine

WORKDIR /usr/src/app

# copy package.json and package-lock.json
COPY . .

# install dependencies
RUN yarn install

# start app
CMD ["yarn", "start"]