package models

func GetAllAdmin() ([]Admin, error) {
	a := []Admin{}
	err := db.Select(&a, "SELECT * FROM admin")
	if err != nil {
		return nil, err
	}
	return a, nil
}

func GetOneAdmin(id int) ([]Admin, error) {
	a := []Admin{}
	err := db.Select(&a, "SELECT id, name, username FROM admin WHERE id = ?", id)
	if err != nil {
		return a, err
	}
	return a, nil
}

func AddAdmin(data Admin) error {
	_, err := db.NamedExec("INSERT INTO admin(name, username, password, created) VALUES(:name, :username, :password, :created)", data)
	return err
}

func EditAdmin(data Admin) error {
	var err error
	if data.PassWord != "" {
		_, err = db.NamedExec("UPDATE admin SET name=:name, password=:password, updated=:updated WHERE id=:id", data)
	} else {
		_, err = db.NamedExec("UPDATE admin SET name=:name, updated=:updated WHERE id =:id", data)
	}
	return err
}

func DeleteAdmin(data Admin) error {
	_, err := db.NamedExec("DELETE FROM admin WHERE id =:id", data)
	return err
}
