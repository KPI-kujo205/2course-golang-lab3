#!/bin/bash

# Constants for the spiral
a=0.01
b=0.01

# Angle step size
angle_step=0.1

# Number of iterations
iterations=1000

# Define the array of commands to draw the figure
commands=(
"white"
"bgrect 0.25 0.25 0.75 0.75"
"green"
)

# Iterate over the array and execute each command
for command in "${commands[@]}"; do
    curl -X POST http://localhost:17000 -d "$command"
done

# Move the figure in a spiral
for ((i=0; i<$iterations; i++)); do
    # Calculate the angle in radians
    angle=$(echo "$i * $angle_step" | bc -l)

    # Calculate the polar coordinates
    r=$(echo "$a + $b * $angle" | bc -l)
    x=$(echo "$r * c($angle)" | bc -l)
    y=$(echo "$r * s($angle)" | bc -l)

    # Convert the polar coordinates to Cartesian coordinates
    x=$(echo "0.5 + $x / 2" | bc -l)
    y=$(echo "0.5 + $y / 2" | bc -l)

    # Move the figure
    curl -X POST http://localhost:17000 -d "figure $x $y"

    # Update the display
    curl -X POST http://localhost:17000 -d "update"

    # Wait for a short time
    sleep 0.1
done
