package diffrules

func BuiltinRegistry() (*Registry, error) {
	return NewRegistry(NewDIF001())
}
