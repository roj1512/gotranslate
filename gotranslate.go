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
	"net/http"
)

var baseHeaders = map[string]string{
	"User-Agent": "GoogleTranslate/6.6.1.RC09.302039986 (Linux; U; Android 9; Redmi Note 8)",
}

type Translator struct {
	Client  *http.Client
	Url     string
	TTSUrl  string
	Headers map[string]string
}

func NewTranslator() *Translator {
	return &Translator{
		Client:  &http.Client{},
		Url:     "https://translate.googleapis.com/translate_a/single",
		TTSUrl:  "https://translate.google.com/translate_tts",
		Headers: baseHeaders,
	}
}
