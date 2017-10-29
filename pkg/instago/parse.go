package instago

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// Filter check if post passes the filter and returns a bool on whether to skip or stop.
func (post Post) Filter(filter Filters) (skip bool, stop bool) {

	name := post.Code

	// Easily readable way to determine if a post has multiple media.
	m := len(post.CarouselMedia)

	var multi bool

	if m > 1 {
		multi = true
	}

	// -----------
	// TIME FILTER
	// -----------

	// Parse the time that the Instgram post was created.
	var empty time.Time

	t, err := strconv.ParseInt(post.CreatedTime, 10, 64)
	if err != nil {
		panic(err)
	}

	created := time.Unix(t, 0)

	// Post was created before a date filter. Skipping post.
	if filter.After != empty && created.Before(filter.After) {
		log.Printf("%s was posted after %v, stopping the search...\n", name, filter.After)
		skip = true
		return
	}

	// Post was created after date. Stopping the loop.
	if filter.Before != empty && created.After(filter.Before) {
		log.Printf("%s was posted before %v, skipping...\n", name, filter.Before)
		stop = true
		return
	}

	// -------------
	// POST FORMAT FILTER
	// -------------

	// Filter out the single posts.
	if filter.CarouselOnly && !multi {
		log.Printf("%s is a single post and you only wanted carousel posts only, skipping...", name)
		stop = true
		return
	}

	// Filter out the carousel posts.
	if filter.SingleOnly && multi {
		log.Printf("%s is a carousel post and you only wanted single posts, skipping...", name)
		stop = true
		return
	}

	// -----------
	// TEXT FILTER
	// -----------

	// Make sure that the post has the desired text in the post caption.
	if filter.Has != "" && !strings.Contains(post.Caption.Text, filter.Has) {
		log.Printf("%s does not have the text %s in the post, skipping...", name, filter.Has)
		skip = true
		return
	}

	// -------------
	// FORMAT FILTER
	// -------------

	if multi {

		for _, m := range post.CarouselMedia {
			var video bool

			if m.Videos.StandardResolution.URL != "" {
				video = true
			}

			if filter.Videos && !video {
				log.Printf("%s does not have a video, skipping...", name)
				stop = true
				return
			}

			if filter.Images && video {
				log.Printf("%s does not have a image, skipping...", name)
				stop = true
				return
			}

		}

	} else {

		var video bool

		if post.Videos.StandardResolution.URL != "" {
			video = true
		}

		if filter.Videos && !video {
			log.Printf("%s does not have a video, skipping...", name)
			stop = true
			return
		}

		if filter.Images && video {
			log.Printf("%s does not have a image, skipping...", name)
			stop = true
			return
		}

	}

	return
}
