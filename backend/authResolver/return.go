package authResolver

type ReturnFactory func() (any, error)

func Return(a any, b error) ReturnFactory {
	return func() (any, error) {
		return a, b
	}
}
