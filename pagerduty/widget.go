package pagerduty

import (
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" PagerDuty ", "pagerduty", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data, err := Fetch()

	widget.UpdateRefreshedAt()
	widget.View.SetTitle(widget.Name)

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		content = data
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */
