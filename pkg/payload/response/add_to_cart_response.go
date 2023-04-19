package response

import "time"

type CartItem struct {
	Quantity        int
	Created_At      time.Time
	Updated_At      time.Time
	ProductResponse ProductResponse
}
