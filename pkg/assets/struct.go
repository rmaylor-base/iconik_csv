package cmd

type Assets struct {
	Objects []*Object `json:"objects"`
}

type Object struct {
	DateCreated        string         `json:"date_created"`
	DataModified       string         `json:"date_modified"`
	FileNames          []string       `json:"file_names"`
	Files              []*File        `json:"files"`
	Formats            []*Format      `json:"formats"`
	FrameRate          float32        `json:"frame_rate"`
	ID                 string         `json:"id"`
	InCollections      []string       `json:"in_collections"`
	Keyframes          []*Keyframe    `json:"keyframes"`
	OriginalResolution map[string]int `json:"original_resolution"`
}

type File struct {
	OriginalName string `json:"original_name"`
}

type Keyframe struct {
	URL string `json:"url"`
}

type Format struct {
	Status string `json:"status"`
}
