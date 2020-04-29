package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/remotes/request"
	"github.com/joaopandolfi/blackwhale/utils"
	"../models"
	"../dao"
)

type LambdaController struct {
}

func (cc LambdaController) Save(w http.ResponseWriter, r *http.Request){
	var received map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&received)
	if err != nil {
		utils.Error("[LambdaController][Save] - Erron on get body", err.Error())
		handlers.RESTResponseError(w, "Invalid body "+err.Error())
		return
	}

	userIDs := handlers.GetHeader(r,"id")
	userID,_ := strconv.Atoi(userIDs)


	dao := dao.Lambda{}
	id := dao.GenerateID()
	received["id"] = id;
	err = dao.Save(models.Lambda{ UserID:userID, Generic:received})

	if err != nil{
		handlers.RESTResponseError(w,false)
	}else{
		handlers.RESTResponse(w,id)
	}
}

func (cc LambdaController) SaveWithTag(w http.ResponseWriter, r *http.Request){
	var received map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&received)
	if err != nil {
		utils.Error("[LambdaController][Save] - Erron on get body", err.Error())
		handlers.RESTResponseError(w, "Invalid body "+err.Error())
		return
	}

	userID := 0
	userIDs := handlers.GetHeader(r,"id")
	if userIDs != "" {
		userID,_ = strconv.Atoi(userIDs)
	}


	dao := dao.Lambda{}
	id := dao.GenerateID()
	received["id"] = id;
	err = dao.Save(models.Lambda{UserID:userID, Generic:received, Tag:received["tag"].(string)})

	if err != nil{
		handlers.RESTResponseError(w,false)
	}else{
		handlers.RESTResponse(w,id)
	}
}

func (cc LambdaController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars:=handlers.GetVars(r)
	dao:= dao.Lambda{}
	result,err := dao.GetById(vars["id"])
	if err != nil{
		utils.Error("[LambdaController][GetById] - Erron on save", err.Error())
		handlers.RESTResponseError(w,"Internal error")
		return
	}
	handlers.Response(w,result)
}

func (cc LambdaController) Forward(w http.ResponseWriter, r *http.Request) {
	var received map[string]string
	form, err := handlers.GetForm(r)
	err = handlers.DecodeForm(&received, form)
	if err != nil {
		utils.Error("[LambdaController][RESTNewPredict] - Erron on get body", err.Error())
		handlers.RESTResponseError(w, "Invalid body "+err.Error())
		return
	}

	body, _ := json.Marshal(received)

	var header map[string]string
	header = make(map[string]string)
	header["Content-Type"] = "application/json"
	result, err := request.PostWithHeader("", header, body)

	var res []string
	err = json.Unmarshal(result, &res)

	utils.Debug("[LambdaController][RestNewPredict] - JSON", res, string(result))
	handlers.RESTResponse(w, res)
}
