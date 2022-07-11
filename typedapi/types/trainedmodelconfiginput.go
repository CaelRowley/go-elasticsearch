// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// TrainedModelConfigInput type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/TrainedModel.ts#L184-L187
type TrainedModelConfigInput struct {
	// FieldNames An array of input field names for the model.
	FieldNames []Field `json:"field_names"`
}

// TrainedModelConfigInputBuilder holds TrainedModelConfigInput struct and provides a builder API.
type TrainedModelConfigInputBuilder struct {
	v *TrainedModelConfigInput
}

// NewTrainedModelConfigInput provides a builder for the TrainedModelConfigInput struct.
func NewTrainedModelConfigInputBuilder() *TrainedModelConfigInputBuilder {
	r := TrainedModelConfigInputBuilder{
		&TrainedModelConfigInput{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelConfigInput struct
func (rb *TrainedModelConfigInputBuilder) Build() TrainedModelConfigInput {
	return *rb.v
}

// FieldNames An array of input field names for the model.

func (rb *TrainedModelConfigInputBuilder) FieldNames(field_names ...Field) *TrainedModelConfigInputBuilder {
	rb.v.FieldNames = field_names
	return rb
}