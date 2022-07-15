#!/bin/bash

mkdir -p ./out

for filename in in/input*.txt; do
    index=${filename#"in/input"}
    index=${index%".txt"}
    python3 main.py < "$filename" > "out/output$index.txt"
done
