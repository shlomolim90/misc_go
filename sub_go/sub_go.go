package sub_go

import "fmt"

var idCount int;

type SubGo struct {
	ID	int
	Description string
}

func New(description string) (*SubGo, error) {
	idCount++;

	return &SubGo{
		ID: idCount,
		Description: description,
	}, nil
}

func (mg *SubGo) GetInfo() (string) {
	return fmt.Sprintf("Subgo: ID: %d, Description: %s\n", mg.ID, mg.Description);
}
