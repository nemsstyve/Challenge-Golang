package piscine

func SortIntegerTable(table []int) {
	sort := true
	for sort {
		sort = false
		for i := 1; i < len(table); i++ {
			if table[i-1] > table[i] {
				table[i], table[i-1] = table[i-1], table[i]
				sort = true
			}
		}
	}
}
