// Copyright (c) 2022 Institute of Software, Chinese Academy of Sciences (ISCAS)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"net/url"
	"strings"
)

func JoinUrl(base string, path string) (string, error) {
	if !strings.HasSuffix(base, "/") {
		base = base + "/"
	}
	b, err := url.Parse(base)

	if err != nil {
		return "", err
	}
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	p, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	fullUrl := b.ResolveReference(p).String()
	return fullUrl, nil
}

func ParseUrl(uri string) (host string, path string, err error) {
	if u, err := url.ParseRequestURI(uri); err == nil {
		host = strings.ToLower(strings.TrimSpace(u.Host))
		path = u.Path
		return host, path, err
	} else {
		return "", "", err
	}

}
