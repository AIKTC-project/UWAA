package get

import (
	"blog_api/internal/pkg/global"
	"encoding/json"
	"net/http"
)

func SayHello(resp http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(resp)
	enc.Encode(global.Status{"Hi! sir kaise ho ?"})
}
