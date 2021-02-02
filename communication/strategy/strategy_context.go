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

func (d *Data) Send() {
	d.algo.Send()
}
