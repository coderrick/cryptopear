#!/bin/zsh
emulate -LR zsh # reset zsh options
export PATH=/usr/bin:/bin:/usr/sbin:/sbin

function countArguments() {
    echo "${#@}"
}

wordlist="one two three four five"

echo "normal substitution, no quotes:"
countArguments $wordlist
# -> 1

echo "substitution with quotes"
countArguments "$wordlist"
# -> 1