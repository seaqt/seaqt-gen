FROM debian:bookworm

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -qyy \
        golang-go \
        qt6-base-dev \
        qt6-base-private-dev \
        qt6-charts-dev \
        qt6-declarative-dev \
        qt6-multimedia-dev \
        qt6-svg-dev \
        qt6-webengine-dev \
        libqscintilla2-qt6-dev \
        clang \
        git \
        ca-certificates \
        pkg-config \
        build-essential && \
    apt-get clean

