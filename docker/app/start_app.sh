#!/bin/sh

echo ">> Waiting for $MYSQL_SERVER to start"
while ! `nc -z $MYSQL_SERVER 3306`; do sleep 3; done
echo ">> $MYSQL_SERVER has started"

/app/main