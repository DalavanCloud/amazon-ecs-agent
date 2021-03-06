// Copyright 2014-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package pause

import (
	"context"

	"github.com/aws/amazon-ecs-agent/agent/config"
	"github.com/aws/amazon-ecs-agent/agent/dockerclient/dockerapi"
	"github.com/docker/docker/api/types"
)

// Loader defines an interface for loading the pause container image. This is mostly
// to facilitate mocking and testing of the LoadImage method
type Loader interface {
	LoadImage(ctx context.Context, cfg *config.Config, dockerClient dockerapi.DockerClient) (*types.ImageInspect, error)
}

type loader struct{}

// New creates a new pause image loader
func New() Loader {
	return &loader{}
}
