#!/usr/bin/env bash
# B"H



main () {

    name="${1:-you}"
    echo "One for $name, one for me."
}

# call main with all of the positional arguments
main "$@"
