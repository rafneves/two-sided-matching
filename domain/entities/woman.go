package entities

type Woman struct {
	ID         string
	Preference []*Man
}

func (m *Woman) Prefer(reference *Man, over *Man) bool {
	if reference == nil {
		return false
	}

	if over == nil {
		return true
	}

	for _, person := range m.Preference {
		if person.ID == reference.ID {
			return true
		}
		if person.ID == over.ID {
			return false
		}
	}
	return false
}
