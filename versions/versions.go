package versions

import (
	"encoding/json"
	"io"
	"time"
)

type MCVersionsQueryResponse struct {
	Latest   Latest     `json:"latest"`
	Versions []Version `json:"versions"`
}
type Latest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}
type Version struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	URL         string    `json:"url"`
	Time        time.Time `json:"time"`
	ReleaseTime time.Time `json:"releaseTime"`
}

func GetVersions()  {
	
}

func parseVersionsFile(reader io.ReadCloser) (MCVersionsQueryResponse, error) {
	var parsedValue MCVersionsQueryResponse
	err := json.NewDecoder(reader).Decode(&parsedValue)
	return parsedValue, err
}
