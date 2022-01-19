package app

// NewTranslation returns an initialized TTranslation object.
func NewTranslation(langFrom, langTo, originalInput string) TTranslation {
	return TTranslation{
		LangFrom:   langFrom,
		LangTo:     langTo,
		Original:   originalInput,
		Translated: "",
	}
}
