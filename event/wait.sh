#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

until mysqladmin ping -h $host -u root -ppassword -s; do
  echo 'Waiting for mysql'
  sleep 1
done

echo "Mysql is up!"
exec $cmd