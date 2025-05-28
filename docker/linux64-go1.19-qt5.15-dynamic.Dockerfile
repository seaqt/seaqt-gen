FROM debian:bookworm

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -qyy \
        golang-go \
        qtbase5-dev \
        qtbase5-private-dev \
        qtmultimedia5-dev \
        qtpdf5-dev \
        qtscript5-dev \
        libqt5svg5-dev \
        libqt5webkit5-dev \
        qtwebengine5-dev \
        clang \
        git \
        ca-certificates \
        pkg-config \
        build-essential && \
    apt-get clean
