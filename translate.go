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
	"io"
	"strconv"
	"strings"
)

func (t *Translator) Translate(from, to, text string) (*TranslationResult, error) {
	options := NewTranslateOptions()
	options.SourceLang = from
	options.TargetLang = to
	options.Q = text
	res, err := t.doRequest("POST", t.Url, options)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&result)
	out := t.parse(result)
	return &out, err
}

// TranslateCustom function translate your text with your options
func (t *Translator) TranslateWithOptions(text string, options *TranslateOptions) (*TranslationResult, error) {
	if options == nil {
		options = NewTranslateOptions()
	}
	options.Q = text

	res, err := t.doRequest("POST", t.Url, options)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&result)
	out := t.parse(result)
	return &out, err
}



// Detect is wrapper of Translate but return only language of text
func (t *Translator) Detect(text string) (string, error) {
	result, err := t.TranslateWithOptions(text, nil)
	if err != nil {
		return "", err
	}
	
	return result.Lang, err
}




// TTS speech your text into a file
func (t Translator) TTS(text string, dst io.Writer, options *TTSOptions) (int64, error) {
	if options == nil {
		options = NewTTSOptions()
	}
	options.Q = text
	options.TextLen = strconv.Itoa(len(text))
	
	res, err := t.doRequest("GET", t.TTSUrl, options)
	if err != nil {
		return 0, err
	}
	
	return io.Copy(dst, res.Body)
}



// parse answer from the Google API
func (t Translator) parse(raw map[string]interface{}) TranslationResult {
	x := TranslationResult{}
	if _, ok := raw["sentences"]; ok {
		for _, sentence := range raw["sentences"].([]interface{}) {
			sentence := sentence.(map[string]interface{})
			
			if trans, ok := sentence["trans"].(string); ok {
				x.TextRaw = append(x.TextRaw, trans)
			}
			if orig, ok := sentence["orig"].(string); ok {
				x.OrigRaw = append(x.OrigRaw, orig)
			}
		}
	}

	
	x.Orig = strings.Join(x.OrigRaw, " ")
	x.Text = strings.Join(x.TextRaw, " ")
	if _, ok := raw["src"]; ok {
		x.Lang = raw["src"].(string)
	}
	return x
}