# GOAL
# To build a development environment managed entirely within docker.
# Golang + Docker as an alternative to python (virtualenv) and vagrant.

# Concerns: host IDE is going to eat shit when searching for dependencies locally.

FROM golang

# Copy everything instead of volume
# ADD . /src/tryhard-platform

RUN mkdir -p /src/tryhard-platform
VOLUME /src/tryhard-platform

ENV GOPATH /go:/
ENV GOBIN /go/bin
ENV PATH $PATH:/go/bin

WORKDIR /src/tryhard-platform
RUN go get github.com/tockins/realize

# ENTRYPOINT ["realize", "start"]
CMD bash -c "realize start > logs/docker.log 2>&1"

# simple entry point for continuous server
# ENTRYPOINT ["tail", "-f", "/dev/null"]

# Helpful Commands
# docker exec -t -i 50f331760ba7 /bin/bash
# docker remove $(docker ps -aq)
# docker rmi $(docker images -q)
# docker start -a -i `docker ps -q -l`