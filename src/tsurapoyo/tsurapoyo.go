package tsurapoyo
import "github.com/fatih/color"

type Poyo interface {
    HontoNoKimochi() string
}

func Doubled(p Poyo) string {
    doubled := p.HontoNoKimochi() + p.HontoNoKimochi()
    color.Red(doubled)
    return doubled
}

type Tsurapoyo struct {
    Kimochi string
    Sugosa  string
}

func (t *Tsurapoyo) HontoNoKimochi() string {
    return t.Sugosa + t.Kimochi
}