package entity

type Tent struct {
	ID   int
	From *User
	To   *User
}

func (t *Tent) Target(id ID) (ID, error) {
	if t.From.ID == id {
		return t.To.ID, nil
	} else if t.To.ID == id {
		return t.From.ID, nil
	} else {
		return 0, ExternalError{"No such private channel."}
	}
}
