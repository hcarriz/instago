package instago

import "time"

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
	Media
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
		Media
		UsersInPhoto []interface{} `json:"users_in_photo"`
		Type         string        `json:"type"`
	} `json:"carousel_media,omitempty"`
	VideoViews int `json:"video_views,omitempty"`
}

// Media combines Images and Videos
type Media struct {
	Images Images `json:"images"`
	Videos Videos `json:"videos,omitempty"`
}

// Images has the media information about a picture
type Images struct {
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
}

// Videos has the information about a video.
type Videos struct {
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
}

// Filters is used in a channel to relay information if a post is or isn't wanted by the user.
type Filters struct {
	After        time.Time
	Amount       int
	Before       time.Time
	CarouselOnly bool
	Text         string
	Images       bool
	Overwrite    bool
	SingleOnly   bool
	Videos       bool
	Directory    string
}
