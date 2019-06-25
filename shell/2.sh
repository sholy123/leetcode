#!/usr/bin/env bash
set -e
#name="sholy"
#echo "hello ${name}"
#
#your_name="tom"
#echo $your_name
#your_name="alibaba"
#echo $your_name
#echo ${#your_name}
test=(1 2 3)
#echo ${test[0]}
#for name in ${test} $name ; do
#    printf "$name hello \n"
#    printf "hello workd \n"
#done
test[0]="A"
test[1]="B"
test[2]="C"
num=0
while ((num<=3))
do
    echo ${test[${num}]}
    echo $num
    let num++
done
