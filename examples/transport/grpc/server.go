package test

import "context"

// 1.2.3.4:50051
// ¬false

type TicketService interface {
	// CRD
	Get(ctx context.Context, request GetTicketRequest) (response GetTicketResponse, err error)
	Create(ctx context.Context, request CreateTicketRequest) (response CreateTicketResponse, err error)
	Delete(ctx context.Context, request DeleteTicketRequest) (err error)
	// List
	List(ctx context.Context, request ListTicketsRequest) (response ListTicketsResponse, err error)
	// Persons
	SetAssignee(ctx context.Context, request SetAssigneeRequest) (response SetAssigneeResponse, err error)
	AddPerson(ctx context.Context, request AddPersonRequest) (response AddPersonResponse, err error)
	RemovePerson(ctx context.Context, request RemovePersonRequest) (response RemovePersonResponse, err error)
	// Update ticket details
	Update(ctx context.Context, request UpdateRequest) (response UpdateResponse, err error)
	Close(ctx context.Context, request CloseRequest) (err error)
	UpdateReference(ctx context.Context, request UpdateReferenceRequest) (response UpdateReferenceResponse, err error)
	UpdateType(ctx context.Context, request UpdateTypeRequest) (response UpdateTypeResponse, err error)
	UpdateStatus(ctx context.Context, request UpdateStatusRequest) (response UpdateStatusResponse, err error)
}

// Ticket shared model
type Ticket struct {
	Id              string
	Key             string
	AccountId       int64
	Reference       Reference
	Source          Source
	Type            Type
	Status          Status
	Title           string
	Description     string
	MediaIds        []string
	AppointmentIds  []string
	Priority        Priority
	DueDate         int64
	CreationDate    int64
	Assignee        Assignee
	InvolvedPersons []InvolvedPerson
}

// Reference shared model
type Reference struct {
	ReferenceType ReferenceType
	ReferenceId   string
}

// Assignee shared model
type Assignee struct {
	UserReference string
	AccountId     int64
}

// InvolvedPerson shared model
type InvolvedPerson struct {
	UserReference string
	AccountId     int64
}

// TicketCreated shared model
type TicketCreated struct {
	UserId    string
	AccountId int64
	Ticket    Ticket
	Timestamp int64
}

// TicketReferenceChanged shared model
type TicketReferenceChanged struct {
	UserId       string
	AccountId    int64
	Ticket       Ticket
	OldReference Reference
	Timestamp    int64
}

// TicketAssigneeChanged shared model
type TicketAssigneeChanged struct {
	UserId      string
	AccountId   int64
	Ticket      Ticket
	OldAssignee Assignee
	Timestamp   int64
}

// TicketPersonAdded shared model
type TicketPersonAdded struct {
	UserId         string
	AccountId      int64
	Ticket         Ticket
	InvolvedPerson InvolvedPerson
	Timestamp      int64
}

// TicketPersonRemoved shared model
type TicketPersonRemoved struct {
	UserId         string
	AccountId      int64
	Ticket         Ticket
	InvolvedPerson InvolvedPerson
	Timestamp      int64
}

// TicketTypeChanged shared model
type TicketTypeChanged struct {
	UserId    string
	AccountId int64
	Ticket    Ticket
	OldType   Type
	Timestamp int64
}

// TicketStatusChanged shared model
type TicketStatusChanged struct {
	UserId    string
	AccountId int64
	Ticket    Ticket
	OldStatus Status
	Timestamp int64
}

// ReferenceType enumeration
type ReferenceType int32

const (
	ReferenceType_WITHOUT  = 0
	ReferenceType_ADVERT   = 1
	ReferenceType_BUILDING = 2
	ReferenceType_GROUP    = 3
)

var ReferenceType_name = map[int32]string{
	0: "WITHOUT",
	1: "ADVERT",
	2: "BUILDING",
	3: "GROUP",
}

var ReferenceType_value = map[string]int32{
	"WITHOUT":  0,
	"ADVERT":   1,
	"BUILDING": 2,
	"GROUP":    3,
}

// Source enumeration
type Source int32

const (
	Source_UNDEFINED_SOURCE = 0
	Source_MAIL             = 1
	Source_PHONE            = 2
	Source_OTHER            = 3
)

var Source_name = map[int32]string{
	0: "UNDEFINED_SOURCE",
	1: "MAIL",
	2: "PHONE",
	3: "OTHER",
}

var Source_value = map[string]int32{
	"UNDEFINED_SOURCE": 0,
	"MAIL":             1,
	"PHONE":            2,
	"OTHER":            3,
}

// Status enumeration
type Status int32

const (
	Status_OPEN        = 0
	Status_IN_PROGRESS = 1
	Status_DONE        = 2
)

var Status_name = map[int32]string{
	0: "OPEN",
	1: "IN_PROGRESS",
	2: "DONE",
}

var Status_value = map[string]int32{
	"OPEN":        0,
	"IN_PROGRESS": 1,
	"DONE":        2,
}

// Type enumeration
type Type int32

const (
	Type_UNDEFINED_TYPE = 0
	Type_DAMAGE         = 1
)

var Type_name = map[int32]string{
	0: "UNDEFINED_TYPE",
	1: "DAMAGE",
}

var Type_value = map[string]int32{
	"UNDEFINED_TYPE": 0,
	"DAMAGE":         1,
}

// Priority enumeration
type Priority int32

const (
	Priority_DEFAULT = 0
	Priority_LOW     = 1
	Priority_HIGH    = 2
)

var Priority_name = map[int32]string{
	0: "DEFAULT",
	1: "LOW",
	2: "HIGH",
}

var Priority_value = map[string]int32{
	"DEFAULT": 0,
	"LOW":     1,
	"HIGH":    2,
}

type GetTicketRequest struct {
	Id string
}

type ListTicketsRequest struct {
	AccountId int64
}

type CreateTicketRequest struct {
	AccountId       int64
	Reference       Reference
	Source          Source
	Type            Type
	Status          Status
	Title           string
	Description     string
	MediaIds        []string
	AppointmentIds  []string
	Priority        Priority
	DueDate         int64
	Assignee        Assignee
	InvolvedPersons []InvolvedPerson
}

type UpdateRequest struct {
	Id          string
	Title       string
	Description string
}

type CloseRequest struct {
	Id string
}

type UpdateReferenceRequest struct {
	Id        string
	Reference Reference
}

type DeleteTicketRequest struct {
	Id string
}

type SetAssigneeRequest struct {
	Id       string
	Assignee Assignee
}

type AddPersonRequest struct {
	Id     string
	Person InvolvedPerson
}

type RemovePersonRequest struct {
	Id     string
	Person InvolvedPerson
}

type UpdateTypeRequest struct {
	Id   string
	Type Type
}

type UpdateStatusRequest struct {
	Id     string
	Status Status
}

type GetTicketResponse struct {
	Ticket Ticket
}

type ListTicketsResponse struct {
	Tickets []Ticket
}

type CreateTicketResponse struct {
	Ticket Ticket
}

type UpdateResponse struct {
	Ticket Ticket
}

type CloseResponse struct {
}

type UpdateReferenceResponse struct {
	Ticket Ticket
}

type DeleteTicketResponse struct {
}

type SetAssigneeResponse struct {
	Ticket Ticket
}

type AddPersonResponse struct {
	Ticket Ticket
}

type RemovePersonResponse struct {
	Ticket Ticket
}

type UpdateTypeResponse struct {
	Ticket Ticket
}

type UpdateStatusResponse struct {
	Ticket Ticket
}
