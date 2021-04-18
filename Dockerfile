FROM faddat/archlinux


COPY . /sentinel

RUN pacman -Syyu --noconfirm go base-devel protobuf git

RUN cd /sentinel && \
      make
