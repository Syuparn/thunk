package definedtypeslash

type MyIF interface {
	WithTypes(i MyInt, s mySecret) MyStruct
}

type MyInt int

// unexported
type mySecret string

type MyStruct struct {
	a int
}
