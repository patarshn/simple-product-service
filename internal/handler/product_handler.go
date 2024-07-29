package handler

import (
	"encoding/json"
	"net/http"
	"product-service/internal/models"
	"product-service/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var (
		product models.Product
	)

	reqID := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(reqID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ResponseBody{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	product, err = h.service.GetProductByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ResponseBody{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ResponseBody{
		Message: "Success Get Product",
		Data:    product,
	})
}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var (
		req models.GetAllProductRequest
	)

	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}

	req.Limit = limit
	req.Page = page

	products, err := h.service.GetAllProduct(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ResponseBody{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ResponseBody{
		Message: "Success get all product",
		Data:    products,
	})
	return
}
