type guestData struct {
	station string
	time    int
}

type StationKey struct {
	in  string
	out string
}

type StationValue struct {
	countRides int
	sumTime    int
}

type UndergroundSystem struct {
	guestMap   map[int]guestData
	stationMap map[StationKey]StationValue
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		guestMap:   map[int]guestData{},
		stationMap: map[StationKey]StationValue{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.guestMap[id] = guestData{
		station: stationName,
		time:    t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkinData := this.guestMap[id]
	delete(this.guestMap, id)

	// create key
	sk := StationKey{
		in:  checkinData.station,
		out: stationName,
	}

	data := this.stationMap[sk]
	data.countRides++
	data.sumTime += (t - checkinData.time)
	this.stationMap[sk] = data
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := StationKey{
		in:  startStation,
		out: endStation,
	}
	if _, ok := this.stationMap[key]; !ok {
		return 0.0
	}

	stationData := this.stationMap[key]
	return float64(stationData.sumTime) / float64(stationData.countRides)

}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */