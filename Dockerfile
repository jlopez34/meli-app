FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build

CMD ./meli-api
EXPOSE 8080