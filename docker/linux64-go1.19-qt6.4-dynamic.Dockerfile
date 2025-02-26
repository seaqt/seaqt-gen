FROM debian:bookworm

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -qyy golang-go qt6-base-dev qt6-base-private-dev build-essential git && \
    apt-get clean

