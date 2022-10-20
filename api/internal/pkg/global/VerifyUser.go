package global

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func VerifyUser(req *http.Request, enc *json.Encoder, username_ptr *string) (userID int) {
	cookie, _ := req.Cookie("data")

	if cookie == nil {
		if enc != nil {
			enc.Encode(Status{"ltc"})
		}
		return 0
	}

	ctx, rel := context.WithTimeout(req.Context(), 5*time.Second)
	defer rel()

	usr_id, err := strconv.Atoi(RDB.Get(ctx, "ssn:"+cookie.Value).Val())
	if err != nil {
		log.Println(err)
		return 0
	}

	// Verifies the user with session-id
	if usr_id < 1 {
		if enc != nil {
			enc.Encode(Status{"ltc"})
		}
		return 0
	}

	if username_ptr != nil {
		*username_ptr = RDB.HGet(ctx, "user:"+strconv.Itoa(usr_id), "uname").Val()
	}

	return usr_id
}
