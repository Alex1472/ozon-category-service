package film

type Category struct {
	ID   uint64
	Name string
}

type Categories []*Category

func (c Categories) FilterById(id uint64) *Category {
	for _, v := range c {
		if v.ID == id {
			return v
		}
	}
	return nil
}
