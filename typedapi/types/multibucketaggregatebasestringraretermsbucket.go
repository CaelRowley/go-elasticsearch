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

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// MultiBucketAggregateBaseStringRareTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b2c13a00c152a97cb41193deda8ed9b37fd06796/specification/_types/aggregations/Aggregate.ts#L327-L329
type MultiBucketAggregateBaseStringRareTermsBucket struct {
	Buckets BucketsStringRareTermsBucket `json:"buckets"`
	Meta    Metadata                     `json:"meta,omitempty"`
}

func (s *MultiBucketAggregateBaseStringRareTermsBucket) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "buckets":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]StringRareTermsBucket, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Buckets", err)
				}
				s.Buckets = o
			case '[':
				o := []StringRareTermsBucket{}
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Buckets", err)
				}
				s.Buckets = o
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		}
	}
	return nil
}

// NewMultiBucketAggregateBaseStringRareTermsBucket returns a MultiBucketAggregateBaseStringRareTermsBucket.
func NewMultiBucketAggregateBaseStringRareTermsBucket() *MultiBucketAggregateBaseStringRareTermsBucket {
	r := &MultiBucketAggregateBaseStringRareTermsBucket{}

	return r
}
