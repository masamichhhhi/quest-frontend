#!/bin/sh

set -e

host="$1"
shift
chm="$@"

# until mysqladmin ping -h -u root -ppassword $host -s; do
#   echo $host
#   echo 'Waiting for mysql'
#   sleep 1
# done

until mysql -h $host -u root -ppassword &> /dev/null
do
  echo "."
  sleep 1
done


echo "Mysql is up!"
exec $cmd