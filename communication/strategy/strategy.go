package strategy

type BeforeHandler func(algorithm Algorithm)
type AfterHandler func(algorithm Algorithm, response interface{})

type Algorithm interface {
	Send()
	Before(handler BeforeHandler)
	After(handler AfterHandler)
}
