package sqlite3

import (
	"github.com/okmall/okmall/models/model"
	"net/http"
)

// GetMainPage get main page data
func (r *sqlite3Repository) GetMainPage() (interface{}, error) {
	data := &model.MainPage{Url: "/index/index"}
	return &model.Result{ErrorNo: http.StatusOK, Data: data}, nil
}
