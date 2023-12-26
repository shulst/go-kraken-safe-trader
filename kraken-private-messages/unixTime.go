package kraken_private_messages

import (
	"strconv"
	"time"
)

type unixTime string

func (u unixTime) Time() time.Time {
	i, _ := strconv.ParseFloat(string(u), 64)
	return time.Unix(int64(i), 0)
}
