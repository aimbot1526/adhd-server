package response

import "time"

type ProductResponse struct {
	Name        string
	Description string
	Price       int
	Created_At  time.Time
	Updated_At  time.Time
}
