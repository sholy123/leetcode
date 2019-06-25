#!/usr/bin/env bash
shuzu=(val1 val2 val3)

shuzu[0]=1
shuzu[1]=2
shuzu[2]=3
num=0
while ((num<3))
do
    echo ${shuzu[${num}]}
    echo $num
    let num++
done