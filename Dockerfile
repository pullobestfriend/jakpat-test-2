FROM golang
WORKDIR /usr/src/app

COPY ./.bin/app ./bin/app
COPY .env .env
COPY config config

EXPOSE 8880

ENTRYPOINT [ "./bin/app" ]