package translation

import (
	gt "github.com/bas24/googletranslatefree"
)

// translateText translates the text to the target language using external package
func translateText(text string, targetLang string) (string, error) {
	translatedText, err := gt.Translate(text, "auto", targetLang)
	if err != nil {
		return "", err
	}

	return translatedText, nil
}
