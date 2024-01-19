FROM scratch

COPY sleeper /go/bin/sleeper
ENTRYPOINT ["/go/bin/sleeper"]
