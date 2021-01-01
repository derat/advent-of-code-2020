#!/bin/sh -e

die() {
  echo "$1" >&2
  exit 2
}

usage() {
  prog=$(basename "$0")
  cat <<EOF >&2
Usage:
  $prog <YEAR> <DAY>   Print (and init) dir for specified year and day.
  $prog <DAY>          Print dir for specified day in current dir's year.
  $prog input          Print input for current dir.
  $prog lib            Print library directory.
  $prog next           Print dir for day after current dir.
  $prog prev           Print dir for day before current dir.
  $prog run            Run code in current dir.
  $prog stdin          Run code in current dir with input from stdin.
  $prog today          Print dir for today.
  $prog web            Open webpage for current dir.
  $prog                Print repo directory.
EOF
  exit 2
}

script_dir="$(dirname "$(realpath -s "$0")")"

# Figure out if we're already in a year/day or year directory.
cur_dir=$(pwd)
cur_year=
cur_day=
case "$cur_dir" in
  ${script_dir}/20[1-9][0-9]/[0-2][0-9])
    cur_year="$(basename "$(dirname "$cur_dir")")"
    cur_day="$(basename "$cur_dir")"
    break
    ;;
  ${script_dir}/20[1-9][0-9])
    cur_year="$(basename "$cur_dir")"
    break
    ;;
esac

# Dies with an error if not already in a year/day directory.
check_in_day_dir() {
  if [ -z "$cur_year" ] || [ -z "$cur_day" ]; then
    die "Must be in year/day directory"
  fi
}

year=
day=

[ $# -eq 0 ] && exec echo "$script_dir"

case "$1" in
  -h|--help)
    usage
    ;;
  input)
    check_in_day_dir
    exec cat "$HOME/.cache/advent-of-code/$(printf "%d/%d" $cur_year $cur_day)"
    ;;
  lib)
    exec echo "${script_dir}/lib"
    ;;
  next)
    check_in_day_dir
    year=$cur_year
    day=$(($cur_day + 1))
    ;;
  prev)
    check_in_day_dir
    year=$cur_year
    day=$(($cur_day - 1))
    ;;
  run)
    check_in_day_dir
    exec go run main.go
    ;;
  stdin)
    check_in_day_dir
    exec go run main.go -
    ;;
  today)
    [ $(date +%m) -eq 12 ] || die "Not in December"
    year=$(date +%Y)
    day=$(date +%d)
    ;;
  web)
    check_in_day_dir
    exec xdg-open "$(printf "https://adventofcode.com/%d/day/%d" $cur_year $cur_day)"
    ;;
  *)
    if [ $# -eq 1 ]; then
      [ -n "$cur_year" ] || die "Must be in year or year/day directory"
      year="$cur_year"
      day="$1"
    elif [ $# -eq 2 ]; then
      year="$1"
      day="$2"
    else
      usage
    fi
    ;;
esac

# Validate the year and day that we're using.
if ! echo "$year" | grep -E -q '^[0-9]{4}$' || [ "$year" -lt 2015 ]; then
  die "Year '${year}' not in range [2015, ...]"
fi
if ! echo "$day" | grep -E -q '^[0-9][0-9]?$' || [ "$day" -lt 1 ] || [ "$day" -gt 25 ]; then
  die "Day '${day}' not in range [1, 25]"
fi

# Remove zero-padding.
year=$(printf "%d" "$year")
day=$(printf "%d" "$day")

dir="${script_dir}/$(printf "%04d/%02d" "$year" "$day")"

if [ ! -e "$dir" ]; then
  mkdir -p "$dir"
  cat <<EOF >"${dir}/main.go"
package main

import (
	"fmt"

	"github.com/derat/advent-of-code/lib"
)

func main() {
	for _, ln := range lib.InputLines("${year}/${day}") {
		fmt.Println(ln)
	}
}
EOF
fi

echo "$dir"
