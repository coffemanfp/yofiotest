#!/bin/bash

cd migrations/postgresql || exit
go build -o ../../bin/postgresql

echo 'Starting migrations...'
../../bin/postgresql --with-examples

cd ../../ || exit
