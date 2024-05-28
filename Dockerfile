FROM golang:1.21.3

WORKDIR /app

COPY . .

# EXPOSE 8080

CMD [ "go", "run", "main.go" ]