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
	SetID(string)
	GetID() string
	SetRole(string)
	GetRole() string
	SetDescription(string)
	GetDescription() string
	SetCompany(string)
	GetCompany() string
	SetLocation(string)
	GetLocation() string
	SetRemote(bool)
	GetRemote() bool
	SetLink(string)
	GetLink() string
	SetSalary(uint64)
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

func (o *opening) SetRole(role string) {
	o.role = role
}

func (o *opening) GetRole() string {
	return o.role
}

func (o *opening) SetDescription(description string) {
	o.description = description
}

func (o *opening) GetDescription() string {
	return o.description
}

func (o *opening) SetCompany(company string) {
	o.company = company
}

func (o *opening) GetCompany() string {
	return o.company
}

func (o *opening) SetLocation(location string) {
	o.location = location
}

func (o *opening) GetLocation() string {
	return o.location
}

func (o *opening) SetRemote(remote bool) {
	o.remote = remote
}

func (o *opening) GetRemote() bool {
	return o.remote
}

func (o *opening) SetLink(link string) {
	o.link = link
}

func (o *opening) GetLink() string {
	return o.link
}

func (o *opening) SetSalary(salary uint64) {
	o.salary = salary
}

func (o *opening) GetSalary() uint64 {
	return o.salary
}
