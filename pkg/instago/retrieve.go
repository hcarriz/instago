package instago

import (
	"encoding/json"
	"net/http"
)

// Retrieve will return the data from the user
func Retrieve(user string, data chan Instagram, err chan error) {

	url := "https://www.instagram.com/" + user + "/media/"

	d, e := fetch(url)
	if e != nil {
		err <- e
		return
	}

	data <- d

	return
}

func fetch(url string) (data Instagram, err error) {

	res, err := http.Get(url)
	if err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	res.Body.Close()

	return
}

// // Retrieve gets all of the information about a user.
// func Retrieve(user string, id string, data chan instagram) {

// 	url := "https://www.instagram.com/" + user + "/media/"

// 	if id != "" {

// 		url = url + "?max_id=" + id
// 	}

// 	res, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	res.Body.Close()

// }

// func retrieve(u string) {

// 	res, err := http.Get(u)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	var j instagram
// 	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	res.Body.Close()

// 	if len(j.Items) < 1 {
// 		log.Println("user is private or does not exist.")
// 		return
// 	}

// 	if j.MoreAvailable {
// 		id := j.Items[len(j.Items)-1].ID
// 		u := "https://www.instagram.com/" + user + "/media/?max_id=" + id
// 		retrieve(u)
// 	}

// 	j.parse()

// }
