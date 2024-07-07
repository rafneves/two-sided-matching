package entities

type Man struct {
	ID         string
	Preference []*Woman
}

func (m *Man) Prefer(reference *Woman, over *Woman) bool {
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
