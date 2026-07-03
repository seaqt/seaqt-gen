FROM miqt-docker/genbindings-base

ENV QT6_VERSION=6.8.3
ENV QT_ARCH=gcc_64

RUN \
    /aqtinstall/bin/aqt install-qt linux desktop ${QT6_VERSION} linux_${QT_ARCH} \
        -m qtmultimedia qtpdf qtpositioning qtscxml qtwebchannel qtwebengine qtwebsockets \
        --outputdir /opt/Qt && \
    rm -rf /root/.cache

ENV PATH="/opt/Qt/${QT6_VERSION}/${QT_ARCH}/bin:$PATH"
ENV LD_LIBRARY_PATH=/opt/Qt/${QT6_VERSION}/${QT_ARCH}/lib
ENV QT_PLUGIN_PATH=/opt/Qt/${QT6_VERSION}/${QT_ARCH}/plugins
ENV PKG_CONFIG_PATH=/opt/Qt/${QT6_VERSION}/${QT_ARCH}/lib/pkgconfig:/usr/local/lib/pkgconfig
ENV QT6_INCLUDE=/opt/Qt/${QT6_VERSION}/${QT_ARCH}/include
