#!/usr/bin/env bash
# B"H

main(){
    
    # Rename parameter.
    stringToRevers=$1
    # get string length.
    srtingLength=${#stringToRevers}

    # start from end of string and concat each character to new string.
    for ((i = $srtingLength - 1; i >= 0; i--));
        do
            reverse="$reverse${stringToRevers:$i:1}" 
        done

    echo "$reverse"
}

main "$@"