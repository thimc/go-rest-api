package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []User `json:"result,omitempty"`
}

func handleCreate(res http.ResponseWriter, req *http.Request) {
	// if req.Method != http.MethodPost {
	// 	writeResponse(res, http.StatusMethodNotAllowed, ApiResponse{
	// 		Succes:  false,
	// 		Message: "Invalid HTTP method",
	// 	})
	// 	return
	// }

	if _, ok := req.URL.Query()["id"]; !ok {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: "Invalid query, this method needs an id",
		})
		return
	}
	id, err := strconv.Atoi(req.URL.Query()["id"][0])
	if err != nil {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: fmt.Sprintf("Unknown ID: %d", id),
		})
		return
	}

	if _, ok := database[id]; ok {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: fmt.Sprintf("ID: %d is already in use", id),
		})
		return
	}

	requiredFields := []string{"name", "nick", "mail"}
	fields := make(map[string]string, len(requiredFields))

	for _, field := range requiredFields {
		if _, ok := req.URL.Query()[field]; ok {
			fields[field] = req.URL.Query()[field][0]
		}
	}

	if len(fields) < 1 || fields["name"] == "" {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: "Invalid query, you need to supply a id and at least a name, nickname or mail",
		})
		return
	}

	message := fmt.Sprintf("ID %d was created. User %s", id, fields["name"])
	if fields["nick"] != "" {
		message = fmt.Sprintf("%s, nick: %s", message, fields["nick"])
	}
	if fields["mail"] != "" {
		message = fmt.Sprintf("%s, mail: %s", message, fields["mail"])
	}

	database[id] = User{
		Name:     fields["name"],
		Nickname: fields["nick"],
		Mail:     fields["mail"],
	}

	log.Println(message)
	writeResponse(res, http.StatusOK, ApiResponse{
		Success: true,
		Message: message,
	})
}

func handleUpdate(res http.ResponseWriter, req *http.Request) {
	// if req.Method != http.MethodPost {
	// 	writeResponse(res, http.StatusMethodNotAllowed, ApiResponse{
	// 		Succes:  false,
	// 		Message: "Invalid HTTP method",
	// 	})
	// 	return
	// }

	if _, ok := req.URL.Query()["id"]; !ok {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: "Invalid query, this method needs an id",
		})
		return
	}
	id, err := strconv.Atoi(req.URL.Query()["id"][0])
	if err != nil {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: fmt.Sprintf("Unknown ID: %d", id),
		})
		return
	}

	validFields := []string{"name", "nick", "mail"}
	fields := make(map[string]string, len(validFields))

	for _, field := range validFields {
		if _, ok := req.URL.Query()[field]; ok {
			fields[field] = req.URL.Query()[field][0]
		}
	}

	if len(fields) < 1 {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: "Invalid query, requires either a new name, nickname or a mail",
		})
		return
	}

	message := ""

	for i, field := range fields {
		user := database[id]
		switch i {
		case "name":
			message = fmt.Sprintf("Updated the username, %s -> %s", user.Name, field)
			user.Name = field
		case "nick":
			message = fmt.Sprintf("Updated the nickname, %s -> %s", user.Nickname, field)
			user.Nickname = field
		case "mail":
			message = fmt.Sprintf("Updated the mail, %s -> %s", user.Mail, field)
			user.Mail = field
		}

		database[id] = user
	}

	log.Println(message)
	writeResponse(res, http.StatusOK, ApiResponse{
		Success: true,
		Message: message,
	})
}

func handleDelete(res http.ResponseWriter, req *http.Request) {
	// if req.Method != http.MethodDelete {
	// 	writeResponse(res, http.StatusMethodNotAllowed, ApiResponse{
	// 		Succes:  false,
	// 		Message: "Invalid HTTP method",
	// 	})
	// 	return
	// }

	if _, ok := req.URL.Query()["id"]; !ok {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: "Invalid query, this method needs an id",
		})
		return
	}

	id, err := strconv.Atoi(req.URL.Query()["id"][0])
	if err != nil {
		writeResponse(res, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: fmt.Sprintf("Unknown ID: %d", id),
		})
		return
	}

	user, ok := database[id]
	if !ok {
		writeResponse(res, http.StatusNotFound, ApiResponse{
			Success: false,
			Message: fmt.Sprintf("Unknown user with ID: %d", id),
		})
		return
	}

	delete(database, id)
	writeResponse(res, http.StatusOK, ApiResponse{
		Success: true,
		Message: fmt.Sprintf("User %s was removed from the database", user.Name),
	})
}

func handleRead(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		writeResponse(res, http.StatusMethodNotAllowed, ApiResponse{
			Success: false,
			Message: "Invalid HTTP method",
		})
	}

	if idStr, ok := req.URL.Query()["id"]; ok {
		id, err := strconv.Atoi(idStr[0])
		if err != nil {
			writeResponse(res, http.StatusBadRequest, ApiResponse{
				Success: false,
				Message: fmt.Sprintf("Invalid ID: %d", id),
			})
			return
		}

		user, ok := database[id]
		if !ok {
			writeResponse(res, http.StatusNotFound, ApiResponse{
				Success: false,
				Message: fmt.Sprintf("Unknown user with ID: %d", id),
			})
			return
		}

		log.Printf("Serving user %d\n", id)
		writeResponse(res, http.StatusOK, ApiResponse{
			Success: true,
			Message: "",
			Result:  []User{user},
		})
		return
	}

	var users []User
	for _, user := range database {
		users = append(users, user)
	}

	log.Println("Serving all users")
	writeResponse(res, http.StatusOK, ApiResponse{
		Success: true,
		Message: "",
		Result:  users,
	})
}
