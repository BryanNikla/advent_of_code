#!/bin/bash

# Run the application
PS3='Please enter your choice: '
options=("Go" "Javascript" "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Go")
            clear
            go run .
            break
            ;;
        "Javascript")
            clear
            node index
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done