package strategy

type Data struct {
	algo Algorithm
}

func NewStrategy(algo Algorithm) *Data {
	return &Data{algo: algo}
}

func (d *Data) SetAlgorithm(algo Algorithm) {
	d.algo = algo
}

func (d *Data) Before(handler BeforeHandler) {
	d.algo.Before(handler)
}

func (d *Data) After(handler AfterHandler) {
	d.algo.After(handler)
}

func (d *Data) Send() {
	d.algo.Send()
}
