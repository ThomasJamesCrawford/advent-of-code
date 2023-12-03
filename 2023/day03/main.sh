#!/usr/bin/env bash

# sum=0
#
# parts=()
#
# mapfile -t < <(cat -)
#
# for ((y = 0; y < ${#MAPFILE}; y++)); do
# 	line="${MAPFILE[$y]}"
#
# 	for ((x = 0; x < ${#line}; x++)); do
# 		char="${line:$x:1}"
#
# 		if [[ "$char" =~ ^[0-9]$ ]]; then
# 			left=""
# 			downleft=""
# 			upleft=""
# 			right=""
# 			downright=""
# 			upright=""
# 			above=""
# 			below=""
# 			if [[ "$x" -gt 0 ]]; then
# 				left="${MAPFILE[$y]:$x-1:1}"
#
# 				if [[ "$y" -lt ${#MAPFILE} ]]; then
# 					downleft="${MAPFILE[$y + 1]:$x-1:1}"
# 				fi
#
# 				if [[ "$y" -gt 0 ]]; then
# 					upleft="${MAPFILE[$y - 1]:$x-1:1}"
# 				fi
# 			fi
#
# 			if [[ "$x" -lt ${#line} ]]; then
# 				right="${MAPFILE[$y]:$x+1:1}"
#
# 				if [[ "$y" -lt ${#MAPFILE} ]]; then
# 					downright="${MAPFILE[$y + 1]:$x+1:1}"
# 				fi
#
# 				if [[ "$y" -gt 0 ]]; then
# 					upright="${MAPFILE[$y - 1]:$x+1:1}"
# 				fi
# 			fi
#
# 			if [[ "$y" -gt 0 ]]; then
# 				above="${MAPFILE[$y - 1]:$x:1}"
# 			fi
#
# 			if [[ "$y" -lt ${#MAPFILE} ]]; then
# 				below="${MAPFILE[$y + 1]:$x:1}"
# 			fi
#
# 			if [[ "$left$downleft$upleft$right$downright$upright$above$below" =~ ([^\.0-9]) ]]; then
# 				part="$char"
# 				for ((i = x + 1; i < ${#line}; i++)); do
# 					if [[ "${line:$i:1}" =~ [0-9] ]]; then
# 						part+="${line:$i:1}"
# 					else
# 						break
# 					fi
# 				done
#
# 				end=$((i - 1))
#
# 				for ((i = x - 1; i >= 0; i--)); do
# 					if [[ "${line:$i:1}" =~ [0-9] ]]; then
# 						part="${line:$i:1}$part"
# 					else
# 						break
# 					fi
# 				done
#
# 				# jump to the end of this part number
# 				x=$((end))
#
# 				echo "$part, y = $y, x = $x, matches=$matches"
#
# 				sum=$((sum + part))
# 				parts+=("$part")
# 			fi
# 		fi
# 	done
# done
#
# echo "$sum"

is_number() {
	if [[ "$1" =~ [0-9] ]]; then
		return 0
	else
		return 1
	fi
}

number_idx() {
	x_idx="$1"
	y_idx="$2"

	number_line="${MAPFILE[$y_idx]}"

	for ((i = x_idx - 1; i >= 0; i--)); do
		if ! [[ "${number_line:$i:1}" =~ [0-9] ]]; then
			break
		fi
	done

	echo "$((i + 1)),$y_idx"
}

sum=0

parts=()

mapfile -t < <(cat -)

for ((y = 0; y < ${#MAPFILE}; y++)); do
	line="${MAPFILE[$y]}"

	for ((x = 0; x < ${#line}; x++)); do
		char="${line:$x:1}"

		if [[ "$char" = "*" ]]; then
			echo "found gear"
			left=""
			downleft=""
			upleft=""
			right=""
			downright=""
			upright=""
			above=""
			below=""
			xy_idxs=()

			if [[ "$x" -gt 0 ]]; then
				left="${MAPFILE[$y]:$x-1:1}"

				if is_number "$left"; then
					y_idx=$((y))
					x_idx=$((x - 1))
					xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")

				fi

				if [[ "$y" -lt ${#MAPFILE} ]]; then
					downleft="${MAPFILE[$y + 1]:$x-1:1}"

					if is_number "$downleft"; then
						y_idx=$((y + 1))
						x_idx=$((x - 1))
						xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")
					fi
				fi

				if [[ "$y" -gt 0 ]]; then
					upleft="${MAPFILE[$y - 1]:$x-1:1}"
					if is_number "$upleft"; then
						y_idx=$((y - 1))
						x_idx=$((x - 1))
						xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")
					fi
				fi
			fi

			if [[ "$x" -lt ${#line} ]]; then
				right="${MAPFILE[$y]:$x+1:1}"

				if is_number "$right"; then
					y_idx=$((y))
					x_idx=$((x + 1))
					xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")
				fi

				if [[ "$y" -lt ${#MAPFILE} ]]; then
					downright="${MAPFILE[$y + 1]:$x+1:1}"
					if is_number "$downright"; then
						y_idx=$((y + 1))
						x_idx=$((x + 1))
						xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")
					fi
				fi

				if [[ "$y" -gt 0 ]]; then
					upright="${MAPFILE[$y - 1]:$x+1:1}"
					if is_number "$upright"; then
						y_idx=$((y - 1))
						x_idx=$((x + 1))
						xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")
					fi
				fi
			fi

			if [[ "$y" -gt 0 ]]; then
				above="${MAPFILE[$y - 1]:$x:1}"
				if is_number "$above"; then
					y_idx=$((y - 1))
					x_idx=$((x))
					xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")
				fi
			fi

			if [[ "$y" -lt ${#MAPFILE} ]]; then
				below="${MAPFILE[$y + 1]:$x:1}"
				if is_number "$below"; then
					y_idx=$((y + 1))
					x_idx=$((x))
					xy_idxs+=("$(number_idx "$x_idx" "$y_idx")")
				fi
			fi

			coords=()
			for n in $(echo "${xy_idxs[@]}" | grep -oE "[0-9]+,[0-9]+" | sort -u); do
				coords+=("$n")
			done

			if [[ "${#coords[@]}" = 2 ]]; then
				sum_part=1
				echo "${coords[@]}"
				for ((c = 0; c < ${#coords[@]}; c++)); do
					x_idx="${coords[$c]%,*}"
					y_idx="${coords[$c]#*,}"

					part="${MAPFILE[$y_idx]:$x_idx:1}"
					number_line="${MAPFILE[$y_idx]}"

					for ((i = x_idx + 1; i < ${#number_line}; i++)); do
						if [[ "${number_line:$i:1}" =~ [0-9] ]]; then
							part+="${number_line:$i:1}"
						else
							break
						fi
					done

					for ((i = x_idx - 1; i >= 0; i--)); do
						if [[ "${number_line:$i:1}" =~ [0-9] ]]; then
							part="${number_line:$i:1}$part"
						else
							break
						fi
					done

					echo "$part"

					sum_part=$((sum_part * part))
				done

				sum=$((sum + sum_part))
				parts+=("$part")
			fi
		fi
	done
done

echo "$sum"
