package versions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Eldius/minecraft-manager-go/config"
)

type MCVersionsQueryResponse struct {
	Latest   Latest    `json:"latest"`
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

func GetVersions() ([]Version, error) {
	res, err := http.Get(config.GetMojangVersionsURL())
	if err != nil {
		fmt.Println("Failed to request versions\n---\n", err.Error())
		return make([]Version, 0), err
	}
	defer res.Body.Close()
	body, err := parseVersionsFile(res.Body)
	if err != nil {
		fmt.Println("Failed to parse versions response\n---\n", err.Error())
		return make([]Version, 0), err
	}
	return body.Versions, err
}

func parseVersionsFile(reader io.ReadCloser) (MCVersionsQueryResponse, error) {
	var parsedValue MCVersionsQueryResponse
	err := json.NewDecoder(reader).Decode(&parsedValue)
	return parsedValue, err
}
