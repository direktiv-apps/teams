#!/bin/sh

docker build -t teams . && docker run -p 9191:8080 teams