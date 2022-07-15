#!/bin/bash

if [ $1 == "bonus" ]; then
    while IFS=, read -r id city name phone salary address; do
        if [ $2 = $id ]; then
            bonus=`expr $salary \* 5 \/ 100`
            echo "$name will get \$$bonus bonus"
            break
        fi
    done < employee.csv
elif [ $1 == "city" ]; then
    while IFS=, read -r id city name phone salary address; do
        if [ $2 = $city ]; then
            echo "Customer Name: $name"
            echo "Mobile No: $phone"
        fi
    done < employee.csv
else
    echo "command not found"
fi
