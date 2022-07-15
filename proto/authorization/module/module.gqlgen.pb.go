package module

import (
	context "context"
)

type ModuleResolvers struct {
	Service ModuleServer
}

func (s *ModuleResolvers) ModuleFetch(ctx context.Context, in *FetchRequestModule) (*FetchResponseModule, error) {
	return s.Service.Fetch(ctx, in)
}
func (s *ModuleResolvers) ModuleStore(ctx context.Context, in *ModuleRequest) (*MainResponseModule, error) {
	return s.Service.Store(ctx, in)
}
func (s *ModuleResolvers) ModuleFindByID(ctx context.Context, in *FindByIDModuleRequest) (*MainResponseModule, error) {
	return s.Service.FindByID(ctx, in)
}
func (s *ModuleResolvers) ModuleUpdate(ctx context.Context, in *ModuleRequest) (*MainResponseModule, error) {
	return s.Service.Update(ctx, in)
}
func (s *ModuleResolvers) ModuleDelete(ctx context.Context, in *FindByIDModuleRequest) (*MainResponseModule, error) {
	return s.Service.Delete(ctx, in)
}

type FetchRequestModuleInput = FetchRequestModule
type MainResponseModuleInput = MainResponseModule
type FetchResponseModuleInput = FetchResponseModule
type ModulesInput = Modules
type ModuleResponseInput = ModuleResponse
type ModuleRequestInput = ModuleRequest
type FindByIDModuleRequestInput = FindByIDModuleRequest
