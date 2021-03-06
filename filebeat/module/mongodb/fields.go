// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mongodb

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "mongodb", asset.ModuleFieldsPri, AssetMongodb); err != nil {
		panic(err)
	}
}

// AssetMongodb returns asset data.
// This is the base64 encoded gzipped contents of module/mongodb.
func AssetMongodb() string {
	return "eJyUkdFO6zAMhu/7FL92vz1AL4502LS7wjOE1g3WkjhK0rHy9CjtJFoIBXxp6/f3Wd7jQmMNK05L91wBiZOhGrt7Z1cBHcU2sE8srsa/CgAa6QZD6CXAqxDZaTQ5cHqAEY2eDcVDBfRMpov1lNnDKUtLVq40eqqhgwz+3ing5jpPy9AHsUvaBMq1hC2BRvRiy1fgJjTXUVxS7OKd8L1ByeLDoxXrxZFLq/WbZOA8uDbPlEGrEmkJ/KZyA9LDUoxK0ypEN2V9fuHxqWn+P55Ww/n4C42vErqiokt0+5PgcY78qMOOk3Kd4ZjI/UKqYBfpSoHTWEgrwyqu+l6ll+n3B0NXMquZZR3UfFAKAxVQpVO2SaXEZ857AAAA//9KC+uk"
}
