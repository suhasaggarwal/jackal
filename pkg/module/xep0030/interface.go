// Copyright 2020 The jackal Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xep0030

import (
	"context"

	c2smodel "github.com/ortuman/jackal/pkg/model/c2s"

	"github.com/ortuman/jackal/pkg/component"
	"github.com/ortuman/jackal/pkg/module"
	"github.com/ortuman/jackal/pkg/router"
	"github.com/ortuman/jackal/pkg/storage/repository"
)

//go:generate moq -out roster_repository.mock_test.go . rosterRepository
type rosterRepository interface {
	repository.Roster
}

//go:generate moq -out router.mock_test.go . globalRouter:routerMock
type globalRouter interface {
	router.Router
}

//go:generate moq -out resourcemanager.mock_test.go . resourceManager
type resourceManager interface {
	GetResources(ctx context.Context, username string) ([]c2smodel.Resource, error)
}

//go:generate moq -out modules.mock_test.go . modules
type modules interface {
	AllModules() []module.Module
}

//go:generate moq -out components.mock_test.go . components
type components interface {
	AllComponents() []component.Component
}

//go:generate moq -out component.mock_test.go . discoComponent:componentMock
type discoComponent interface {
	component.Component
}

//go:generate moq -out module.mock_test.go . discoModule:moduleMock
type discoModule interface {
	module.Module
}
