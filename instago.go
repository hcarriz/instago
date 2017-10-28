package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	dir        string
	pics       bool
	skip       bool
	user       string
	vids       bool
	maxdaysold int
)

type instagram struct {
	Items []struct {
		ID     string `json:"id"`
		Images struct {
			StandardResolution struct {
				URL string `json:"url"`
			} `json:"standard_resolution"`
		} `json:"images"`
		CreatedTime   string `json:"created_time"`
		CarouselMedia []struct {
			Images struct {
				StandardResolution struct {
					URL string `json:"url"`
				} `json:"standard_resolution"`
			} `json:"images"`
			Videos struct {
				StandardResolution struct {
					URL string `json:"url"`
				} `json:"standard_resolution"`
			} `json:"videos,omitempty"`
		} `json:"carousel_media,omitempty"`
		Videos struct {
			StandardResolution struct {
				URL string `json:"url"`
			} `json:"standard_resolution"`
		} `json:"videos,omitempty"`
	} `json:"items"`
	MoreAvailable bool `json:"more_available"`
}

type media struct {
	Image string
	Video string
}

func main() {

	flag.StringVar(&user, "user", "", "user to scrape (required)")
	flag.StringVar(&dir, "dir", "", "where to save the scraped media files (required)")
	flag.IntVar(&maxdaysold, "maxdaysold", 0, "only scrape media within past x days. F.e. use '5' to download content from the past 5 days")
	flag.BoolVar(&pics, "pics", false, "only download images (optional)")
	flag.BoolVar(&vids, "vids", false, "only download videos (optional)")
	flag.BoolVar(&skip, "overwrite", false, "will overwrite previous downloaded images or videos (optional)")

	flag.Parse()

	if user == "" || dir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if pics == true && vids == true {
		pics = false
		vids = false
	}

	if c := checkDir(dir); !c {
		return
	}

	u := "https://www.instagram.com/" + user + "/media/"
	log.Printf("searching for %s....\n", user)

	parse(u)
	log.Printf("finished. exiting....")
	os.Exit(1)

}

func parse(u string) {

	res, err := http.Get(u)
	if err != nil {
		log.Println(err)
		return
	}

	var j instagram
	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
		log.Fatal(err)
		return
	}

	res.Body.Close()

	if len(j.Items) < 1 {
		log.Println("user is private or does not exist.")
		return
	}

	for _, post := range j.Items {

		if maxdaysold != 0 {
			daysAgo, err := daysAgo(post.CreatedTime)
			if err != nil {
				log.Println(err)
			}

			if daysAgo > maxdaysold {
				log.Printf("skipping post form %d days ago (%s)", daysAgo, post.Images.StandardResolution.URL)
				break
			}
		}

		carousel := post.CarouselMedia

		if len(carousel) > 0 {
			for _, seat := range carousel {

				m := media{
					Video: seat.Videos.StandardResolution.URL,
					Image: seat.Images.StandardResolution.URL,
				}

				if err := m.save(); err != nil {
					log.Println(err)
					break
				}

			}
		}

		m := media{
			Video: post.Videos.StandardResolution.URL,
			Image: post.Images.StandardResolution.URL,
		}

		if err := m.save(); err != nil {
			log.Println(err)
			break
		}

	}

	if j.MoreAvailable {
		id := j.Items[len(j.Items)-1].ID
		u := "https://www.instagram.com/" + user + "/media/?max_id=" + id
		parse(u)
	}

}

func (m media) save() error {

	url := m.Video

	if url == "" && vids {
		return nil
	}

	if url == "" || pics {
		url = m.Image
		url = strings.Replace(url, "s640x640", "s1080x1080", -1)
	}

	d := filepath.Join(dir, user)

	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.Mkdir(d, 0755); err != nil {
			log.Fatal(err)
			return err
		}
	}

	_, name := filepath.Split(url)
	newPath := filepath.Join(d, name)

	if _, err := os.Stat(newPath); !os.IsNotExist(err) && !skip {
		log.Printf("%s has already been saved. skipping....\n", name)
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	file, err := os.Create(newPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("successfully downloaded %s\n", name)
	return nil

}

func checkDir(d string) (forward bool) {

	if _, err := os.Stat(d); os.IsNotExist(err) {
		log.Println("directory does not exist")
		return
	}

	new := filepath.Join(d, "tempDir")

	if err := os.Mkdir(new, 0755); err != nil {
		log.Println("unable to create directory")
		return
	}

	if err := os.Remove(new); err != nil {
		log.Println("unable to remove directory")
		return
	}

	return true

}

func daysAgo(unixTimestamp string) (int, error) {
	i, err := strconv.ParseInt(unixTimestamp, 10, 64)
	if err != nil {
		return 0, err
	}

	then := time.Unix(i, 0)
	now := time.Now()
	diff := now.Sub(then)
	days := int(diff.Hours() / 24)
	return days, nil
}
