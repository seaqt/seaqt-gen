FROM debian:bookworm

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install --no-install-recommends -qyy \
        golang-go \
        qtbase5-dev \
        qtbase5-private-dev \
        qtmultimedia5-dev \
        qtpdf5-dev \
        qtpositioning5-dev \
        qtscript5-dev \
        qttools5-dev \
        libqt5svg5-dev \
        libqt5webkit5-dev \
        qtwebengine5-dev \
        qt6-base-dev \
        qt6-base-private-dev \
        qt6-charts-dev \
        qt6-declarative-dev \
        qt6-multimedia-dev \
        qt6-pdf-dev \
        qt6-positioning-dev \
        qt6-tools-dev \
        qt6-svg-dev \
        qt6-webengine-dev \
        clang \
        git \
        ca-certificates \
        pkg-config \
        build-essential \
        wget && \
    apt-get clean && \
    wget https://github.com/nim-lang/nimble/releases/download/v0.18.2/nimble-linux_x64.tar.gz && \
    tar xvf nimble-linux_x64.tar.gz && \
    mv nimble /usr/local/bin

ENV GOFLAGS=-buildvcs=false
