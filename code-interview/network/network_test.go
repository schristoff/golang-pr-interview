package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadToFile(t *testing.T) {

	files := map[string]string{
		"README.md": "Hello World\n",
		".adr-dir":  "adr\n",
	}
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, fileName := path.Split(req.URL.Path)
		content, ok := files[fileName]
		if !ok {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		//nolint:errcheck // ignore
		rw.Write([]byte(content))
	}))

	tests := []struct {
		name     string
		fileName string
	}{
		{
			name:     "existing file",
			fileName: "README.md",
		},
		{
			name:     "non existing file",
			fileName: "README.md.bad",
		},
		{
			name:     "existing file with shasum",
			fileName: ".adr-dir",
		},
		{
			name:     "existing file with shasum and query parameters",
			fileName: ".adr-dir",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			src := fmt.Sprintf("%s/%s", srv.URL, tt.fileName)
			fmt.Println(src)
			dst := filepath.Join(t.TempDir(), tt.fileName)
			err := DownloadToFile(src, dst, "")

			require.NoError(t, err)
			require.FileExists(t, dst)
			require.NoError(t, err)
		})
	}
}
