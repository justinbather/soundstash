package ports

import "os"

type Downloader interface {
	Download(url string) (*os.File, error)
}
