package models

func GetAllOrder() ([]OrderStorage, error) {
	o := []OrderStorage{}
	err := db.Select(&o, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	return o, nil
}

func EditOrder(data Order) error {
	_, err := db.NamedExec("UPDATE orders SET status=:status WHERE id=:id", data)
	return err
}
