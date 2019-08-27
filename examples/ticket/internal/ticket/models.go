package ticket

/*
The models are based off the specified messages in protobuf. You can use them
as a base for the domain's hidden models as they will most likely correlate with the
shared models.
Godin will not render this file when calling 'godin generate', only if it doesn't exist.
*/

type Ticket struct {
	Id     string
	Title  string
	Status Status
}

// Status enumeration
type Status int32

const (
	Status_OPEN        = 0
	Status_IN_PROGRESS = 1
	Status_CLOSED      = 2
)

var StatusName = map[int32]string{
	0: "OPEN",
	1: "IN_PROGRESS",
	2: "CLOSED",
}

var StatusValue = map[string]int32{
	"OPEN":        0,
	"IN_PROGRESS": 1,
	"CLOSED":      2,
}
