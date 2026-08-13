type TimeMap struct {
	items map[string]string
	timestamps map[string][]int
}
func Constructor() TimeMap {
	return TimeMap{ 
		items: make(map[string]string),
		timestamps: make(map[string][]int),
	}
}
func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, e := this.timestamps[key]; !e {
		this.timestamps[key] = []int{}
	}
	this.timestamps[key] = append(this.timestamps[key], timestamp)

	timekey := key + strconv.Itoa(timestamp)
	this.items[timekey] = value
}
func (this *TimeMap) Get(key string, timestamp int) string {
	var timestamps []int
	var e bool
	
	if timestamps, e = this.timestamps[key]; !e {
		fmt.Println("Error: key does not exist")
		return ""
	}

	time := -1
	for _, t := range timestamps {
		if t == timestamp {
			time = t
			break
		}
		if t != timestamp && t < timestamp {
			time = t
		}
	}
	if time == -1 {
		fmt.Println("Error: No timestamps in key")
		return ""
	}

	timekey := key + strconv.Itoa(time)
	return this.items[timekey]
}
