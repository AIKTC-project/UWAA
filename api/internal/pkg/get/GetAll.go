package get

import (
	"blog/internal/pkg/global"
	"context"
	"encoding/json"
	"time"
)

func GetAll(id int, enc *json.Encoder) {
	ctx, rel := context.WithTimeout(context.Background(), time.Second*10)
	defer rel()

	rows, err := global.DB.QueryContext(ctx, "SELECT id, title, _from, md FROM content WHERE id = $1", id)
	global.HandleErr(err)

	m := make(map[string][]interface{})
	j := json.RawMessage{}

	for rows.Next() {
		var id int
		var md, title string
		var _from [2]string

		rows.Scan(&id, &title, &_from[0], &md)
		_from[1] = global.RDB.HGet(ctx, "user:"+_from[0], "uname").Val()

		m["id"] = append(m["id"], id)
		m["_from"] = append(m["_from"], _from)
		m["title"] = append(m["title"], title)
		m["md"] = append(m["md"], md)
	}
	j, _ = json.Marshal(m)
	enc.Encode(j)
}
