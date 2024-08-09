package productcontroller

import (
	"encoding/json"
	"github/fdjrr/go-restapi-mux/helpers"
	"github/fdjrr/go-restapi-mux/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helpers.ResponseJson
var ResponseError = helpers.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, products)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	ResponseJson(w, http.StatusCreated, product)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found!")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ? ", id).Updates(&product).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Product not found!")
		return
	}

	product.Id = id

	ResponseJson(w, http.StatusOK, product)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if models.DB.Delete(&models.Product{}, id).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Product not found!")
		return
	}

	ResponseJson(w, http.StatusOK, map[string]string{"message": "Product deleted successfully!"})
}
