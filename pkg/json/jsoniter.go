// +build jsoniter

package json

import "github.com/json-iterator/go"

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Marshal json marshal
	Marshal = json.Marshal
	// MarshalIndent json marshal indent
	MarshalIndent = json.MarshalIndent
	// NewDecoder json new decoder function
	NewDecoder = json.NewDecoder
	// NewEncoder json new encoder function
	NewEncoder = json.NewEncoder
)
