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
