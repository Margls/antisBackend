package handlers

import (
	"antis/backend/models"
	"antis/backend/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)


type ItemHandler struct {
	service *service.ItemService
}

func NewItemHandler (service *service.ItemService) *ItemHandler{
	return &ItemHandler{service: service}
}

var validate = validator.New()

func (h *ItemHandler) CreateItem (w http.ResponseWriter, r *http.Request){

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
 		http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
	}

	 if err := validate.Struct(item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


	createdItem, err := h.service.CreateItem(r.Context(), &item)

	if err != nil {

		  http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
	}

	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdItem)
}

func (h *ItemHandler) GetItemById (w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idstr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
	}
	 item, err := h.service.GetItemById(r.Context(), id)

	 if err != nil {
		http.Error(w,"Not found", 404)
		return

	 }

	  w.Header().Set("Content-Type", "application/json")
	  json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) GetAllItems (w http.ResponseWriter, r *http.Request) {

	items := h.service.GetAllItems(r.Context())

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(items)
}