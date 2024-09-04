package main

type TVer interface {
	switchOFF()
	switchOn()
	GetStatus()
	GetModel()
}

type TV struct {
	status bool
	model  string
}

func (t *TV) switchOFF() {
	t.status = false
}

func (t *TV) switchON() {
	t.status = true
}

func (t *TV) GetModel() string {
	return t.model
}

func (t *TV) GetStatus() bool {
	return t.status
}

type Samsunger struct {
	TV
}

func (s *Samsunger) SamsungHub() string {
	return "SamsungHub is opened"
}

type LGer struct {
	TV
}

func (s *LGer) LGHub() string {
	return "LGHub is opened"
}
