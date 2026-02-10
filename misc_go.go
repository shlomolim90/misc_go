package misc_go

import "fmt"

var idCount int;

type MiscGo struct {
	ID	int
	Description string
}

func New(description string) (*MiscGo, error) {
	idCount++;

	return &MiscGo{
		ID: idCount,
		Description: description,
	}, nil
}


func (mg *MiscGo) GetInfo() (string) {
	return fmt.Sprintf("ID: %d, Description: %s\n", mg.ID, mg.Description);
}
