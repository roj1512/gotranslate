/*
 * gotranslate - A Go package for translating text using Google Translate
 * Copyright (C) 2021  Roj Serbest
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package gotranslate

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func getQuery(options interface{}) (string, error) {
	bytes, err := json.Marshal(options)
	if err != nil {
		return "", err
	}

	values, toAdd := url.Values{}, map[string]string{}
	err = json.Unmarshal(bytes, &toAdd)
	if err != nil {
		return "", err
	}

	for k, v := range toAdd {
		values.Add(k, v)
	}

	return values.Encode(), nil
}

func (t Translator) doRequest(method string, url string, options interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	query, err := getQuery(options)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query
	for k, v := range t.Headers {
		req.Header.Set(k, v)
	}

	res, err := t.Client.Do(req)
	if err != nil {
		return nil, err
	}
	
	if res.StatusCode != 200 {
		var ans string
		json.NewDecoder(res.Body).Decode(&ans)
		if ans == "" {
			ans = "empty page"
		}
		return nil, HTTPError{
			Code:        res.StatusCode,
			Description: ans,
		}
	}
	
	return res, nil
}
