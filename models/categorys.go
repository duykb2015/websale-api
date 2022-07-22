package models

func GetAllCategory() ([]Category, error) {
	c := []Category{}

	err := db.Select(&c, "SELECT id, name, parent_id FROM category WHERE parent_id = 0")
	if err != nil {
		return nil, err
	}
	for key, val := range c {
		subc := []SubCategory{}
		err = db.Select(&subc, "SELECT id, name, parent_id FROM category WHERE parent_id = ?", val.ID)
		if err != nil {
			return nil, err
		}
		c[key].SubCategory = subc
	}

	return c, nil
}

func AddCategory(data Category) error {
	_, err := db.NamedExec("INSERT INTO category(name, parent_id) VALUES(:name, :parent_id)", data)
	return err
}

func EditCategory(data Category) {
	db.NamedExec("UPDATE category SET name=:name, parent_id=:parent_id WHERE id=:id", data)
}

func DeleteCategory(data Category) {
	db.NamedExec("DELETE FROM category WHERE id =:id", data)
}
