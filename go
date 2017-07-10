#!/bin/bash

SSHCONFIG="$HOME/.ssh/config"
LIST=`grep -i "^Host" "$SSHCONFIG"| grep -v \* | awk '{ print $2 }' | sort`

if [ "$1" = "--list" -o "$1" = "-l" ]; then
  echo "$LIST"
  exit
fi

function VRFY_INPUT(){
  INPUT=$1
  if echo $LIST | grep $INPUT; then
    ssh $INPUT
  else
    echo "Fail"
  fi
}

echo -e "\nPlease choose a server # to access.\n"
PS3="
Server #: "
select SERVER in $LIST Exit ; do
  case $SERVER in
    Exit )
      exit 0
      ;;
    * )
      if [ -n "$SERVER" ]; then
        ssh $SERVER
      else
        echo "Fail"
      fi
      break
      ;;
  esac
done
