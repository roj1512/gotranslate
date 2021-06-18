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

type TranslationResult struct {
	Orig    string
	Text    string
	OrigRaw []string
	TextRaw []string
	Lang    string
}

type TranslateOptions struct {
	Q          string `json:"q"`
	SourceLang string `json:"sl"`
	TargetLang string `json:"tl"`
	Client     string `json:"client"`
	Dt         string `json:"dt"`
	Dj         string `json:"dj"`
	Ie         string `json:"ie"`
	Oe         string `json:"oe"`
}

func NewTranslateOptions() *TranslateOptions {
	return &TranslateOptions{
		SourceLang: "auto",
		TargetLang: "en",
		Client:     "gtx",
		Dt:         "t",
		Dj:         "1",
		Ie:         "utf-8",
		Oe:         "utf-8",
	}
}

func (t Translator) Translate(text string, options *TranslateOptions) (TranslationResult, error) {
	if options == nil {
		options = NewTranslateOptions()
	}
	options.Q = text

	result := map[string]interface{}{}
	res, err := t.doRequest("POST", t.Url, options)
	if err != nil {
		return t.parse(result), err
	}

	err = decodeJSON(res, &result)
	return t.parse(result), err
}
