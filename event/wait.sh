#!/bin/sh

set -e

host="$1"
shift
chm="$@"

until mysqladmin ping -h -u root -ppassword $host -s; do
  echo 'Waiting for mysql'
  sleep 1
done

echo "Mysql is up!"
exec $cmd