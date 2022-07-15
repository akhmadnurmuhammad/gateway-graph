// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package module

type FetchRequestModuleInput struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

type FetchResponseModule struct {
	IsSuccess *bool    `json:"isSuccess"`
	Message   *string  `json:"message"`
	Status    *int     `json:"status"`
	Data      *Modules `json:"data"`
}

type FindByIDModuleRequestInput struct {
	ID *string `json:"iD"`
}

type MainResponseModule struct {
	IsSuccess *bool           `json:"isSuccess"`
	Message   *string         `json:"message"`
	Status    *int            `json:"status"`
	Data      *ModuleResponse `json:"data"`
}

type ModuleRequestInput struct {
	ID   *string `json:"iD"`
	Name *string `json:"name"`
	Path *string `json:"path"`
}

type ModuleResponse struct {
	ID        *string `json:"iD"`
	Name      *string `json:"name"`
	Path      *string `json:"path"`
	CreatedBy *string `json:"createdBy"`
	UpdatedBy *string `json:"updatedBy"`
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
}

type Modules struct {
	Items []*ModuleResponse `json:"items"`
}
