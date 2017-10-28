# instago
[![Go Report Card](https://goreportcard.com/badge/github.com/pagumin/instago)](https://goreportcard.com/report/github.com/pagumin/instago)
[![GitHub release](https://img.shields.io/github/release/qubyte/rubidium.svg)](https://github.com/pagumin/instago/releases)

Scrape an Instagram user with Go. Get the latest release [here](https://github.com/pagumin/instago/releases).

## Usage

Default usage:

``` bash
instago -user <user> -dir <save_directory>
```

Only download images:

``` bash
instago -user <user> -dir <save_directory> -pics
```

Only download videos:

``` bash
instago -user <user> -dir <save_directory> -vids
```

Download content from past 5 days:

``` bash
instago -user <user> -dir <save_directory> -maxdaysold 5
```

Using `-pics` and `-vids` together will download images and videos together.


## To Do

* Add filters

Inspired by [insta-dl](https://github.com/sdushantha/insta-dl).
