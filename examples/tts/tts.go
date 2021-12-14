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

package main

import (
	"fmt"
	"os"

	"github.com/roj1512/gotranslate"
)

func main() {
	translator := gotranslate.NewTranslator()
	file, err := os.Create("test.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	options := gotranslate.NewTTSOptions()
	options.TargetLang = "en"
	written, err := translator.TTS(
		"Hello world!",
		file,
		options,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(written)
}
