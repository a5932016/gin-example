#!/bin/bash

# install go
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.24.0 -B
gvm use go1.24.0 --default