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

// CumulativeCardinalityAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/Aggregate.ts#L698-L706
type CumulativeCardinalityAggregate struct {
	Meta          *Metadata `json:"meta,omitempty"`
	Value         int64     `json:"value"`
	ValueAsString *string   `json:"value_as_string,omitempty"`
}

// CumulativeCardinalityAggregateBuilder holds CumulativeCardinalityAggregate struct and provides a builder API.
type CumulativeCardinalityAggregateBuilder struct {
	v *CumulativeCardinalityAggregate
}

// NewCumulativeCardinalityAggregate provides a builder for the CumulativeCardinalityAggregate struct.
func NewCumulativeCardinalityAggregateBuilder() *CumulativeCardinalityAggregateBuilder {
	r := CumulativeCardinalityAggregateBuilder{
		&CumulativeCardinalityAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the CumulativeCardinalityAggregate struct
func (rb *CumulativeCardinalityAggregateBuilder) Build() CumulativeCardinalityAggregate {
	return *rb.v
}

func (rb *CumulativeCardinalityAggregateBuilder) Meta(meta *MetadataBuilder) *CumulativeCardinalityAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *CumulativeCardinalityAggregateBuilder) Value(value int64) *CumulativeCardinalityAggregateBuilder {
	rb.v.Value = value
	return rb
}

func (rb *CumulativeCardinalityAggregateBuilder) ValueAsString(valueasstring string) *CumulativeCardinalityAggregateBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}