package api2

type ListFilesReq struct {
	Path         string `json:"path,omitempty"`
	JustNames    bool   `json:"justNames,omitempty"`
	Flat         bool   `json:"flat,omitempty"`
	CrossFS      bool   `json:"crossFS,omitempty"`
	Continuation string `json:"continuation,omitempty"`
	MaxFiles     int    `json:"maxFiles,omitempty"`
}

type ListFilesRes struct {
	Files        map[string]*FileInfo `json:"files,omitempty"`
	Names        []string             `json:"names,omitempty"`
	Continuation string               `json:"continuation,omitempty"`
	FileCount    int                  `json:"fileCount,omitempty"`
}

type FileInfo struct {
	ETAG  string `json:"etag,omitempty"`
	FS    string `json:"fs,omitempty"`
	Bytes int    `json:"bytes,omitempty"`
}
