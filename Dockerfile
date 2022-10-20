FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . ./

COPY wait-for-it.sh /usr/bin/wait-for-it.sh
RUN chmod +x /usr/bin/wait-for-it.sh

RUN go build -o /Baetyl

EXPOSE 80

ENTRYPOINT /bin/bash /usr/bin/wait-for-it.sh -t 60 $POSTGRES_HOST:$POSTGRES_PORT -s -- /Baetyl