package utils

import (
	"bytes"
	"github.com/pkg/errors"
	"html/template"
)

func ParseTemplate(tmplFiles string, data any) (string, error) {
	tmpl, err := template.ParseFiles(tmplFiles)
	if err != nil {
		return "", errors.WithMessage(err, "failed to parse template")
	}

	var emailBuff bytes.Buffer
	err = tmpl.Execute(&emailBuff, data)
	if err != nil {
		return "", errors.WithMessage(err, "failed to execute template")
	}

	return emailBuff.String(), nil
}
