package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/external/mysql"
)

func ReadAllUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := mysql.GetAllUser()
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, user, nil)
	}
}

func UpdateEnrollmentStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user mysql.User
		args := mux.Vars(r)
		payloads, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		json.Unmarshal(payloads, &user)

		i, _ := strconv.ParseUint(args["id"], 10, 64)

		err = mysql.UpdateUser(uint(i), user)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		userUpdated, err := mysql.GetUserById(uint(i))
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, userUpdated, nil)
	}
}
