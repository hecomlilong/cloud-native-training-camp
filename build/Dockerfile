FROM centos
COPY http-server /
COPY config/mod8.ini /config.ini
ENTRYPOINT ["/http-server", "--config","file:///config.ini","--v","2","--logtostderr","true"]
