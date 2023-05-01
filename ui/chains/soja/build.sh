#!/bin/bash

killall sojanoded

rm $(which sojanoded) 2> /dev/null || echo sojanoded not install yet ...

rm -rf ~/.sojanoded

cd ../../../ && make install 