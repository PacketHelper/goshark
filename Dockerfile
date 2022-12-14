FROM golang:1.19

# install tshark
RUN echo "wireshark-common wireshark-common/install-setuid boolean true" | debconf-set-selections
RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get -y install wireshark

RUN apt-get update && apt-get install -y --allow-change-held-packages --force-yes tshark

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v main.go
RUN chmod a+x main

CMD ./main
