package entity

type opening struct {
	id          string
	role        string
	description string
	company     string
	location    string
	remote      bool
	link        string
	salary      uint64
}

type Opening interface {
	SetID(id string)
	GetID() string
	GetRole() string
	GetDescription() string
	GetCompany() string
	GetLocation() string
	GetRemote() bool
	GetLink() string
	GetSalary() uint64
}

func NewOpening(role, description, company, location, link string, remote bool, salary uint64) Opening {
	id := ""

	return &opening{
		id,
		role,
		description,
		company,
		location,
		remote,
		link,
		salary,
	}
}

func (o *opening) SetID(id string) {
	o.id = id
}

func (o *opening) GetID() string {
	return o.id
}

func (o *opening) GetRole() string {
	return o.role
}

func (o *opening) GetDescription() string {
	return o.description
}

func (o *opening) GetCompany() string {
	return o.company
}

func (o *opening) GetLocation() string {
	return o.location
}

func (o *opening) GetRemote() bool {
	return o.remote
}

func (o *opening) GetLink() string {
	return o.link
}

func (o *opening) GetSalary() uint64 {
	return o.salary
}
