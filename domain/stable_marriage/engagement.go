package stable_marriage

type engagement map[string]string

func newEngagement() engagement {
	return make(map[string]string)
}

func (e engagement) Breakup(womanID string) {
	delete(e, womanID)
}

func (e engagement) Engage(manID string, womanID string) {
	e[womanID] = manID
}

func (e engagement) GetManID(womanID string) string {
	return e[womanID]
}

func (e engagement) GetWomanID(manID string) string {

	for wID, mID := range e {
		if mID == manID {
			return wID
		}
	}

	return ""
}
