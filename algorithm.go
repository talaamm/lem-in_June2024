package main

func dfs(antFarm *AntFarm, current, end Room, visited map[string]bool, path []string, allPaths *[][]string) {
	// Add the current room to the path
	path = append(path, current.Name)

	// If we have reached the end room, add the current path to allPaths
	if current == end {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*allPaths = append(*allPaths, newPath)
		return
	}

	// Mark the current room as visited
	visited[current.Name] = true

	// Visit each neighbor room
	for _, tunnel := range antFarm.Tunnels {
		var neighborName string
		if tunnel.Room1 == current.Name {
			neighborName = tunnel.Room2
		} else if tunnel.Room2 == current.Name {
			neighborName = tunnel.Room1
		} else {
			continue
		}

		// If the neighbor has not been visited, continue the search
		if !visited[neighborName] {
			dfs(antFarm, antFarm.RoomsMap[neighborName], end, visited, path, allPaths)
		}
	}

	// Backtrack: unmark the current room as visited
	visited[current.Name] = false
}

func findAllPaths(antFarm *AntFarm, start, end Room) [][]string {
	allPaths := [][]string{}
	visited := make(map[string]bool)
	dfs(antFarm, start, end, visited, []string{}, &allPaths)
	return allPaths
}

func divideAnts(antFarm *AntFarm, paths [][]string) [][]string {
	// 4 --> 3      b  a  c  5
	// 6 --> 2
	// 7 --> 1 -2

	antPaths := make([][]string, antFarm.Ants) //array string with len the ants

	if len(paths) == 1 {
		for antnum := 0; antnum < antFarm.Ants; antnum++ {
			antPaths[antnum] = paths[0]
		}
		return antPaths
	}

	pathLengths := make([]int, len(paths))     // saves in it the length of the path without the start room
	antDistribution := make([]int, len(paths)) //taqseem el ants (num of ants to go to certain path)

	for i, path := range paths {
		pathLengths[i] = len(path) - 1
	}

	for antnum := 0; antnum < antFarm.Ants; antnum++ {
		shortestIndex := 0
		for i := 1; i < len(paths); i++ {
			if antDistribution[i]+pathLengths[i] < antDistribution[shortestIndex]+pathLengths[shortestIndex] {
				shortestIndex = i
			}
		}
		antDistribution[shortestIndex]++
		antPaths[antnum] = paths[shortestIndex]
	}

	return antPaths
}

func countStepsToEnd(antFarm *AntFarm, paths1 [][]string) int {
	if len(paths1) == 0 {
		return -1 // No valid path found
	}

	antPaths := divideAnts(antFarm, paths1)
	antPositions := make([]int, antFarm.Ants)
	roomOccupancy := make(map[string]int)
	steps := 0

	for _, path := range antPaths {
		for _, room := range path {
			roomOccupancy[room] = 0
		}
	}
	roomOccupancy[antFarm.Start.Name] = antFarm.Ants

	for {
		allFinished := true

		for ant := 0; ant < antFarm.Ants; ant++ {
			if antPositions[ant] < len(antPaths[ant])-1 {
				allFinished = false
				nextRoom := antPaths[ant][antPositions[ant]+1]
				if roomOccupancy[nextRoom] == 0 || nextRoom == antFarm.End.Name {
					currentRoom := antPaths[ant][antPositions[ant]]
					if currentRoom != antFarm.Start.Name {
						roomOccupancy[currentRoom]--
					}
					antPositions[ant]++
					roomOccupancy[nextRoom]++
				}
			}
		}

		if allFinished {
			break
		}

		steps++
	}

	return steps
}
