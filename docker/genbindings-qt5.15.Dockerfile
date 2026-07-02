FROM miqt-docker/genbindings-base

ENV QT5_VERSION=5.15.8
ENV QT_ARCH=gcc_64

# Unlike the other images, we install Qt using the native packages since
# aqt only has a buggy version (5.15.2)

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install --no-install-recommends -qyy \
        qtbase5-dev \
        qtbase5-private-dev \
        qtmultimedia5-dev \
        qtpdf5-dev \
        qtpositioning5-dev \
        qtscript5-dev \
        qttools5-dev \
        libqt5opengl5-dev \
        libqt5svg5-dev \
        libqt5webkit5-dev \
        libqt5websockets5-dev \
        libqwt-qt5-dev \
        qtwebengine5-dev && \
    apt-get clean

ENV QT5_INCLUDE=/usr/include/x86_64-linux-gnu/qt5
