package create

import (
	"blog/internal/pkg/global"
	"encoding/json"
	"net/http"
)

type cc struct {
	Title    string `json:"title"`
	Markdown string `json:"md"`
}

func CreateBlog(req *http.Request, enc *json.Encoder) {
	usr_id := global.VerifyUser(req, enc, nil)
	if usr_id < 1 {
		enc.Encode(global.Status{"ltc"})
		return
	}

	var p cc
	json.NewDecoder(req.Body).Decode(&p)

	_, err := global.DB.Exec("INSERT INTO content(md, _from, title) VALUES($1, $2, $3)", p.Markdown, usr_id, p.Title)
	if err != nil {
		enc.Encode(global.Status{"An error has occured"})
		return
	}

	enc.Encode(global.Status{"sfl"})
}
