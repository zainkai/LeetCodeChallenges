import "sort"

const (
	restObjId            = 0
	restObjRating        = 1
	restObjVeganFriendly = 2
	restObjPrice         = 3
	restObjDistance      = 4
)

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filteredResturants := make([][]int, 0, len(restaurants))
	for _, r := range restaurants {
		if veganFriendly == 1 && r[restObjVeganFriendly] != 1 {
			continue
		} else if r[restObjPrice] > maxPrice {
			continue
		} else if r[restObjDistance] > maxDistance {
			continue
		} else {
			filteredResturants = append(filteredResturants, r)
		}
	}

	sort.Slice(filteredResturants, func(i, j int) bool {
		if filteredResturants[i][restObjRating] != filteredResturants[j][restObjRating] {
			return filteredResturants[i][restObjRating] > filteredResturants[j][restObjRating]
		}
		return filteredResturants[i][restObjId] > filteredResturants[j][restObjId]
	})

	ids := make([]int, len(filteredResturants))
	for i, r := range filteredResturants {
		ids[i] = r[0]
	}

	return ids
}
