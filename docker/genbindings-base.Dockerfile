# Base image for genbindings: Go + dev tools + aqtinstall (no Qt)
# Used as the parent for version-specific Dockerfiles.

FROM debian:bookworm

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG=C.UTF-8
ENV LC_ALL=C.UTF-8

# Layer: System deps + aqtinstall + Go + Nimble

RUN \
    apt-get update && \
    apt-get install --no-install-recommends -qyy \
        libglib2.0-0 \
        libgl-dev \
        clang \
        git \
        ca-certificates \
        pkg-config \
        build-essential \
        python3 \
        python3-pip \
        python3-venv \
        wget \
        # for tests
        fontconfig \
        libegl1 \
        libfreetype6 \
        libxkbcommon0 \
        libxkbfile1 \
        libnspr4 \
        libxtst6 \
        libxcomposite1 \
        libxdamage1 \
        libnss3 \
        libxi6 \
        libxrandr2 \
        libasound2 \
        libdbus-1-3 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# aqtinstall for downloading Qt versions
RUN \
    python3 -m venv /aqtinstall && \
    /aqtinstall/bin/pip install aqtinstall && \
    rm -rf /root/.cache

# Go (for running tests; TODO move to separate image)
RUN \
    wget 'https://go.dev/dl/go1.24.13.linux-amd64.tar.gz' && \
    tar -C /usr/local -xzf go1.24.13.linux-amd64.tar.gz && \
    rm go1.24.13.linux-amd64.tar.gz

ENV PATH=/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV GOFLAGS=-buildvcs=false

# Nimble (for running; TODO move to separate image)
RUN \
    wget https://github.com/nim-lang/nimble/releases/download/v0.18.2/nimble-linux_x64.tar.gz && \
    tar xvf nimble-linux_x64.tar.gz && \
    mv nimble /usr/local/bin && \
    rm nimble-linux_x64.tar.gz
