package gotranslate

import "net/http"

func NewTranslateOptions() *TranslateOptions {
	return &TranslateOptions{
		SourceLang: 		   "auto",
		TargetLang: 		   "en",
		Client:     		   "gtx",
		Dt:         		   "t",
		Dj:        			   "1",
		InputEncoding:         "utf-8",
		OutputEncoding:        "utf-8",
	}
}

func NewTTSOptions() *TTSOptions {
	return &TTSOptions{
		TargetLang: 		   "en",
		Client:     		   "at",
		IDx:        		   "0",
		Prev:       		   "input",
		InputEncoding:         "utf-8",
		OutputEncoding:        "utf-8",
	}
}

func NewTranslator() *Translator {
	return &Translator{
		Client:  &http.Client{},
		Url:     "https://translate.googleapis.com/translate_a/single",
		TTSUrl:  "https://translate.google.com/translate_tts",
		Headers: map[string]string{"User-Agent": "GoogleTranslate/6.6.1.RC09.302039986 (Linux; U; Android 9; Redmi Note 8)"},
	}
}