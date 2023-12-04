#!/usr/bin/env bash

# Part 1

# sum=0
# while read -r line; do
# 	sum=$((sum + $(comm -12 \
# 		<(echo "$line" | cut -d: -f2 | cut -d"|" -f1 | grep -oP "\d+" | sort) \
# 		<(echo "$line" | cut -d: -f2 | cut -d"|" -f2 | grep -oP "\d+" | sort) |
# 		wc -l | xargs -I{} echo "2^({}-1)" | bc)))
# done
# echo "$sum"

# Part 2

winning=(0 0)
card=1
while read -r line; do
	win=$(comm -12 \
		<(echo "$line" | cut -d: -f2 | cut -d"|" -f1 | grep -oP "\d+" | sort) \
		<(echo "$line" | cut -d: -f2 | cut -d"|" -f2 | grep -oP "\d+" | sort) |
		wc -l)

	prev="${winning[card]:-0}"

	if [[ "$prev" = 0 ]]; then
		winning[card]=0
	fi

	for ((i = 1; i <= win; i++)); do
		winning[card + i]=$(("${winning[card + i]}" + prev + 1))
	done

	card=$((card + 1))
done

echo "${winning[*]} $((${#winning[@]} - 1))" | grep -oP "\d+" | paste -sd+ | bc
