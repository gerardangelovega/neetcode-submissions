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

	l, r := 0, len(timestamps)-1
	t := -1
	for l <= r {
		m := (l + r) / 2
		if timestamp == timestamps[m] {
			t = timestamps[m]
			break
		} else if timestamp < timestamps[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if t == -1 {
		if timestamps[max(0, min(l,r))] < timestamp {
			t = timestamps[max(0, min(l,r))]
		} else {
			return ""
		}
	}

	timekey := key + strconv.Itoa(t)
	return this.items[timekey]
}

// 1,2,5,7,11,13