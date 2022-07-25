package jetrender

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/CloudyKit/jet/v6/loaders/multi"
)

// MailTemplateRenderer jet renderer
type MailTemplateRenderer struct {
	tpl *jet.Set
}

// NewRenderer new jet renderer
func NewRenderer(loaders ...jet.Loader) *MailTemplateRenderer {
	return &MailTemplateRenderer{
		tpl: jet.NewSet(multi.NewLoader(loaders...)),
	}
}

func (r *MailTemplateRenderer) renderTemplate(name string, data map[string]interface{}) (out string, err error) {
	var view *jet.Template
	view, err = r.tpl.GetTemplate(name)
	if err != nil {
		return
	}

	vars := make(jet.VarMap)
	for k := range data {
		vars.Set(k, data[k])
	}

	buf := bytes.NewBuffer(nil)
	err = view.Execute(buf, vars, nil)
	if err != nil {
		return
	}

	out = strings.TrimSpace(buf.String())
	return
}

func (r *MailTemplateRenderer) findSubject(name string) string {
	return fmt.Sprintf("%s/subject.jet", name)
}

func (r *MailTemplateRenderer) findHTML(name string) string {
	return fmt.Sprintf("%s/html.jet", name)
}

func (r *MailTemplateRenderer) findTXT(name string) string {
	return fmt.Sprintf("%s/text.jet", name)
}

// TextBody render txt message body
func (r *MailTemplateRenderer) TextBody(name string, data map[string]interface{}) (string, error) {
	return r.renderTemplate(r.findTXT(name), data)
}

// Subject render subject message
func (r *MailTemplateRenderer) Subject(name string, data map[string]interface{}) (string, error) {
	return r.renderTemplate(r.findSubject(name), data)
}

// HtmlBody render html message body
func (r *MailTemplateRenderer) HtmlBody(name string, data map[string]interface{}) (string, error) {
	return r.renderTemplate(r.findHTML(name), data)
}
