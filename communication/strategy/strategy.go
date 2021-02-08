package strategy

type BeforeHandler func(algorithm Algorithm)
type AfterHandler func(response interface{})

type Algorithm interface {
	Send()
	Before(handler BeforeHandler)
	After(handler AfterHandler)
}
