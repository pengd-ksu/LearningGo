#!/usr/bin/env bash

find . -name "*.txt" | xargs rm -f
go run fetchToFile.go https://github.com/adonovan/gopl.io/
diff 0.txt 1.txt