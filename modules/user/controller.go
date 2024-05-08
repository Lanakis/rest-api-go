package user

import (
	"authorization/api"
	"authorization/modules/user/dto"
	"authorization/modules/user/entity"
	"authorization/utils"
	"authorization/utils/filter"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Handler(service entity.IUserService) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Create(w, r, service)
		case http.MethodGet:
			parts := strings.Split(r.URL.Path, "/")
			fmt.Print("parts ", strings.Split(r.URL.Path, "/"))

			if len(parts) > 2 {
				id := parts[len(parts)-1]
				_, err := strconv.Atoi(id)
				if err != nil {
					FindAll(w, r, service)
				} else {
					FindOne(w, r, service, id)
				}
			}
		case http.MethodPatch:
			parts := strings.Split(r.URL.Path, "/")
			id := parts[len(parts)-1]
			Update(w, r, service, id)
		case http.MethodDelete:
			parts := strings.Split(r.URL.Path, "/")
			id := parts[len(parts)-1]
			Delete(w, r, service, id)
		}
	}
}

func Create(writer http.ResponseWriter, requests *http.Request, service entity.IUserService) {
	var userCreateRequest dto.Create
	err := json.NewDecoder(requests.Body).Decode(&userCreateRequest)
	if err != nil {
		http.Error(writer, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// Вызываем метод Create сервиса для создания пользователя
	service.Create(requests.Context(), userCreateRequest)

	webResponse := utils.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   userCreateRequest,
	}
	api.ResponseJSON(writer, webResponse)
}

func FindAll(writer http.ResponseWriter, requests *http.Request, service entity.IUserService) error {
	var filterOptions filter.Option
	if options, ok := requests.Context().Value(filter.OptionsContextKey).(filter.Option); ok {
		filterOptions = options
	}
	username := requests.URL.Query().Get("username")
	if username != "" {
		err := filterOptions.AddField("username", filter.OperatorLike, username, filter.DataTypeStr)
		if err != nil {
			return err
		}
	}
	role := requests.URL.Query().Get("role")
	if role != "" {
		err := filterOptions.AddField("role", filter.OperatorLike, role, filter.DataTypeStr)
		if err != nil {
			return err
		}
	}
	firstName := requests.URL.Query().Get("first_name")
	if firstName != "" {
		if err := filterOptions.AddField("first_name", filter.OperatorLike, firstName, filter.DataTypeStr); err != nil {
			return err
		}

	}

	//fmt.Printf("%v\n Смотрим что в фильтре ", b)
	middleName := requests.URL.Query().Get("middle_name")
	if middleName != "" {
		err := filterOptions.AddField("middle_name", filter.OperatorLike, middleName, filter.DataTypeStr)
		if err != nil {
			return err
		}
	}

	lastName := requests.URL.Query().Get("last_name")
	if lastName != "" {
		err := filterOptions.AddField("last_name", filter.OperatorLike, lastName, filter.DataTypeStr)
		if err != nil {
			return err
		}
	}

	age := requests.URL.Query().Get("age")
	if age != "" {
		operator := filter.OperatorEq
		value := age
		// Поиск конструкции age=lt:Число или age=Число
		if strings.Index(age, ":") != -1 {
			splitValue := strings.Split(age, ":")
			operator = splitValue[0]
			value = splitValue[1]
		}
		err := filterOptions.AddField("age", operator, value, filter.DataTypeInt)
		if err != nil {
			return err
		}
	}

	createAt := requests.URL.Query().Get("created_at")
	if createAt != "" {
		var operator string
		if strings.Index(createAt, ":") != -1 {
			// Тут диапазон
			operator = filter.OperatorBetween
		} else {
			//Одиночная дата
			operator = filter.OperatorEq
		}
		err := filterOptions.AddField("age", operator, createAt, filter.DataTypeDate)
		if err != nil {
			return err
		}

	}

	head := requests.URL.Query().Get("head")
	if head != "" {
		_, err := strconv.ParseBool(head)
		if err != nil {
			validationError := utils.BadRequestError("filter params validation failed", "bool value wrong parameter")
			validationError.WithParams(map[string]string{"head": "this field should be boolean"})
			return validationError
		}
		err = filterOptions.AddField("head", filter.OperatorEq, head, filter.DataTypeBool)
		if err != nil {
			return err
		}
	}

	usersResponse, _ := service.FindAll(requests.Context(), filterOptions)
	response := dto.Response{
		Users: usersResponse,
		Count: len(usersResponse),
	}
	api.ResponseJSON(writer, response)
	return nil
}

func FindOne(w http.ResponseWriter, r *http.Request, service entity.IUserService, id string) {

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	userID, _ := strconv.Atoi(id)
	el, err := service.FindOne(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to find entity", http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v \n", el)
	api.ResponseJSON(w, el)
}

func Update(writer http.ResponseWriter, requests *http.Request, service entity.IUserService, id string) {
	if id == "" {
		http.Error(writer, "ID is required", http.StatusBadRequest)
		return
	}
	userID, _ := strconv.Atoi(id)
	userUpdateRequest := dto.Update{}
	service.Update(requests.Context(), userUpdateRequest, userID)
	webResponse := utils.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   userUpdateRequest,
	}
	api.ResponseJSON(writer, webResponse)
}

func Delete(writer http.ResponseWriter, requests *http.Request, service entity.IUserService, id string) {
	if id == "" {
		http.Error(writer, "ID is required", http.StatusBadRequest)
		return
	}
	userID, _ := strconv.Atoi(id)
	service.Delete(requests.Context(), userID)
	response := utils.DeleteResponse{Code: 200, Status: "Ok"}
	api.ResponseJSON(writer, response)

}
