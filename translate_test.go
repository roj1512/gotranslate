package gotranslate

import (
	"testing"
	"time"
)

func TestTranslate(t *testing.T) {
	tests := []struct{
		From, To, Text, Result string
	}{
		{"ru", "en", "привет", "Hello"},
		{"ru", "en", "", ""},
		{"auto", "ja", "!!!", "!!!"},
		{"auto", "es", "👌👍😂", "👌👍😂"}, // emojis
		{"ur", "en", "ディック", ""}, // invalid languages
		{"sidahd!212", "pis--p1ko2", "", ""}, // invalid languages
	}
	translator := NewTranslator()
	for _, test := range tests {
		tr, err := translator.Translate(test.From, test.To, test.Text)
		if err != nil {
			t.Error(err)
			continue
		}
		if tr.Text != test.Result {
			t.Error("wanted:", test.Result, "got:", tr.Text)
		}
		time.Sleep(1 * time.Second)
	}
}


// Write TestTTS soon
