package Day09

type distMap map[string]map[string]int

func (dm *distMap) addDistance(city1, city2 string, dist int) {
	if _, ok := (*dm)[city1]; !ok {
		(*dm)[city1] = map[string]int{}
	}

	if _, ok := (*dm)[city2]; !ok {
		(*dm)[city2] = map[string]int{}
	}

	(*dm)[city1][city2] = dist
	(*dm)[city2][city1] = dist
}

func (dm *distMap) totalDistance() {
	for _, m := range *dm {
		for _, dist := range m {

		}
	}
}