package instago

import (
	"encoding/json"
	"net/http"
)

// Retrieve will return the data from the user.
func Retrieve(user string, id string, data chan Instagram, e chan error) {

	var d Instagram

	url := "https://www.instagram.com/" + user + "/media/"

	if id != "" {
		url = url + "?max_id=" + id
	}

	res, err := http.Get(url)
	if err != nil {
		e <- err
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		e <- err
		return
	}

	res.Body.Close()

	if d.MoreAvailable {

		last := d.Items[len(d.Items)-1].ID
		go Retrieve(user, last, data, e)

	}

	data <- d

	return

}
