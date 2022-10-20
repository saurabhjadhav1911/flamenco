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
