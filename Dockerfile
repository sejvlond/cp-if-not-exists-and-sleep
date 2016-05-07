FROM scratch
MAINTAINER Ondřej Šejvl
COPY run /
ENTRYPOINT [ "/run" ]
