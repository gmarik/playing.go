#!/usr/bin/env bash

set -e

GOROOT=${GOROOT:-/usr/local/Cellar/go/HEAD}

eval $(gomake --no-print-directory -f ${GOROOT}/src/Make.inc go-env)

if [ -z "$O" ]; then
	echo 'missing $O - maybe no Make.$GOARCH?' 1>&2
	exit 1
fi

rm -f *.$O

for i in \
  fib.go\
  main.go\
; do
	$GC $i
done

$LD main.$O && ./$O.out
