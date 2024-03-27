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
// https://github.com/elastic/elasticsearch-specification/tree/b2c13a00c152a97cb41193deda8ed9b37fd06796

package types

// TrainedModelConfigMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b2c13a00c152a97cb41193deda8ed9b37fd06796/specification/ml/_types/TrainedModel.ts#L207-L215
type TrainedModelConfigMetadata struct {
	// FeatureImportanceBaseline An object that contains the baseline for feature importance values. For
	// regression analysis, it is a single value. For classification analysis, there
	// is a value for each class.
	FeatureImportanceBaseline map[string]string `json:"feature_importance_baseline,omitempty"`
	// Hyperparameters List of the available hyperparameters optimized during the
	// fine_parameter_tuning phase as well as specified by the user.
	Hyperparameters []Hyperparameter `json:"hyperparameters,omitempty"`
	ModelAliases    []string         `json:"model_aliases,omitempty"`
	// TotalFeatureImportance An array of the total feature importance for each feature used from the
	// training data set. This array of objects is returned if data frame analytics
	// trained the model and the request includes total_feature_importance in the
	// include request parameter.
	TotalFeatureImportance []TotalFeatureImportance `json:"total_feature_importance,omitempty"`
}

// NewTrainedModelConfigMetadata returns a TrainedModelConfigMetadata.
func NewTrainedModelConfigMetadata() *TrainedModelConfigMetadata {
	r := &TrainedModelConfigMetadata{
		FeatureImportanceBaseline: make(map[string]string, 0),
	}

	return r
}
