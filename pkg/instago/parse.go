package instago

import (
	"fmt"
	"strconv"
	"time"
)

// Filters is used in a channel to relay information if a post is or isn't wanted by the user.
type Filters struct {
	Before       time.Time
	After        time.Time
	CarouselOnly bool
	SingleOnly   bool
	Videos       bool
	Pictures     bool
	Amount       int
}

// Filter check if post passes the filter and returns a bool on whether to skip or stop.
func (post Post) Filter(filter Filters) (skip bool, stop bool) {

	// TIME FILTER

	var empty time.Time

	t, err := strconv.ParseInt(post.CreatedTime, 10, 64)
	if err != nil {
		panic(err)
	}

	created := time.Unix(t, 0)

	// Post was created before a date filter. Skipping post.
	if filter.After != empty && created.Before(filter.After) {
		skip = true
		return
	}

	// Post was created after date. Stopping the loop.
	if filter.Before != empty && created.After(filter.Before) {
		stop = true
		return
	}

	// FORMAT FILTER

	return
}

// Save will
func (post Post) Save(filter Filters) {

	fmt.Printf("post %s passed filters. saving....\n", post.ID)

}
