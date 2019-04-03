package core

// MinaAction indicator mina service interface
type MinaAction interface {
	// GetMainPage fetch main page data
	GetMainPage() (interface{}, error)
}
