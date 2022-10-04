#!/usr/bin/env bash
make cover
curl -Os https://uploader.codecov.io/latest/macos/codecov
chmod 777 codecov

echo ${CODECOV}
./codecov -t 339562ba-7a26-4b07-9dea-9f1aef74eefe