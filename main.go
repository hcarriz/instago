package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/araddon/dateparse"
	"github.com/pagumin/instago/pkg/instago"
	"github.com/pagumin/instago/pkg/utilities"
)

// Flags
var (
	after      string
	afterDate  time.Time
	before     string
	beforeDate time.Time
	carousels  bool
	dir        string
	has        string
	max        int
	pics       bool
	singles    bool
	user       string
	vids       bool
	zone       string
)

func main() {

	// Initialize the flags
	flag.StringVar(&user, "user", "", "user to scrape (required)")
	flag.StringVar(&dir, "dir", "~/", "where to save the scraped media files (required)")
	flag.BoolVar(&pics, "pics", false, "only download images (optional)")
	flag.BoolVar(&vids, "vids", false, "only download videos (optional)")
	flag.StringVar(&before, "before", "", "get posts before date (optional)")
	flag.StringVar(&after, "after", "", "get posts after a date (optional)")
	flag.StringVar(&zone, "timezone", "UTC", "Timezone aka `America/Los_Angeles` formatted time-zone (optional)")
	flag.BoolVar(&carousels, "carousel", false, "only download media from carousel posts (optional)")
	flag.BoolVar(&singles, "single", false, "only download media from single posts (optional)")
	flag.IntVar(&max, "max", 0, "the maximum amount of valid/filtered posts to download (0 means all valid posts)")
	flag.StringVar(&has, "has", "", "download a post if it has certain text (optional)")

	flag.Parse()

	// Always require the user to search and the directory to download to.
	if user == "" || dir == "" {

		flag.PrintDefaults()
		os.Exit(2)

	}

	// Check the directory
	if c := utilities.ValidDir(dir); !c {
		log.Fatal("need a valid directory with write access to download to")
	}

	// Fix the timezone
	if zone != "" {

		loc, err := time.LoadLocation(zone)
		if err != nil {
			log.Fatal(err)
		}

		time.Local = loc

	}

	// Dates
	if b, err := dateparse.ParseLocal(before); err == nil {
		beforeDate = b
	}

	if a, err := dateparse.ParseLocal(after); err == nil {
		afterDate = a
	}

	// Allow for both commands to be used at the same time.
	if pics == true && vids == true {

		pics = false
		vids = false

	}

	// If both commands are used together, they won't cancel each other out.
	if carousels == true && singles == true {

		carousels = false
		singles = false

	}

	// Provide feedback that the search has started.
	fmt.Printf("searching for %s\n", user)

	// Get the posts.
	data := make(chan instago.Instagram)
	err := make(chan error)
	stop := make(chan bool)
	go instago.Retrieve(user, "", data, err)

	// Filter the posts

	filter := instago.Filters{
		Before:       beforeDate,
		After:        afterDate,
		CarouselOnly: carousels,
		SingleOnly:   singles,
		Videos:       vids,
		Pictures:     pics,
		Amount:       max,
		Has:          has,
	}
	x := 1

	for i := 0; i < x; i++ {

		select {
		case e := <-err:
			fmt.Println(e)
		case d := <-data:

			for _, post := range d.Items {

				skip, stop := post.Filter(filter)

				switch {
				case stop:
					break
				case skip:
					return
				default:
					post.Save(filter)
				}

			}

			if d.MoreAvailable {
				x++
			}

		case <-stop:
			break

		}

	}

	return

}
