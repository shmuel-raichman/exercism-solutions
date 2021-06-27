#!/usr/bin/env bash
# B"H

main () {

    statementPartOne="One for "
    statementPartTwo=", one for me."

    name="you"
    givenName=$1

    if [ "x$givenName" != "x" ]
    then 
        name=$givenName
    fi

    echo "$statementPartOne$name$statementPartTwo"
}

# call main with all of the positional arguments
main "$@"
