package unsafe_local_test

type Programmer struct {
	name     string
	age      int
	language string
}

func GetLocalValue() Programmer {
	p := Programmer{
		name:     "yang",
		age:      18,
		language: "chinese",
	}
	return p
}
