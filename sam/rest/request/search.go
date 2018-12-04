package request

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `search.go`, `search.util.go` or `search_test.go` to
	implement your API calls, helper functions and tests. The file `search.go`
	is only generated the first time, and will not be overwritten if it exists.
*/

import (
	"encoding/json"
	"github.com/crusttech/crust/internal/rbac"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

var _ = chi.URLParam
var _ = types.JSONText{}
var _ = multipart.FileHeader{}
var _ = rbac.Operation{}

// Search messages request parameters
type SearchMessages struct {
	InChannel uint64 `json:",string"`
	FromUser  uint64 `json:",string"`
	FirstID   uint64 `json:",string"`
	LastID    uint64 `json:",string"`
	Query     string
}

func NewSearchMessages() *SearchMessages {
	return &SearchMessages{}
}

func (s *SearchMessages) Fill(r *http.Request) (err error) {
	if strings.ToLower(r.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(r.Body).Decode(s)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return errors.Wrap(err, "error parsing http request body")
		}
	}

	if err = r.ParseForm(); err != nil {
		return err
	}

	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	if val, ok := get["inChannel"]; ok {

		s.InChannel = parseUInt64(val)
	}
	if val, ok := get["fromUser"]; ok {

		s.FromUser = parseUInt64(val)
	}
	if val, ok := get["firstID"]; ok {

		s.FirstID = parseUInt64(val)
	}
	if val, ok := get["lastID"]; ok {

		s.LastID = parseUInt64(val)
	}
	if val, ok := get["query"]; ok {

		s.Query = val
	}

	return err
}

var _ RequestFiller = NewSearchMessages()
