package coder

type (
	Tester interface {
		Test() (int, error)
	}

	superTester struct {
	}

	coder struct {
	}
)

func NewCoder() *coder {
	return &coder{}
}

func (s *superTester) Test() (int, error) {
	return 1, nil
}

func (c *coder) Code(tester Tester) (int, error) {
	return tester.Test()
}
