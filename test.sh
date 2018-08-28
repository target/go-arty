#!/usr/bin/env bash

echo "mode: set" > coverage.out
for Dir in $(find ./* -maxdepth 10 -type d );
do
        if ls $Dir/*.go &> /dev/null;
        then
            go test -coverprofile=profile.out $Dir
            if [ -f profile.out ]
            then
                cat profile.out | grep -v "mode: set" >> coverage.out
            fi
fi
done
