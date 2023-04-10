FROM golang:1.19

LABEL maintainer="Ansh Jamwal"

RUN mkdir /adhd-server

WORKDIR /adhd-server

#COPY go.mod go.sum ./

#RUN go mod download

RUN apt-get update && apt-get install make

COPY . .

COPY .env .

RUN make build

EXPOSE 8000

CMD [ "./build/adhdserver" ]