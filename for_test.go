package youtubedataapi

import (
	"os"
)

func getKey() string {
	return os.Getenv("YoutubeAPIKey")
}
