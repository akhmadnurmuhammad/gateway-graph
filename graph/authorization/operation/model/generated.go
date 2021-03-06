// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package operation

type FetchRequestOperationInput struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

type FetchResponseOperation struct {
	IsSuccess *bool       `json:"isSuccess"`
	Message   *string     `json:"message"`
	Status    *int        `json:"status"`
	Data      *Operations `json:"data"`
}

type FindByIDOperationRequestInput struct {
	ID *string `json:"iD"`
}

type MainResponseOperation struct {
	IsSuccess *bool              `json:"isSuccess"`
	Message   *string            `json:"message"`
	Status    *int               `json:"status"`
	Data      *OperationResponse `json:"data"`
}

type OperationRequestInput struct {
	ID     *string `json:"iD"`
	Name   *string `json:"name"`
	URL    *string `json:"url"`
	Path   *string `json:"path"`
	Method *string `json:"method"`
}

type OperationResponse struct {
	ID        *string `json:"iD"`
	Name      *string `json:"name"`
	URL       *string `json:"url"`
	Path      *string `json:"path"`
	Method    *string `json:"method"`
	CreatedBy *string `json:"createdBy"`
	UpdatedBy *string `json:"updatedBy"`
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
}

type Operations struct {
	Items []*OperationResponse `json:"items"`
}
