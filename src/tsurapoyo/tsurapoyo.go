package tsurapoyo

type Tsurapoyo struct {
    Kimochi string
    Sugosa  string
}

func (t *Tsurapoyo) HontoNoKimochi() string {
    return t.Sugosa + t.Kimochi
}