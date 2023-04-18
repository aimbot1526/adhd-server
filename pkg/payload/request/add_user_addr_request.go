package request

type AddUserAddrRequest struct {
	Address_line1 string
	Address_line2 string
	City          string
	PostalCode    int
	Country       string
	Mobile        int
}
