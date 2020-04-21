package mimestypes

import "strings"

var MimeType map[string]string

func init() {
	MimeType = map[string]string{
		"map":  "application/x-navimap",
		"js":   "application/x-javascript",
		"css":  "text/css",
		"html": "text/html",
		"txt":  "text/plain",
		"ico":  "image/x-icon",
		"png":  "image/png",
		"json": "application/json",
	}
}

// GetFileContentType Get the file MIME using his File Extension
func GetFileContentType(path string) (string, error) {
	contentType := MimeType[getTheFileExtension(path)]
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return contentType, nil
}

//getTheFileExtension get the fileExtention of the requested resource
func getTheFileExtension(path string) string {
	exte := strings.Split(path, ".")
	if len(exte) > 1 {
		return exte[len(exte)-1]
	}
	return ""
}
