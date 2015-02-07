package opc

type Part struct {
	Name          string
	ContentType   string
	GrowthHint    string
	Relationships []*Relationship
	Content       interface{}
}
