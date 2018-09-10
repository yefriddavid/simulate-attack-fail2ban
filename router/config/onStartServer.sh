#!/bin/bash

echo "------------START-------------"
#sed -i "s|server-ip|80|g" /app/test.conf

cd /app
#chmod 777 test.conf
#sed -i -e "s|server-ip|80|g" /app/test1.conf



sed "11ilisten   80" /app/test.conf
  #listen       80;
ls
pwd
dir

cat /app/test.conf



echo "------------END-------------"
