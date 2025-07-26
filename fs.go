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

type ListFileSystemsRes struct {
	FileSystems []*FileSystem          `json:"fileSystems"`
	FSMap       map[string]*FileSystem `json:"fsMap,omitempty"`
}

type FileSystem struct {
	ID          string   `json:"id"`
	Parent      string   `json:"parent,omitempty"`
	Ancestors   []string `json:"ancestors,omitempty"`
	Children    []string `json:"children,omitempty"`
	Created     string   `json:"created"`
	LastMounted string   `json:"lastMounted"`
	IsMounted   bool     `json:"isMounted"`
	Frozen      string   `json:"frozen,omitempty"`
	Bytes       int      `json:"bytes"`
}
