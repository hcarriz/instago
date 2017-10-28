package instago

func (i Instagram) parse() (posts []post) {

	for _, single := range i.Items {

		temp := post{
			ID:       single.ID,
			Created:  single.CreatedTime,
			Caption:  single.Caption.Text,
			Likes:    single.Likes.Count,
			Comments: single.Comments.Count,
			Location: single.Location.Name,
		}

		carousel := single.CarouselMedia

		if len(carousel) > 0 {

			for _, seat := range carousel {

				t := solve(seat.Images.StandardResolution.URL, seat.Videos.StandardResolution.URL)
				temp.Media = append(temp.Media, t)

			}

		} else {

			t := solve(single.Images.StandardResolution.URL, single.Videos.StandardResolution.URL)
			temp.Media = append(temp.Media, t)

		}

		posts = append(posts, temp)

	}

	return

}

func solve(picture, video string) postmedia {

	var url string
	var vid bool

	switch {
	case video != "":
		url = video
		vid = true
	default:
		url = picture
	}

	t := postmedia{
		URL:   url,
		Video: vid,
	}

	return t

}
