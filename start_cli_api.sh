#!/usr/bin/env sh

/app/nfeloader &
P1=$!
/app/nfeloader-api &
P2=$!
wait $P1 $P2

