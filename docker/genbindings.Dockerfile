FROM debian:bookworm

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install --no-install-recommends -qyy \
        golang-go \
        qtbase5-dev \
        qtbase5-private-dev \
        qtmultimedia5-dev \
        qtscript5-dev \
        libqt5svg5-dev \
        libqt5webkit5-dev \
        qtwebengine5-dev \
        qt6-base-dev \
        qt6-base-private-dev \
        qt6-charts-dev \
        qt6-declarative-dev \
        qt6-multimedia-dev \
        qt6-svg-dev \
        qt6-webengine-dev \
        libqscintilla2-qt5-dev \
        libqscintilla2-qt6-dev \
        clang \
        git \
        ca-certificates \
        pkg-config \
        build-essential && \
    apt-get clean

ENV GOFLAGS=-buildvcs=false
