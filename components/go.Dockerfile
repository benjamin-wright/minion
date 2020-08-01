ARG BASE_IMAGE=scratch
FROM $BASE_IMAGE

ARG EXE_NAME

COPY ${EXE_NAME} /go/bin/go_binary

ENTRYPOINT ["/go/bin/go_binary"]