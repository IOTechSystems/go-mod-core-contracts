// Copyright (C) 2022 IOTech Ltd

package utils

import (
	"net/url"
	"path"
)

// EscapeAndJoinPath escape and join the path variables
func EscapeAndJoinPath(apiRoutePath string, pathVariables ...string) string {
	elements := make([]string, len(pathVariables)+1)
	elements[0] = apiRoutePath // we don't need to escape the route path like /device, /reading, ...,etc.
	for i, e := range pathVariables {
		elements[i+1] = url.QueryEscape(e)
	}
	return path.Join(elements...)
}

// edgeXClientReqURI returns the non-encoded path?query that would be used in an HTTP request for u.
func edgeXClientReqURI(u *url.URL) string {
	result := u.Scheme + "://" + u.Host + u.Path
	if u.ForceQuery || u.RawQuery != "" {
		result += "?" + u.RawQuery
	}
	return result
}
