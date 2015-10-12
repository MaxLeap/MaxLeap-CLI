#!/usr/bin/env bash
if ["$1"=="win"];then
    go build -o us/maxleap .
    go build -o cn/maxleap -ldflags "-X  main.region CN" .
else
    go build -o us/maxleap .
    go build -o cn/maxleap -ldflags "-X  main.region CN" .
fi