package main

import (
	"blog/internal/pkg/create"
	"blog/internal/pkg/get"
	"blog/internal/pkg/global"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func route(resp http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(resp)

	switch req.Method {
	case "GET":
		url_query, _ := url.ParseQuery(req.URL.RawQuery)

		if req.URL.RawQuery != "" {
			switch url_query.Get("q") {
			case "id":
				id, _ := strconv.Atoi(url_query.Get("v"))
				get.GetMD(id, enc)
				break
			case "blog":
				id, _ := strconv.Atoi(url_query.Get("v"))
				get.GetAll(id, enc)
				break
			default:
				enc.Encode(global.Status{"Invalid query"})
				break
			}
		} else {
			get.GetPreview(req, enc)
		}
		break
	case "POST":
		create.CreateBlog(req, enc)
		break
	default:
		enc.Encode(global.Status{"Invalid query"})
		break
	}
}
