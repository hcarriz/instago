package instago

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Save will get the filtered media and download the post media.
func (post Post) Save(filter Filters) {

	// Slice to hold valid url(s)
	var urls []string

	// Notify the user that a file is being saved.
	log.Printf("%s passed filters. saving....\n", post.Code)

	// Easily readable way to determine if a post has multiple media.
	m := len(post.CarouselMedia)

	var multi bool

	if m > 1 {
		multi = true
	}

	// Find the valid media
	if multi {

		// If the post has multiple media.
		for _, p := range post.CarouselMedia {
			if url, valid := p.URL(filter); valid {
				urls = append(urls, url)
			}
		}

	} else {

		// If the post has only one media item.
		if url, valid := post.URL(filter); valid {
			urls = append(urls, url)
		}

	}

	// Save the valid urls.
	for x := range urls {

		err := make(chan error)
		go post.Download(urls[x], filter, err)
		if e := <-err; e != nil {
			log.Fatal(e)
		}

	}

}

// URL will return valid url of a media file.
func (m Media) URL(filter Filters) (url string, valid bool) {

	video := m.Videos.StandardResolution.URL
	photo := m.Images.StandardResolution.URL

	switch {
	case filter.Videos && video == "":
		return
	case filter.Videos && video != "":
		url = video
	case filter.Images && video != "":
		return
	default:
		url = photo
		url = strings.Replace(url, "s640x640", "s1080x1080", -1)
	}

	valid = true
	return
}

// Download ---
func (post Post) Download(url string, filter Filters, e chan error) {

	main := filepath.Join(filter.Directory, post.User.Username)

	if _, err := os.Stat(main); os.IsNotExist(err) {
		if err := os.Mkdir(main, 0755); err != nil {
			e <- err
			return
		}
	}

	_, name := filepath.Split(url)
	path := filepath.Join(main, name)

	if _, err := os.Stat(path); !os.IsNotExist(err) && !filter.Overwrite {
		log.Printf("%s has already been downloaded, skipping...\n", name)
		e <- nil
		return
	}

	res, err := http.Get(url)
	if err != nil {
		e <- err
		return
	}

	defer res.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		e <- err
		return
	}

	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		e <- err
		return
	}

	log.Printf("%s has been saved\n", post.Code)

	e <- nil
}
