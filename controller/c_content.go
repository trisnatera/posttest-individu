package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trisnatera/posttest-individu/model"
)

type response struct {
	Status  int
	Message string
	Data    []model.Content
}

type responseObject struct {
	Status  int
	Message string
	Data    model.Content
}

func GetAllDataContent(w http.ResponseWriter, r *http.Request) {
	var result response

	result.Status = 200
	result.Message = "success"
	result.Data = model.SelectContentAll()
	resJson, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}

func GetOneDataContent(w http.ResponseWriter, r *http.Request) {
	var result responseObject
	contentID := mux.Vars(r)["id"]

	result.Status = 200
	result.Message = "success"
	result.Data = model.SelectContentById(contentID)
	resJson, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}

func AddDataContent(w http.ResponseWriter, r *http.Request) {

	var result responseObject
	var request model.Content
	reqBody, err := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &request)
	if err != nil {
		result.Status = 401
		result.Message = "Kindly enter data with the event title and description only in order to add"

	} else {
		status := model.AddContent(request.ID, request.Title, request.Description)
		if status {
			result.Status = 201
			result.Message = "Success save data"

		} else {
			result.Status = 406
			result.Message = "Can't save data!"
		}
	}
	result.Data = request
	resJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}

func UpdateDataContent(w http.ResponseWriter, r *http.Request) {
	contentID := mux.Vars(r)["id"]
	var result responseObject
	var request model.Content
	reqBody, err := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &request)
	if err != nil {
		result.Status = 401
		result.Message = "Kindly enter data with the event title and description only in order to update"

	} else {
		status := model.UpdateContent(contentID, request.Title, request.Description)
		if status {
			result.Status = 201
			result.Message = "Success update data"

		} else {
			result.Status = 406
			result.Message = "Can't update data!"
		}
	}
	result.Data = request
	resJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}

func DeleteDataContent(w http.ResponseWriter, r *http.Request) {
	contentID := mux.Vars(r)["id"]
	var result responseObject
	status := model.DeleteContent(contentID)
	if status {
		result.Status = 201
		result.Message = "Success Delete data"

	} else {
		result.Status = 406
		result.Message = "Can't Delete data!"
	}
	resJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}
