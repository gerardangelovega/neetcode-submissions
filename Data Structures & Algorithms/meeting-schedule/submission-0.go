import (
	"slices"
	"cmp"
)
/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})

	for i := 0; i < len(intervals) - 1; i = i + 1 {
		if intervals[i + 1].start < intervals[i].end {
			return false
		}
	}

	return true
}
