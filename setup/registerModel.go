package setup

import "TestErajaya/model"

func GetListModel() []interface{} {
	var listModel []interface{}

	listModel = append(listModel, &model.Product{})
	return listModel
}
