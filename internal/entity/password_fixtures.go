package entity

type PasswordMap map[string]Password

func (m PasswordMap) Get(name string) Password {
	if result, ok := m[name]; ok {
		return result
	}

	return Password{}
}

func (m PasswordMap) Pointer(name string) *Password {
	if result, ok := m[name]; ok {
		return &result
	}

	return &Password{}
}

var PasswordFixtures = PasswordMap{
	"alice":  NewPassword("uqxetse3cy5eo9z2", "Alice123!"),
	"bob":    NewPassword("uqxc08w3d0ej2283", "Bobbob123!"),
	"friend": NewPassword("uqxqg7i1kperxvu7", "!Friend321"),
}

// CreatePasswordFixtures inserts known entities into the database for testing.
func CreatePasswordFixtures() {
	for _, entity := range PasswordFixtures {
		Db().Create(&entity)
	}
}
