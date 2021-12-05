#!/usr/bin/env bash

loc=$1

if [ -z "$loc" ]; then
    echo 'provide directory to generate day'
    exit 1
fi

mkdir -p $PWD/$loc

cp $PWD/templates/go/main.go $PWD/$loc
cp $PWD/templates/go/main_test.go $PWD/$loc
cp $PWD/templates/go/input.txt $PWD/$loc