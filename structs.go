package main

type Room struct {
	Name string
	X, Y int
}

type Tunnel struct {
	Room1, Room2 string
}

type AntFarm struct {
	Ants     int
	Rooms    []Room
	Tunnels  []Tunnel
	Start    Room
	End      Room
	RoomsMap map[string]Room //necessary to find allpaths
	// ["fofo"]: Room{Name:"fofo" , X: 5 , Y: 2}
}
