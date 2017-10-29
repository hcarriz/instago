package instago

import "log"

// Save will
func (post Post) Save(filter Filters) {

	log.Printf("http://instagram.com/p/%s passed filters. saving....\n", post.Code)

	if !filter.CarouselOnly {
		files := post.Media
		urls := files.urls(filter)

		for get, url := range urls {
			if get {
				if err := download(url); err != nil {
					log.Fatal(err)
				}
			}
		}

	}

	if !filter.SingleOnly {
		files := post.CarouselMedia

		for _, f := range files {
			file := f.urls(filter)

			for get, url := range file {
				if get {
					if err := download(url); err != nil {
						log.Fatal(err)
					}
				}
			}

		}
	}

}

func (media Media) urls(filter Filters) map[bool]string {

	log.Println(media.Images.StandardResolution.URL)

	return nil

}

func download(url string) error {
	return nil
}

// func (m media) save() error {

// 	url := m.Video

// 	if url == "" && vids {
// 		return nil
// 	}

// 	if url == "" || pics {
// 		url = m.Image
// 		url = strings.Replace(url, "s640x640", "s1080x1080", -1)
// 	}

// 	d := filepath.Join(dir, user)

// 	if _, err := os.Stat(d); os.IsNotExist(err) {
// 		if err := os.Mkdir(d, 0755); err != nil {
// 			log.Fatal(err)
// 			return err
// 		}
// 	}

// 	_, name := filepath.Split(url)
// 	newPath := filepath.Join(d, name)

// 	if _, err := os.Stat(newPath); !os.IsNotExist(err) && !skip {
// 		log.Printf("%s has already been saved. skipping....\n", name)
// 		return nil
// 	}

// 	res, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}

// 	defer res.Body.Close()

// 	file, err := os.Create(newPath)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	defer file.Close()

// 	if _, err := io.Copy(file, res.Body); err != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	log.Printf("successfully downloaded %s\n", name)
// 	return nil

// }
