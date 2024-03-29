#!/bin/bash

function setup() {
  echo -n "building....."

  rm -rf tmp
  mkdir -p tmp
  make clean &> /dev/null
  make &> /dev/null
  if (($? > 0)); then
    echo "error"
    exit 1
  fi
  cp build/main tmp/aoc
  echo "ok!"
  echo
}

function tests() {
  e=0
  while IFS= read -r line
  do
    IFS=' ' read -r problem expected_result <<< "$line"
    expected="tmp/$problem.expected"
    result="tmp/$problem.out"
    echo -n "testing $problem...."

    echo $expected_result > $expected
    case $problem in
      1.1|1.2|2.1|2.2|3.1|3.2|4.1|4.2)
        tmp/aoc $problem > $result
        ;;
      *)
        python3 main.py $problem > $result
        ;;
    esac

    if cmp -s $expected $result; then
      echo OK
    else
      ((e+=1))
      echo ERROR
      echo "    expected: $(cat $expected)"
      echo "    got:      $(cat $result)"
    fi
  done < "tests.txt"

  return $e
}

setup
tests

errors=$?
if (($errors > 0)); then
  echo
  echo "$errors errors found"
fi

exit $errors