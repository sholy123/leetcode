#!/usr/bin/env bash

echo $0 > test.txt
echo $1
echo $2
echo $*
echo $* >> test.txt

i=0
while((i<=5))
do
    echo $i
    let i++
done