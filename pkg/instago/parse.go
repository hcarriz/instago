package instago

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// Filter will check if post passes the filter and returns a bool on whether to skip or stop.
func (post Post) Filter(filter Filters) (stop, skip bool) {

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
		log.Panic(err)
	}

	created := time.Unix(t, 0)

	// Post was created before a date filter. Skipping post.
	if filter.After != empty && created.Before(filter.After) {
		log.Printf("%s was posted after %v, stopping the search...\n", name, filter.After)
		stop = true
		return
	}

	// Post was created after date. Stopping the loop.
	if filter.Before != empty && created.After(filter.Before) {
		log.Printf("%s was posted before %v, skipping...\n", name, filter.Before)
		skip = true
		return
	}

	// -------------
	// POST FORMAT FILTER
	// -------------

	// Filter out the single posts.
	if filter.CarouselOnly && !multi {
		log.Printf("%s is not a carousel post, skipping...", name)
		skip = true
		return
	}

	// Filter out the carousel posts.
	if filter.SingleOnly && multi {
		log.Printf("%s is not a regular post, skipping...", name)
		skip = true
		return
	}

	// -----------
	// TEXT FILTER
	// -----------

	// Make sure that the post has the desired text in the post caption.
	if filter.Text != "" && !strings.Contains(post.Caption.Text, filter.Text) {
		log.Printf("%s does not have the text %s in the post, skipping...", name, filter.Text)
		skip = true
		return
	}

	// -------------
	// FORMAT FILTER
	// -------------

	if multi {

		// For a carousel post.
		for _, m := range post.CarouselMedia {
			var video bool

			if m.Videos.StandardResolution.URL != "" {
				video = true
			}

			if filter.Videos && !video {
				log.Printf("%s does not have a video, skipping...", name)
				skip = true
				return
			}

			if filter.Images && video {
				log.Printf("%s does not have a image, skipping...", name)
				skip = true
				return
			}

		}

	} else {

		// For a single post.
		var video bool

		if post.Videos.StandardResolution.URL != "" {
			video = true
		}

		if filter.Videos && !video {
			log.Printf("%s does not have a video, skipping...", name)
			skip = true
			return
		}

		if filter.Images && video {
			log.Printf("%s does not have a image, skipping...", name)
			skip = true
			return
		}

	}

	return
}
