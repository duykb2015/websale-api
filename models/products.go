package models

func GetAllProduct() ([]Product, error) {
	p := []Product{}
	err := db.Select(&p, "SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetOneProduct(slug string) ([]Product, error) {
	a := []Product{}
	err := db.Select(&a, "SELECT * FROM product WHERE slug = ?", slug)
	if err != nil {
		return a, err
	}
	return a, nil
}

func AddProduct(data Product) error {
	_, err := db.NamedExec("INSERT INTO product(category_id, name, slug, price, discount, specifications, promotion, guarantee, total, product_image, image_list, status, created) VALUES(:category_id, :name, :slug, :price, :discount, :specifications, :promotion, :guarantee, :total, :product_image, :image_list, :status, :created)", data)
	return err
}

func EditProduct(data Product) error {
	var err error
	if data.ProductImage == "" {
		_, err = db.NamedExec("UPDATE product SET category_id=:category_id, name=:name, slug=:slug, price=:price, specifications=:specifications, discount=:discount, promotion=:promotion, guarantee=:guarantee, total=:total, status=:status, updated=:updated WHERE id=:id", data)
	} else {
		if data.ImageList == "" {
			_, err = db.NamedExec("UPDATE product SET category_id=:category_id, name=:name, slug=:slug, price=:price, specifications=:specifications, discount=:discount, product_image=:product_image, promotion=:promotion, guarantee=:guarantee, total=:total, status=:status, updated=:updated WHERE id=:id", data)
		} else {
			_, err = db.NamedExec("UPDATE product SET category_id=:category_id, name=:name, slug=:slug, price=:price, specifications=:specifications, discount=:discount, product_image=:product_image, image_list=:image_list, promotion=:promotion, guarantee=:guarantee, total=:total, status=:status, updated=:updated WHERE id=:id", data)
		}
	}
	return err
}

func DeleteProduct(data Product) error {
	_, err := db.NamedExec("DELETE FROM product WHERE id =:id", data)
	return err
}
