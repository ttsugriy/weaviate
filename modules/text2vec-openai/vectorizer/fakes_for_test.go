//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package vectorizer

import (
	"context"

	"github.com/weaviate/weaviate/modules/text2vec-openai/ent"
)

type fakeClient struct {
	lastInput  string
	lastConfig ent.VectorizationConfig
}

func (c *fakeClient) Vectorize(ctx context.Context,
	text string, cfg ent.VectorizationConfig,
) (*ent.VectorizationResult, error) {
	c.lastInput = text
	c.lastConfig = cfg
	return &ent.VectorizationResult{
		Vector:     []float32{0, 1, 2, 3},
		Dimensions: 4,
		Text:       text,
	}, nil
}

func (c *fakeClient) VectorizeQuery(ctx context.Context,
	text string, cfg ent.VectorizationConfig,
) (*ent.VectorizationResult, error) {
	c.lastInput = text
	c.lastConfig = cfg
	return &ent.VectorizationResult{
		Vector:     []float32{0.1, 1.1, 2.1, 3.1},
		Dimensions: 4,
		Text:       text,
	}, nil
}

type fakeSettings struct {
	skippedProperty    string
	vectorizeClassName bool
	excludedProperty   string
	openAIType         string
	openAIModel        string
	openAIModelVersion string
	resourceName       string
	deploymentID       string
	isAzure            bool
}

func (f *fakeSettings) PropertyIndexed(propName string) bool {
	return f.skippedProperty != propName
}

func (f *fakeSettings) VectorizePropertyName(propName string) bool {
	return f.excludedProperty != propName
}

func (f *fakeSettings) VectorizeClassName() bool {
	return f.vectorizeClassName
}

func (f *fakeSettings) Type() string {
	return f.openAIType
}

func (f *fakeSettings) Model() string {
	return f.openAIModel
}

func (f *fakeSettings) ModelVersion() string {
	return f.openAIModelVersion
}

func (f *fakeSettings) ResourceName() string {
	return f.resourceName
}

func (f *fakeSettings) DeploymentID() string {
	return f.deploymentID
}

func (f *fakeSettings) IsAzure() bool {
	return f.isAzure
}
