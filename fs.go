package api2

type ListFilesReq struct {
	Path         string `json:"path,omitempty"`
	JustNames    bool   `json:"just_names,omitempty"`
	Flat         bool   `json:"flat,omitempty"`
	CrossFS      bool   `json:"cross_fs,omitempty"`
	Continuation string `json:"continuation,omitempty"`
	MaxFiles     int    `json:"max_files,omitempty"`
}

type ListFilesRes struct {
	Message      string               `json:"message,omitempty"`
	Files        map[string]*FileInfo `json:"files,omitempty"`
	Names        []string             `json:"names,omitempty"`
	Continuation string               `json:"continuation,omitempty"`
	FileCount    int                  `json:"file_count,omitempty"`
}

type FileInfo struct {
	ETAG  string `json:"etag,omitempty"`
	FS    string `json:"fs,omitempty"`
	Bytes int    `json:"bytes,omitempty"`
}
