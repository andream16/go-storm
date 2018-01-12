FROM node:8.9-alpine as angular-built

# Gyp dependencies
RUN apk add --no-cache make gcc g++ python2 git && \
  npm install --production --silent && \
  apk del make gcc g++ python2
RUN ln -s /usr/bin/python2.7 /usr/bin/python

# Angular
RUN npm i -g @angular/cli@1.6.3
RUN git clone https://github.com/AndreaM16/price-probe-ui-v2.git
WORKDIR ./price-probe-ui-v2
RUN npm i
RUN npm rebuild node-sass --force
RUN ng build

# Nginx
FROM nginx:alpine
COPY --from=angular-built /price-probe-ui-v2/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 4200
CMD [ "nginx", "-g", "daemon off;" ]