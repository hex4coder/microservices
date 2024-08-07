package main

import (
	"fmt"
	"net/http"

	"github.com/hex4coder/ams/common"
	pb "github.com/hex4coder/ams/common/api"
)

type handler struct {
	client pb.UserServiceClient
}

func NewHandler(client pb.UserServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("POST %s", common.BASE_URL), h.HandleCreateUser)
	mux.HandleFunc(fmt.Sprintf("GET %s", common.BASE_URL), h.HandleGetUser)
}

func (h *handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user")

	var userReq pb.CreateUserRequest

	if err := common.ReadJSON(r, &userReq); err != nil {
		common.WriteError(w, http.StatusBadRequest, fmt.Sprintf("%s and user request not valid", err.Error()))
		return
	}

	user, err := h.client.CreateUser(r.Context(), &userReq)

	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
	}

	common.WriteJSON(w, http.StatusCreated, &user)
}

func (h *handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	var getReq pb.GetUserRequest
	if err := common.ReadJSON(r, &getReq); err != nil {
		common.WriteError(w, http.StatusBadRequest, fmt.Sprintf("failed decode get user req : %s", err.Error()))
		return
	}

	getRes, err := h.client.GetUser(r.Context(), &getReq)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, fmt.Sprintf("get response failed GetUser Methods : %s", err.Error()))
		return
	}

	common.WriteJSON(w, http.StatusOK, &getRes)
}
