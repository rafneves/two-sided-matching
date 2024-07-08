package stable_marriage

type engagement map[string]string

func newEngagement() engagement {
	return make(map[string]string)
}

func (e engagement) Breakup(manID string) {
	if manID == "" {
		return
	}
	delete(e, manID)
}

func (e engagement) Engage(manID string, womanID string) {
	if manID == "" || womanID == "" {
		return
	}
	e[manID] = womanID
}

func (e engagement) GetManID(womanID string) string {
	for mID, wID := range e {
		if wID == womanID {
			return mID
		}
	}

	return ""
}

func (e engagement) GetWomanID(manID string) string {
	return e[manID]
}

func (e engagement) Size() int {
	return len(e)
}
