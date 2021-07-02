package gotranslate

import (
	"fmt"
	"net/http"
)

type Translator struct {
	Client  *http.Client
	Url     string
	TTSUrl  string
	Headers map[string]string
}

type TranslationResult struct {
	Orig    string
	Text    string
	OrigRaw []string
	TextRaw []string
	Lang    string
}

type TranslateOptions struct {
	Q          string `json:"q"` // Text that we need to translate
	SourceLang string `json:"sl"`
	TargetLang string `json:"tl"`
	Client     string `json:"client"` // I don't know what is it...
	Dt         string `json:"dt"` // I don't know what is it...
	Dj         string `json:"dj"` // I don't know what is it...
	InputEncoding         string `json:"ie"`
	OutputEncoding         string `json:"oe"`
}


type TTSOptions struct {
	Q          string `json:"q"` // Text that we need to translate
	TextLen    string `json:"textlen"`
	TargetLang string `json:"tl"`
	Client     string `json:"client"` // I don't know what is it...
	IDx        string `json:"idx"` // I don't know what is it...
	Prev       string `json:"prev"`
	InputEncoding         string `json:"ie"`
	OutputEncoding         string `json:"oe"`
}

type HTTPError struct {
	Code int
	Description string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("Got HTTP error [%d]: %s", e.Code, e.Description)
}