package tpl

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/CloudyKit/jet/v6/loaders/httpfs"
)

func NewPrefixedFilesystem(fsys fs.FS, prefix string) (jet.Loader, error) {
	return httpfs.NewLoader(http.FS(prefixedFilesystem{fsys, prefix}))
}

type prefixedFilesystem struct {
	fsys   fs.FS
	prefix string
}

func (p prefixedFilesystem) Open(name string) (fs.File, error) {
	return p.fsys.Open(fmt.Sprintf("%s/%s", p.prefix, name))
}
