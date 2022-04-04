// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
