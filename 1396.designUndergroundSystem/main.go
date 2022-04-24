package main

// ["UndergroundSystem","checkIn","checkIn","checkIn","checkOut","checkOut","checkOut","getAverageTime","getAverageTime","checkIn","getAverageTime","checkOut","getAverageTime"]
// [[],[45,"Leyton",3],[32,"Paradise",8],[27,"Leyton",10],[45,"Waterloo",15],[27,"Waterloo",20],[32,"Cambridge",22],["Paradise","Cambridge"],["Leyton","Waterloo"],[10,"Leyton",24],["Leyton","Waterloo"],[10,"Waterloo",38],["Leyton","Waterloo"]]

func main() {
	system := Constructor()
	system.CheckIn(45, "Leyton", 3)
	system.CheckIn(32, "Paradise", 8)
	system.CheckIn(27, "Leyton", 10)
	system.CheckOut(45, "Waterloo", 15)
	system.CheckOut(27, "Waterloo", 20)
	system.CheckOut(32, "Cambridge", 22)
	system.GetAverageTime("Paradise", "Cambridge")
	system.GetAverageTime("Leyton", "Waterloo")
	system.CheckIn(10, "Leyton", 24)
	system.GetAverageTime("Leyton", "Waterloo")
	system.CheckOut(10, "Leyton", 38)
	system.GetAverageTime("Leyton", "Waterloo")
}

type checkIn struct {
	checkInStation string
	checkInTime    int
}

type UndergroundSystem struct {
	users           map[int]checkIn  // key is {id}
	travellingTimes map[string][]int // key is {startStation:endStation}
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		users:           make(map[int]checkIn),
		travellingTimes: make(map[string][]int),
	}
}

func (s *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	s.users[id] = checkIn{
		checkInStation: stationName,
		checkInTime:    t,
	}
}

func (s *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	info := s.users[id]
	travellingKey := info.checkInStation + ":" + stationName
	s.travellingTimes[travellingKey] = append(s.travellingTimes[travellingKey], t-info.checkInTime)
	delete(s.users, id)
}

func (s *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	travellingKey := startStation + ":" + endStation

	var avg float64
	for _, t := range s.travellingTimes[travellingKey] {
		avg += float64(t)
	}
	avg /= float64(len(s.travellingTimes[travellingKey]))

	// fmt.Printf("avg time from %s to %s = %.2f\n", startStation, endStation, avg)
	return avg
}
