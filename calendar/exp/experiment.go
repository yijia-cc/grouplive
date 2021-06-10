package exp

type Experiment interface {
	Save()
	Start()
	End()
}
