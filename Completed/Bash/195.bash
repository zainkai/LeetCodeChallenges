#!/bin/bash
# Read from the file file.txt and output the tenth line to stdout.

i=0
while read line;
do
    ((i++))
    if [[ $i -eq 10 ]]; then
        echo $line
    fi
    
done < "file.txt"
