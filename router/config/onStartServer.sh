#!/bin/bash

echo "------------START-------------"
#sed -i "s|server-ip|80|g" /app/test.conf

cd /app
chmod 777 test.conf
sed -i "s|server-ip|80|g" /app/test.conf
ls
pwd
dir

cat /app/test.conf



echo "------------END-------------"
