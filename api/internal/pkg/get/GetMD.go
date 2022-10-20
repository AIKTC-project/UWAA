package get

import (
	"blog/internal/pkg/global"
	"context"
	"encoding/json"
	"time"
)

func GetMD(id int, enc *json.Encoder) {
	ctx, rel := context.WithTimeout(context.Background(), time.Second*10)
	defer rel()

	var md string

	global.DB.QueryRowContext(ctx, "SELECT md FROM content WHERE id = $1", id).Scan(&md)

	enc.Encode(global.Status{md})
}
