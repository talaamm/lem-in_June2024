# Lem-in

This project implements a digital ant farm simulation to find the quickest path for ants to traverse a colony of interconnected rooms.

## Overview

The program:
- Reads colony data from a file provided as an argument.
- Simulates the movement of ants from a starting room (`##start`) to an ending room (`##end`) along the shortest path(s).
- Outputs the simulation, showing the steps the ants take in an optimized manner.

The project was developed in Go and adheres to good coding practices, passing all audits.

## Features

1. **Input Parsing**:
   - Handles formatted input for rooms, tunnels, and ants.
   - Validates input for consistency and correctness.

2. **Ant Simulation**:
   - Optimally calculates paths to move ants with minimal traffic jams.
   - Displays step-by-step movements of ants in the required format.

3. **Error Handling**:
   - Detects invalid or poorly formatted input.
   - Provides specific error messages, such as:
     - `ERROR: invalid data format`
     - `ERROR: invalid number of Ants`
     - `ERROR: no start room found`

4. **Bonus**:
   - A visualizer feature (if implemented) displays the ants' movement graphically, enhancing understanding and debugging.

## Input Format

### Ants
The first line specifies the number of ants.

### Rooms
Each room is defined as:
```
name coord_x coord_y
```
- The room name cannot start with `L` or `#`.
- Coordinates are integers.

Special commands:
- `##start` marks the starting room.
- `##end` marks the ending room.

### Tunnels
Each tunnel connects two rooms:
```
room1-room2
```

### Example Input
```
3
##start
0 1 2
##end
1 9 2
2 5 0
0-2
2-1
```

## Output Format
The program outputs:
1. Input data (number of ants, room definitions, and tunnels).
2. Steps taken by ants:
   ```
   L1-room L2-room ...
   ```
   - `L1`, `L2` refer to ant numbers.
   - `room` refers to the destination room.

### Example Output
```
3
##start
0 1 2
##end
1 9 2
2 5 0
0-2
2-1

L1-2
L1-1 L2-2
L2-1
```

## How to Run

1. Compile the program (if needed).
2. Run with:
   ```
   go run . <filename>
   ```
   Replace `<filename>` with the path to the input file.

### Example Command
```
go run . test1.txt
```

## Contributors
- me: **Tala Amm**
- **Moaz Razem**
- **Amro Khweis**
- **Noor Halabi**

## Status
- Successfully passed all audit tests.
- Developed using only standard Go libraries.
