# instago
[![Go Report Card](https://goreportcard.com/badge/github.com/pagumin/instago)](https://goreportcard.com/report/github.com/pagumin/instago)

Scrape an Instagram user with Go. Get the latest release [here](https://github.com/pagumin/instago/releases).

## Usage

Default usage:

``` bash
instago -user <user> -dir <save_directory> [additional commands...]
```

## Commands

These are the commands that can be used in instago. All of the commands can be used together.
Commands that would otherwise conflict with each other (`-carousel` and `-single` or `-pics` and `-vids`) are canceled out.


``` bash
REQUIRED
-user       the user to scape
-dir        the directory to save files

OPTIONAL
-after      get the posts after a certain date
-before     get the posts before a certain date
-carousel   only download media from carousel posts
-has        only download a file if the post has this text
-overwrite  overwrite media that has already been saved
-pics       only download images
-single     only download media from single posts
-vids       only download videos
-timezone   timezone aka `America/Los_Angeles`
```


## To Do

* Add more filters.

Inspired by [insta-dl](https://github.com/sdushantha/insta-dl).
