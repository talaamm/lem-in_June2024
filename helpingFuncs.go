package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(0)
	}
}

func ReadFile_ln(s string) []string {
	opener, err := os.Open(s)
	CheckErr(err)
	defer opener.Close()
	reader, err := io.ReadAll(opener)
	CheckErr(err)
	casting := string(reader)
	lineoFiles := strings.Split(casting, "\n")
	return lineoFiles
}

func ManipulateLines(lineoFiles []string) []string {
	arr := []string{}
	for index, i := range lineoFiles {
		if len(i) == 0 || i == "" || i == " " {
			continue
		}
		if i[0] == '#' && i[1] != '#' { //#this is a comment  //##start
			continue
		}
		if index != len(lineoFiles)-1 {
			arr = append(arr, i[:len(i)-1])
		} else {
			arr = append(arr, i)
		}
	}
	return arr
}

func CheckRoom(arr1 []string, str string) bool {
	for _, i := range arr1 {
		if i == str {
			return true
		}
	}
	return false
}

func RoomsTunnels_Validity(arr []string) {
	TheRooms := AllRoomsNames(arr)
	for _, i := range arr[1:] {
		if strings.Contains(i, " ") {
			concaconate := strings.Fields(i)
			if len(concaconate) != 3 {
				fmt.Println("Room not valid:", i, "\n"+"name coord_x coord_y")
				os.Exit(0)
			} else if concaconate[0][0] == 'L' {
				fmt.Println("Room Name not valid:", i, "\n"+"A room will never start with the letter L")
				os.Exit(0)
			}
		}

		if strings.Contains(i, "-") {
			concaconate1 := strings.Split(i, "-")
			if concaconate1[1] == "" || concaconate1[0] == "" || len(concaconate1) != 2 {
				fmt.Println("Invalid Data Format:", i)
				os.Exit(0)
			}
			if concaconate1[0] == concaconate1[1] {
				fmt.Println("Cannot link the room with itself!", i)
				os.Exit(0)
			}
			if !bol(concaconate1[0], TheRooms) {
				fmt.Println("Room in Tunnel not found:", concaconate1[0])
				fmt.Println("found rooms:", TheRooms)
				os.Exit(0)
			}
			if !bol(concaconate1[1], TheRooms) {

				fmt.Println("room not found:", concaconate1[1])
				fmt.Println("found rooms:", TheRooms)
				os.Exit(0)
			}
		}
	}
}

func AllRoomsNames(arr []string) []string {
	TheRooms := []string{}
	for index, i := range arr {
		if index == 0 {
			continue
		}
		if strings.Contains(i, " ") {
			concaconate2 := strings.Split(i, " ")
			TheRooms = append(TheRooms, concaconate2[0])
		}
	}
	return TheRooms
}

func bol(s string, str []string) bool {
	for _, i := range str {
		if i == s {
			return true
		}
	}
	return false
}

func AllAntsSteps(dodo [][]string) [][]string {
	popo := [][]string{}
	res := []string{}
	ant := 1
	// for i:= 1 ;
	for _, p1 := range dodo {
		for _, m := range p1[1:] {
			res = append(res, fmt.Sprint("L"+strconv.Itoa(ant), "-", m, " "))
		}
		popo = append(popo, res)
		res = []string{}
		ant++
	}
	return popo
}

func Check2DArrs(xx1, str1 [][]string) bool {
	if len(xx1) == len(str1) { //checks if all paths are shortest path
		for i := range str1 {
			if !notConatians(xx1[i], str1[i]) { //to check if we can tak'em all or just one or what 
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func notConatians(res, last []string) bool {
	for _, r := range last {
		if !bol(r, res) {
			return false
		}
	}
	return true
}

func printStepps(popo, dodo, allpaths [][]string) {
	lastROOMS := []string{}
	for _, steps := range popo {
		lastROOMS = append(lastROOMS, steps[len(steps)-1])
	}

	allrooms := map[string]bool{}
	for _, path := range allpaths {
		for _, room := range path {
			allrooms[room] = false
		}
	}

	res := []string{}
	for !notConatians(res, lastROOMS) {
		for bigInd, steps := range dodo { // step --> [start t  E  a  m  end ]    [h a c K]

			for ind, oneStep := range steps[1:] { // t

				if !allrooms[oneStep] || oneStep == steps[len(steps)-1] {
					if !bol(popo[bigInd][ind], res) {
						res = append(res, popo[bigInd][ind])
						allrooms[oneStep] = true // t: true   h: true
						if oneStep == steps[len(steps)-1] {
							allrooms[oneStep] = false
						}
						break
					}
				} else {
					break
				}
			}
		}
		res = append(res, "\n")
		for ss := range allrooms {
			allrooms[ss] = false
		}
	}
	for _, FINALLY := range res {
		fmt.Print(FINALLY)
	}
}

func findlen(str [][]string) int {
	noor := []int{}
	for _, i := range str {
		noor = append(noor, len(i))
	}
	sort.Ints(noor)
	return noor[0]
}

func findpath(i int, arr [][]string) [][]string {
	bv := [][]string{}
	for _, i1 := range arr {
		if len(i1) == i {
			bv = append(bv, i1)
		}
	}
	return bv
}

func connectStrs(lol []string) string {
	res := ""
	for _, l := range lol {
		res += l
	}
	return res
}

func found(sliceOfSlices [][]string, itemsToDelete []string) bool {
	for _, ar := range sliceOfSlices {
		if connectStrs(ar) == connectStrs(itemsToDelete) {
			return true
		}
	}

	return false
}

func isAtSindex(k []string, p []string) bool {
	kk := (k[1 : len(k)-1])
	pp := (p[1 : len(p)-1])
	for m, n := 0, 0; m <= len(kk)-1 && n <= len(pp)-1; m = m + 1 {
		if kk[m] == pp[n] {
			return true
		}
		n = n + 1
	}
	return false
}

func exclude(pt [][]string, ex []string) [][]string {
	res := [][]string{}
	for _, k := range pt {
		if !isAtSindex(k, ex) {
			if !containsAny(k[1:len(k)-1], ex[1:len(ex)-1]) {
				res = append(res, k)
			}
		}
	}
	return res
}

func containsAny(slice1, slice2 []string) bool {
	for _, s := range slice1 {
		if bol(s, slice2) {
			return true
		}
	}
	return false
}

func deleteItemsFromSliceOfSlices(sliceOfSlices [][]string, itemsToDelete []string) [][]string {
	result := [][]string{}

	for _, ar := range sliceOfSlices {
		if connectStrs(ar) == connectStrs(itemsToDelete) {
			continue
		} else {
			result = append(result, ar)
		}
	}

	return result
}

func removeRepeated(arr [][]string) [][]string {
	rem := [][]string{}
	for i := 0; i < len(arr)-1; i++ {
		for k := 1 + i; k <= len(arr)-1; k++ {
			if isAtSindex(arr[i], arr[k]) {
				if !found(rem, arr[k]) {
					rem = append(rem, arr[k])
				}

				break
			}
		}
	}
	res := arr
	for _, koko := range rem {
		res = deleteItemsFromSliceOfSlices(res, koko)
	}
	return res
}

func printing(possible [][]string, fileData []string) {
	for _, line := range fileData {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Println("Found Paths: ", possible)
	fmt.Println()
	dodo := divideAnts(&farm, possible)
	coco := (AllAntsSteps(dodo))
	printStepps(coco, dodo, possible)
}

//21 funcs
