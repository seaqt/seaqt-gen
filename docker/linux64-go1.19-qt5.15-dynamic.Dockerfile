FROM debian:bookworm

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -qyy golang-go qtbase5-dev qtbase5-private-dev qtscript5-dev qtmultimedia5-dev libqt5svg5-dev build-essential git && \
    apt-get clean
