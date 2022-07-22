package models

func GetAllTransaction() ([]Transaction, error) {
	t := []Transaction{}
	err := db.Select(&t, "SELECT * FROM transaction")
	if err != nil {
		return nil, err
	}
	return t, nil
}

func EditTransaction(data Transaction) error {
	_, err := db.NamedExec("UPDATE orders SET status=:transaction_status WHERE id=:id", data)
	return err
}
