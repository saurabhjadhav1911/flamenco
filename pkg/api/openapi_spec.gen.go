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

	"H4sIAAAAAAAC/9xb3W4ctxV+FWJSIAk6uytbboHuVR07dmT4R8jKyEUsrLjDszuUOOSE5Ox6awjIQ/RN",
	"2gC9aK76AsobFYfk/M9qpVhykwZBMNohD8/vd36G+RAlKsuVBGlNNP0QmSSFjLrHx8bwlQR2Qs0F/s3A",
	"JJrnlisZTVtvCTeEEotP1BBu8W8NCfA1MLLYEpsC+U7pC9DjKI5yrXLQloM7JVFZRiVzz9xC5h7+oGEZ",
	"TaPPJjVzk8DZ5InfEF3Gkd3mEE0jqjXd4t/naoG7w8/Gai5X4fd5rrnS3G4bC7i0sAJdrvC/DmyXNBt+",
	"cT1NY6kt9oqD+pv5lSgRNRe7GSkKzvDFUumM2mjqf4i7Cy/jSMMPBdfAoun35SJUTpCl4q0hQkdLDZU0",
	"uYpre51W56rFOSQWGXy8plzQhYAXajEDa5GdnufMuFwJIMa/J2pJKHmhFgSpmQEHSRVP/GObzncpSLLi",
	"a5AxETzj1vnZmgrO8L8FGGIV/maABCJj8kaKLSkM8kg23KbEK80djmdXLthTftfZGCxpIWyfr5MUSHjp",
	"+SAmVRsZmCGFAU02yDsDCzrj0p2fclOqZIzkgXGLXHr64aglFQbivh5sChrpUyHUhuDWLk1ClxbXpEDO",
	"1YKk1JAFgCSmWGTcWmBj8p0qBCM8y8WWMBDgtwlB4D03niA1F4Yslfakz9UiJlQyjHWV5VzgGm7H72Tt",
	"kwulBFCJEl3Atq+sIwbS8iUHHehWjhGTrDCWLIAUkv9QeHNxWYlQWqxnqNr3b6E5nmXAOLUgtkQD+jOh",
	"7hgGSy45bojRVZ3geGTs+FGF9T/lVFueFILqyoo71GCKRRng1+HCQCjNws7KGW9N4SRsX3PDu75ldXGd",
	"gtCH2x4VbPH2yIcwKqv0Jk2+EPwCCCVfCZAMNKGMjZT8ckxmYJHcmTPImQ8EnzGoJIieWlJRnWFTavHo",
	"QjD5uXOGKpZAMhdLZljRHRBE5wuLbghcs9pOHfwqFiN8493BO2Npc/Kk0BqkFVuiEGloSdd5dwNrzJic",
	"ffN49s3XT+fPjl5+PT9+fPLNmc+jjGtIrNJbklObkj+Ss3fR5DP3z7vojNA8R5UyLzbIIkP5llzAHNdH",
	"ccS4Lh/dzwHzU2pSYPN65elA8Oxymj7KBQ00pG9ErAdYasjR02OP5lsnNjpNcIkxea2IBGOBoWKKxBYa",
	"DPnCAayJCeMJHkU1B/MloRqIKfJcadsVPTAfY+49fIhCC0VtFDtf2CvksHRlPqrP9HUMN+QVlXQF2iMf",
	"ty70aYZQPpC8BF2AuF1REZR584JoKOn28lUnHIJLePYaZ+6LDdTWQCp+yY0tncF592699XVUFhq/TuKT",
	"FiLuELc+YkjAsqLsiRVeEA25BoMsEEqML19CHeSQ6D0khYV9le6NLN5hbths15rra62VRlLdOptBq3Ys",
	"o6VfuGZgDF0N8dphx9Gs1w9x88KX5FSIN8to+v31dp2VxQjuuox7ImigFobshC+4ksTyDIylWY4oUArK",
	"qIURvhkqFvgAubdvj56W4P7CVc17Cu6b1voYoFWpX+TsjqXpWMdxWuqsPq9i9vTy1BvoFVjKqKXOUIy5",
	"YoeK45buexJ3ukG94FZTvSVZIBaSnRmTV0q7cMkFvG8ifUIl5opMYbHpcKLA2CJndLwYJ2dEKuv1UBaG",
	"F7DFqIL3FGkFF3eONo1mueYWyDPNVyliP1YGY8goF8j1dqFB/nUREo/Sq3KFj8lo5haQmf3Pv9cgGnDS",
	"cuRZI06H9eRrqMG9lYOUaYsmlq9dR0VlghrwzVUuwIZn6ZXFlRwtKfcrqoecFsY9/FBA4R6oTlLsuKtH",
	"nxU9+RF6hku2gUjrB/fsqRSoolHz8CiONtR1FKOl0iOsH8xgWv0WVtxY0MA8BPZBiDKmwQw7lKDGzp1S",
	"2h11I2Xy5GJ3Ly6oxSAZRli1tBuqd8DvjWLXi1SHb5Xg5lV33E5gexvIj+rmK13ElVKbXX2pjDhKfEHq",
	"uIy6Wm5oZodEQ5g+g6TQ3G53ZJobp4/r8kYrFQyWZ3VjVjexmI2fCZqBTFQHKrIGyN0fbIQXh1d/J7/8",
	"ePXT1c9X/7z66Zcfr/519fPVP5rjlumfDtqJP5wyTzIWTaMP4c9LtGBayIu54X+DaHqIMllNEzunBeOq",
	"hBwMSlfTT6OJdjsnZjk5Vwt0YJDw4OHh2JFsppLj18/xz9xE04eP4miJZayJptGD0YMDLKczugIzV3q+",
	"5gwU1gjulyiOVGHzwvpWAt5bkMbbZZw7yPEczP2qNkv+kIqpRlwYjqYaBcFHfoufsrW9q7bjnlxb5bWb",
	"zvCqXhiNMzDQa5hrX5ovlzZ69euDIQRzmLJVXA3FRmNkeIt8UmWOCuox9uvMcpM8EZLOEPgjU29diTHQ",
	"K1bviJsfSIvZnoZKGYPWFyd+/OMkI++Kg4OHfyZCrYyfL7jxMrefm1Bvu0FZ1zua+aPNwxsJI8FlmPZI",
	"xhM8cJNSpJhUXXvq2mssQ9x0EBnCg8fkzRr0BsHCkFzDmqvCiK2XpTy0KnmGKkShBkahL9WKIFONoVrA",
	"6fb+ONpwIbBaKrt/lMLpxnEAVAuOPcdUFkKEOfLs1vPnobrH28indk09313E/4jEDIkGO/zqIxNsJ77C",
	"Sa3cOHhEI7ee7tTHjK/km9tq4o4lapQAN07ddRUDT1IqV9AXwcffvA77WxVGXa13id2IKbaLqzvgZQ8H",
	"bUw1lmrrA5Bu6IWrtowAwI4MXPUTRyYtLFMbN4MEE1ar5RLjegApvdO7+mmGXHvxNo6BOS0whff6UQMa",
	"DY3giYDkF5OjpzHJqTEbpVn5ynu5/ypCqC2X6kb4Ijo5fblxKTU8qeEmtTaPLpFHLpfKjw2kpYmt5xdR",
	"WV+RE6AYRIUWYaeZTibLsvriatJvE7/1w+BnVGck8/Mg8vj4COtSnoA00Djn+fHL9WGP/mazGa9kgcXY",
	"JOwxk1UuRofjgzHIcWoz379xK1rchuOiOFqDDuXKg/HB+ABXqxwkzTlWbu4nzHQ2dZaZ0Jy7Qsr5pDJO",
	"FeiZTplHzA+EM279pCB4+leKbUv1gXR7aJ4LTDpcycm58TDq/XafV7fHIpc9rbphpQpVcNR0eiwOXRSY",
	"XKGm8KSHBweflLMNNcQUSQJmWQixJf5TETDCZcjEa84KKvzXpXHn09qdsOkblQH+3AtS9iEuNosso3pb",
	"WZVQImHjBpuYoit3CtPMxvjPZX2K1aGbN5rotEXuRfl5xKDzEZAsV1xaJ2/lY5MqJ6xgwNGeg61msPdo",
	"1f7Ad0B11aJ66NtR4HOwRPQGw25mmgLXnbn5Naqrj6rUf15/L27p78O5Wsw5u9ypwmdgk9SHan2+G0xy",
	"lCp8NgkQ5In1Iipu6HFf8356j3a6JugcfLfN4SR3Lwhd+O+WznY38Fu/SbIAohlyXqrdZ5iJDrOf0aYe",
	"/QyCZTkkCiOi+0HMgaJ1QFF1+1Fy/0nBszcuG2BRonsJUvLwScGxkPA+h8QCIxDWNB2jZD8g5Ka0Z+lL",
	"4YfTgU3eJIgL9U7T9SjDV3Kklstr8i4W4ctlHwof9Wuo354iQxHosKdV/n1/iqhR6+wV1RfNuo9iQ+zL",
	"yz3afkJFmKx7D3MXYgT40C8z2IV0H/Zh+7kGslL+So4jPx42idxjEXmvQR2O2B3O1YDoU8Zyv6/6XQTz",
	"jX3wcWFTkNYPTcJoBr2hvASyqb553rFDaqBsi6uQnv/m3hoX8drgfXe1YRo1WAk0TBb9rz3DcUoS957U",
	"zfJlvAvMyO4dv22Xur17JCkkF2RT3kRKQYO/LbTdoYRhPxgljdHCIHgNjCHuFciaBw2o93WVGr2cN8Cz",
	"/6+8F/A82M0rYUxOUm5I4q4qLtwNI5ogYAhgvjD1w+KAJfXwuuUrMVEakavUSokvoEdCJVQ4aKPC3DWe",
	"raElTWF6rmrDXesd6TVJgRUCTvyM+P4awObN7wHDujvfzc53F1C9VuHOaPsenBt0l9dkLuPo0cHh3Y0k",
	"Wh8nB5g/Bl024U9Bcg+ajx4eflrEL52bSqksUQtLuXTVsNNXTBaF9bfpVspd7JXKwZ8PglsG0htPnVb0",
	"G7bb5+HO1Cb4nR4YWzQcd/LBTZ9D+z3swo3PRDfpwAPBj2/B7x7FG5LsCpFQpmD7jCz6uxS/AsRPUihp",
	"bRziJZCXie7RwV+GN9jyf8wIwdx0I2+0mJhwJb2mjd7Yov97yRZv6y+IKHlM7DbnCRVi2/rgl2u10mBM",
	"HG4khYvdmiwpF4WGvZBfAr0ByVrTFFR3SR3BBQuVMlL1uvRxP82eRI1aqGu8rxAo8V+Hle6KM4qwAuvG",
	"XtUFqwUVC0Fb0yrjbs11BnXHR+3RZbO2UllWyPAplNu0N98c1+SDNi5PL/8bAAD//4q9w500NAAA",
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
