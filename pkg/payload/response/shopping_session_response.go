package response

import "time"

type ShoppingSessionResponse struct {
	Total      int
	Created_At time.Time
	Updated_At time.Time
	UserID     uint
}
