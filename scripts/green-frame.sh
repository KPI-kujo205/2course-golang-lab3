#!/bin/bash

commands=(
"white"
"bgrect 0.25 0.25 0.75 0.75"
"figure 0.5 0.5"
"green"
"figure 0.6 0.6"
"update"
)

# Iterate over the array and execute each command
for command in "${commands[@]}"; do
    curl -X POST http://localhost:17000 -d "$command"
done
