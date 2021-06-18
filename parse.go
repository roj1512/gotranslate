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
	"strings"
)

func (t Translator) parse(raw map[string]interface{}) TranslationResult {
	x := TranslationResult{}

	for _, sentence := range raw["sentences"].([]interface{}) {
		sentence := sentence.(map[string]interface{})

		if trans, ok := sentence["trans"].(string); ok {
			x.TextRaw = append(x.TextRaw, trans)
		}
		if orig, ok := sentence["orig"].(string); ok {
			x.OrigRaw = append(x.OrigRaw, orig)
		}
	}

	x.Orig = strings.Join(x.OrigRaw, " ")
	x.Text = strings.Join(x.TextRaw, " ")
	x.Lang = raw["src"].(string)
	return x
}
