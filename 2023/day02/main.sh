#!/usr/bin/env bash

# Part 1

# 12 red cubes, 13 green cubes, and 14 blue cubes

# sum=0
# while read -r line; do
# 	game="$(echo "$line" | grep -oP '(?<=Game )[0-9]+')"
#
# 	possible=true
# 	for round in $(echo "$line" | cut -d: -f2 | tr -d "[:space:]" | tr -s ';' '\n'); do
# 		red=0
# 		green=0
# 		blue=0
#
# 		for res in $(echo "$round" | tr -s ',' '\n'); do
# 			count="$(echo "$res" | grep -oP '[0-9]+')"
# 			colour="$(echo "$res" | grep -oP 'red|green|blue')"
#
# 			if [ "$colour" = "red" ]; then red=$((red + count)); fi
# 			if [ "$colour" = "blue" ]; then blue=$((blue + count)); fi
# 			if [ "$colour" = "green" ]; then green=$((green + count)); fi
# 		done
#
# 		if ((red > 12 || green > 13 || blue > 14)); then
# 			possible=false
# 			break
# 		fi
# 	done
#
# 	if $possible; then
# 		sum=$((sum + game))
# 	fi
# done
#
# echo "$sum"

# sum=0
# while read -r line; do
# 	red_max=0
# 	green_max=0
# 	blue_max=0
#
# 	for round in $(echo "$line" | cut -d: -f2 | tr -d "[:space:]" | tr -s ';' '\n'); do
# 		for res in $(echo "$round" | tr -s ',' '\n'); do
# 			count="$(echo "$res" | grep -oP '[0-9]+')"
# 			colour="$(echo "$res" | grep -oP 'red|green|blue')"
#
# 			if [ "$colour" = "red" ] && ((count > red_max)); then
# 				red_max=$((count))
# 			fi
#
# 			if [ "$colour" = "blue" ] && ((count > blue_max)); then
# 				blue_max=$((count))
# 			fi
#
# 			if [ "$colour" = "green" ] && ((count > green_max)); then
# 				green_max=$((count))
# 			fi
#
# 		done
# 	done
#
# 	sum=$((sum + blue_max * red_max * green_max))
# done
#
# echo "$sum"

# Attempt 2

sum=0
while read -r line; do
	sum=$((
		sum + \
			$(echo "$line" | grep -oP "\d+ blue" | grep -oP "\d+" | sort -n | tail -n1) \
			* $(echo "$line" | grep -oP "\d+ red" | grep -oP "\d+" | sort -n | tail -n1) \
			* $(echo "$line" | grep -oP "\d+ green" | grep -oP "\d+" | sort -n | tail -n1) \
		))
done

echo "$sum"
