#!/usr/bin/env bash

# Part 1

# sum=0
# while read -r line; do
# 	line=${line//[^0-9]/}
#
# 	first="$(echo "$line" | cut -c1)"
# 	last="$(echo "$line" | rev | cut -c1)"
#
# 	number="${first}${last}"
#
# 	sum=$((sum + number))
# done
#
# echo "$sum"

# Part 2

sum=0
while read -r line; do
	l=$(echo "$line" | perl -ne 'print "$1\n" while /(?=(\d|one|two|three|four|five|six|seven|eight|nine))/g')

	first=$(echo "$l" | head -n1)
	last=$(echo "$l" | tail -n1)

	number="${first}${last}"

	number="${number//one/1}"
	number="${number//two/2}"
	number="${number//three/3}"
	number="${number//four/4}"
	number="${number//five/5}"
	number="${number//six/6}"
	number="${number//seven/7}"
	number="${number//eight/8}"
	number="${number//nine/9}"

	sum=$((sum + number))
done

echo "$sum"
