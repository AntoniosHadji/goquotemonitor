#!/usr/bin/env bash
for f in *.sql;
do
  docker cp "$f" postgres15-server:/root/. ;
done
