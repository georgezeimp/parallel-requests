package output

import (
	"strings"
)

type Presenter struct{}

func NewPresenter() *Presenter {
	return &Presenter{}
}

func (p *Presenter) Prepare(line []string) string {
	return strings.Join(line, " ")
}
