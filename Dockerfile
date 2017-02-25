FROM golang:1.6

RUN mkdir -p /go/src/app
RUN mkdir -p /haste-data

WORKDIR /go/src/app
COPY . /go/src/app

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]

EXPOSE 8181

# Useful Commands

# Build
# docker build -t bland/brian-bland-me-node .

# Run
# docker run -d -p 80:8080 -t bland/brian-bland-me-node

# To view iptables mapping
# iptables -t nat -L -n

# Delete all containers
# docker rm $(docker ps -a -q)

# Delete all images
# docker rmi $(docker images -q)
