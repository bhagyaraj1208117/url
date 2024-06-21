package URLName

type URLName struct {
	Name          string
	HashedUrlName string
	Rounds        int
}

func NewUrl(name string) *URLName {
	return &URLName{Name: name}

}

func (domainName URLName) Hash() string {
	return "hash"
}
