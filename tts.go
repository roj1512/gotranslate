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
	"io"
	"strconv"
)

type TTSOptions struct {
	Q          string `json:"q"`
	TextLen    string `json:"textlen"`
	TargetLang string `json:"tl"`
	Client     string `json:"client"`
	IDx        string `json:"idx"`
	Prev       string `json:"prev"`
	Ie         string `json:"ie"`
	Oe         string `json:"oe"`
}

func NewTTSOptions() *TTSOptions {
	return &TTSOptions{
		TargetLang: "en",
		Client:     "at",
		IDx:        "0",
		Prev:       "input",
		Ie:         "utf-8",
		Oe:         "utf-8",
	}
}

func (t Translator) TTS(text string, dst io.Writer, options *TTSOptions) (int64, error) {
	if options == nil {
		options = NewTTSOptions()
	}
	options.Q = text
	options.TextLen = string(strconv.AppendInt([]byte{}, int64(len(text)), 10))

	res, err := t.doRequest("GET", t.TTSUrl, options)
	if err != nil {
		return 0, err
	}

	return io.Copy(dst, res.Body)
}
