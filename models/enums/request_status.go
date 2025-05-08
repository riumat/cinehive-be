package enums

type RequestStatus string

const (
	Pending  RequestStatus = "pending"
	Accepted RequestStatus = "accepted"
	Rejected RequestStatus = "rejected"
)
