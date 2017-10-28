package instago

// Instagram is JSON that Instagram uses.
type Instagram struct {
	MoreAvailable bool   `json:"more_available"`
	Status        string `json:"status"`
	Items         []Post `json:"items"`
}

// Post has the information about a single Instagram post.
type Post struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	User struct {
		ID             string `json:"id"`
		FullName       string `json:"full_name"`
		ProfilePicture string `json:"profile_picture"`
		Username       string `json:"username"`
	} `json:"user"`
	Images struct {
		Thumbnail struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"thumbnail"`
		LowResolution struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"low_resolution"`
		StandardResolution struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"standard_resolution"`
	} `json:"images"`
	CreatedTime string `json:"created_time"`
	Caption     struct {
		ID          string `json:"id"`
		Text        string `json:"text"`
		CreatedTime string `json:"created_time"`
		From        struct {
			ID             string `json:"id"`
			FullName       string `json:"full_name"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
		} `json:"from"`
	} `json:"caption"`
	Likes struct {
		Data []struct {
			ID             string `json:"id"`
			FullName       string `json:"full_name"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
		} `json:"data"`
		Count int `json:"count"`
	} `json:"likes"`
	Comments struct {
		Data []struct {
			ID          string `json:"id"`
			Text        string `json:"text"`
			CreatedTime string `json:"created_time"`
			From        struct {
				ID             string `json:"id"`
				FullName       string `json:"full_name"`
				ProfilePicture string `json:"profile_picture"`
				Username       string `json:"username"`
			} `json:"from"`
		} `json:"data"`
		Count int `json:"count"`
	} `json:"comments"`
	CanViewComments   bool   `json:"can_view_comments"`
	CanDeleteComments bool   `json:"can_delete_comments"`
	Type              string `json:"type"`
	Link              string `json:"link"`
	Location          struct {
		Name string `json:"name"`
	} `json:"location"`
	CarouselMedia []struct {
		Videos struct {
			StandardResolution struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
				ID     string `json:"id"`
			} `json:"standard_resolution"`
			LowBandwidth struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
				ID     string `json:"id"`
			} `json:"low_bandwidth"`
			LowResolution struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
				ID     string `json:"id"`
			} `json:"low_resolution"`
		} `json:"videos,omitempty"`
		Images struct {
			Thumbnail struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"thumbnail"`
			LowResolution struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"low_resolution"`
			StandardResolution struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"standard_resolution"`
		} `json:"images"`
		UsersInPhoto []interface{} `json:"users_in_photo"`
		Type         string        `json:"type"`
	} `json:"carousel_media,omitempty"`
	Videos struct {
		StandardResolution struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
			ID     string `json:"id"`
		} `json:"standard_resolution"`
		LowBandwidth struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
			ID     string `json:"id"`
		} `json:"low_bandwidth"`
		LowResolution struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
			ID     string `json:"id"`
		} `json:"low_resolution"`
	} `json:"videos,omitempty"`
	VideoViews int `json:"video_views,omitempty"`
}

type post struct {
	ID       string
	Created  string
	Caption  string
	Likes    int
	Comments int
	Location string
	Media    []postmedia
}

type postmedia struct {
	URL   string
	Video bool
}
