package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// NudeNetClient calls the NudeNet inference service.
// Set BaseURL to the in-cluster address, e.g. http://nudenet.nudenet.svc.cluster.local:8080
type NudeNetClient struct {
	BaseURL string
}

type nudeNetResponse struct {
	Detections []struct {
		Class string  `json:"class"`
		Score float64 `json:"score"`
	} `json:"detections"`
}

var nsfwClasses = []string{
	"EXPOSED_GENITALIA_F", "EXPOSED_GENITALIA_M",
	"EXPOSED_BREAST_F", "EXPOSED_BUTTOCKS", "EXPOSED_ANUS",
}

const nsfwThreshold = 0.6

// IsNSFW uploads imagePath to the NudeNet service and returns true when any
// explicit detection exceeds nsfwThreshold.
func (c *NudeNetClient) IsNSFW(imagePath string) (bool, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return false, fmt.Errorf("open for moderation: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile("media", filepath.Base(imagePath))
	if err != nil {
		return false, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return false, err
	}
	w.Close()

	resp, err := http.Post(c.BaseURL+"/infer", w.FormDataContentType(), &buf)
	if err != nil {
		return false, fmt.Errorf("nudenet call: %w", err)
	}
	defer resp.Body.Close()

	var result nudeNetResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("nudenet decode: %w", err)
	}

	for _, d := range result.Detections {
		if d.Score < nsfwThreshold {
			continue
		}
		for _, cls := range nsfwClasses {
			if strings.EqualFold(d.Class, cls) {
				return true, nil
			}
		}
	}
	return false, nil
}
