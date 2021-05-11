FROM golang:latest as compile
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-w -s' -o server
ENTRYPOINT ["/build/server"]

# scratch stands for a pure empty image, we only copy the 'server' executable file into this empty image.
FROM scratch
# --from: copy from step 'compile',
# it aslo can be a number, sucn as 0,1..
# it also can be another docker image name and then copy from that image.
COPY --from=compile /build/server /
ENTRYPOINT ["/server"]