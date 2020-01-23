package ApplesAndOranges

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (aHits int32, oHits int32) {
	appleHits := 0
	for _, apple := range apples {
		dropPos := apple + a
		if dropPos >= s && dropPos <= t  {
			appleHits++
		}
	}

	orangeHits := 0
	for _, orange := range oranges {
		dropPos := orange + b
		if dropPos <= t && dropPos >= s {
			orangeHits++
		}
	}

	return aHits, oHits
}
