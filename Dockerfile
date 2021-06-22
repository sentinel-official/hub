FROM faddat/archlinux as builder

ENV GOPATH /go
ENV PATH=$PATH:/go/bin

RUN pacman -Syyu --noconfirm git base-devel go gcc

COPY . .

RUN go mod tidy && \
    go mod download && \
        go build ./...

FROM faddat/archlinux

COPY --from=build /go/bin/sentinelhub /usr/bin/sentinelhub
