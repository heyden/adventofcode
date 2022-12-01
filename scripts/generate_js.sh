#!/usr/bin/env bash

loc=$1

if [ -z "$loc" ]; then
    echo 'provide directory to generate day'
    exit 1
fi

rm -r $PWD/$loc
mkdir -p $PWD/$loc/js

cp $PWD/templates/js/main.js $PWD/$loc/js
cp $PWD/templates/js/main.test.js $PWD/$loc/js
cp $PWD/templates/js/input.txt $PWD/$loc