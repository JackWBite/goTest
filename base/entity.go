package base

type Human struct {
	Name string
	Sex  int
}

type Student struct {
	Human

	Sid    int
	Age    int
	Class  string
	School string
}
