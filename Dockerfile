FROM golang:1.19

RUN mkdir /adhd-server

WORKDIR /adhd-server

#COPY go.mod go.sum ./

#RUN go mod download

RUN apt-get update && apt-get install make

COPY .env /adhd-server

COPY . .

RUN make build

EXPOSE 8000

CMD [ "./build/adhdserver" ]