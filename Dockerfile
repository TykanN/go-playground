FROM golang:1.16

RUN mkdir /app
WORKDIR /app
COPY . .

EXPOSE 5000

CMD [ "go", "run", "main.go" ]