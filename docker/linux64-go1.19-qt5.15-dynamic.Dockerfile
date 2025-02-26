FROM debian:bookworm

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -qyy golang-go qtbase5-dev qtbase5-private-dev build-essential git && \
    apt-get clean
