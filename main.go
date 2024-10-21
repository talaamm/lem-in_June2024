package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	roomxy []Room
	tun    []Tunnel
	farm   AntFarm
	start  Room
	end    Room
)

func main() {
	////////////////////////////////////// Handling Errors ////////////////////////////////////////////////////////////////
	if len(os.Args) != 2 {
		return
	}

	filename := os.Args[1]
	FilesLines := ReadFile_ln(filename)
	myLines := ManipulateLines(FilesLines)

	if !CheckRoom(myLines, "##start") {
		fmt.Println("missing start room")
		return
	} else if !CheckRoom(myLines, "##end") {
		fmt.Println("missing end room")
		return
	}

	RoomsTunnels_Validity(myLines)

	////////////////////////////////////// Parsing input ////////////////////////////////////////////////////////////////
	antsNum, err := strconv.Atoi(myLines[0])
	if err != nil || antsNum == 0 {
		fmt.Println("we need a valid number of ants")
		return
	}
	myLines = myLines[1:]
	for index, i := range myLines {
		if strings.Contains(i, " ") {
			GG := strings.Split(i, " ")
			con, err := strconv.Atoi(GG[1])
			CheckErr(err)
			con1, err := strconv.Atoi(GG[2])
			CheckErr(err)
			create := Room{
				Name: GG[0],
				X:    con,
				Y:    con1,
			}
			roomxy = append(roomxy, create)
		} else if strings.Contains(i, "-") {
			GG1 := strings.Split(i, "-")
			rom := Tunnel{
				Room1: GG1[0],
				Room2: GG1[1],
			}
			tun = append(tun, rom)
		} else if i == "##start" {
			GG := strings.Split(myLines[index+1], " ")
			con, err := strconv.Atoi(GG[1])
			CheckErr(err)
			con1, err := strconv.Atoi(GG[2])
			CheckErr(err)
			start = Room{
				Name: GG[0],
				X:    con,
				Y:    con1,
			}
			// did not append to ants bc in the coming iteration it will be appended
		} else if i == "##end" {
			GG := strings.Split(myLines[index+1], " ")
			con, err := strconv.Atoi(GG[1])
			CheckErr(err)
			con1, err := strconv.Atoi(GG[2])
			CheckErr(err)
			end = Room{
				Name: GG[0],
				X:    con,
				Y:    con1,
			}
		}
	}
	roomsMap := make(map[string]Room)
	for _, room := range roomxy {
		roomsMap[room.Name] = room
	}

	farm = AntFarm{
		Ants:     antsNum,
		Rooms:    roomxy,
		Tunnels:  tun,
		Start:    start,
		End:      end,
		RoomsMap: roomsMap,
	}

	////////////////////////////////////// calculating & Printing shortest path ////////////////////////////////////////////////////////////////

	allPaths := findAllPaths(&farm, farm.Start, farm.End) // all possible paths

	if len(allPaths) == 0 {
		fmt.Println("No valid path found")
		os.Exit(0)
	} else if len(allPaths) == 1 {
		printing(allPaths, FilesLines)
		return
	} else {

		xx := findlen(allPaths)       // len of shortest path
		xx1 := findpath(xx, allPaths) // all paths with the len of the shortest path

		str1 := allPaths   //a copy of allpaths to edit
		if len(xx1) == 1 { //only one short path
			finalss := [][]string{}
			finalss = append(finalss, xx1[0]) //add the shortest path to the array

			for countStepsToEnd(&farm, finalss) >= antsNum {
				for _, xx3 := range xx1 {
					str1 = exclude(str1, xx3)
					//exclude the already added path from allpaths
					//& all the paths with similar rooms at same index
				}
				// fmt.Println("All paths:", allPaths)
				// fmt.Println("All paths after exclude:", str1)
				// fmt.Println("excluded as:", xx1)
				// os.Exit(0)
				if len(str1) == 0 { //see myexample
					min := countStepsToEnd(&farm, xx1)
					lenthepath := len(xx1[0])
					for _, r := range allPaths {
						fmt.Println(countStepsToEnd(&farm, [][]string{r}), r)
						if countStepsToEnd(&farm, [][]string{r}) < min {
							min = countStepsToEnd(&farm, [][]string{r})
							lenthepath = len(r)
						}
					}
					theFinal := findpath(lenthepath, allPaths)[0]
					printing([][]string{theFinal}, FilesLines)
					os.Exit(0)
				}
				xx = findlen(str1) //len second shortest path
				xx1 = findpath(xx, str1)
				for _, yu := range finalss { //which conatins the shorteset
					for _, lp := range xx1 { //which contains second shortest
						if !(isAtSindex(yu, lp)) { //making sure no rooms at same index
							if !(found(finalss, lp)) { //making sure no rooms exist in the 2 paths
								if !containsAny(yu[1:len(yu)-1], lp[1:len(lp)-1]) { //[start 3 4 2 1 end]
									finalss = append(finalss, lp)
									break
								}
							}
						}
					}
				}

			}
			printing(finalss, FilesLines) //remember its a for loop
			return
		} else { //multiple short paths
			if Check2DArrs(xx1, str1) {
				// for _, xx3 := range xx1 {
				// 	str1 = deleteItemsFromSliceOfSlices(str1, xx3)
				// }
				printing(str1, FilesLines)
				return
			}
			if countStepsToEnd(&farm, xx1) <= antsNum { // if not all short paths just check for steps count
				xx1 = removeRepeated(xx1) //to get 2/3 short paths that dont have same inex or repeated
				printing(xx1, FilesLines)
				return
			}

			for _, xx3 := range xx1 {
				str1 = deleteItemsFromSliceOfSlices(str1, xx3) //only delete the short paths ONLY
			}

			// fmt.Println(xx1)
			// return

			var1 := findlen(str1)
			var2 := findpath(var1, str1)
			neww := removeRepeated(var2)
			// fmt.Println(neww)
			// return
			possible := [][]string{}
			possible = append(possible, neww...) //check for possible ways from combination of short paths
			for _, yu := range neww {
				for _, lp := range xx1 {
					if !(isAtSindex(yu, lp)) {
						if !found(possible, lp) {
							possible = append(possible, lp)
							break
						}
					}
				}
			}
			// fmt.Println(countStepsToEnd(&farm, neww))
			// fmt.Println(possible,countStepsToEnd(&farm, possible))
			if countStepsToEnd(&farm, neww) > countStepsToEnd(&farm, possible) {
				// fmt.Println("ana hon")
				printing(possible, FilesLines)
			} else {
				printing(neww, FilesLines)
			}
		}
		return
	}
}
