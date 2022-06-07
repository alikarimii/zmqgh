package application

// out port
// implement(method of struct) in output adapter (driven adapter)
// type outputAdapter struct{}
// func (adapter outputAdapter) SavingMessage func() error {}
type ForSavingMessage func(data []byte) error
