FROM faddat/archlinux


COPY . /sentinel

RUN pacman -Syyu --noconfirm go base-devel protoc

RUN cd /sentinel && \
      make all
