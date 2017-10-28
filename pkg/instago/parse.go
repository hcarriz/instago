package instago

import (
	"strconv"
	"strings"
	"time"
)

// Filter check if post passes the filter and returns a bool on whether to skip or stop.
func (post Post) Filter(filter Filters) (skip bool, stop bool) {

	// TIME FILTER

	// Parse the time that the Instgram post was created.
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

	// TEXT FILTER
	if filter.Has != "" && !strings.Contains(post.Caption.Text, filter.Has) {
		skip = true
		return
	}

	return
}
