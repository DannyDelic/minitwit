#! /bin/bash

# default ENV is dev
env=dev

while test $# -gt 0; do
  case "$1" in
    -env)
      shift
      if test $# -gt 0; then
        env=$1
      fi
      # shift
      ;;
    *)
    break
    ;;
  esac
done

cd ../../minitwit || exit
source .env
go build -o cmd/minitwit/minitwit cmd/minitwit/main.go
cmd/minitwit/minitwit -env "$env" &