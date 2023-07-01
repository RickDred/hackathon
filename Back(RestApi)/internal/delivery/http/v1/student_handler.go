package v1

//func RegisterStudent(w http.ResponseWriter, r *http.Request) {
//	var input struct {
//		Id       int64  `json:"id"`
//		Name     string `json:"name"`
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//
//	// Parse the request body
//	err := json.NewDecoder(r.Body).Decode(&input)
//	if err != nil {
//		http.Error(w, "Invalid request body", http.StatusBadRequest)
//		return
//	}
//
//	// Call the student service to register the student
//	err = services.RegisterStudent(&input)
//	if err != nil {
//		http.Error(w, "Failed to register student", http.StatusInternalServerError)
//		return
//	}
//
//	WriteJSON(w, http.StatusOK, input, nil)
//}
//
//// LoginStudent handles the user login endpoint
//func LoginStudent(w http.ResponseWriter, r *http.Request) {
//	var credentials models.UserCredentials
//
//	// Parse the request body
//	err := json.NewDecoder(r.Body).Decode(&credentials)
//	if err != nil {
//		http.Error(w, "Invalid request body", http.StatusBadRequest)
//		return
//	}
//
//	// Call the user service to authenticate the user
//	token, err := services.AuthenticateUser(credentials.Email, credentials.Password)
//	if err != nil {
//		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
//		return
//	}
//
//	// Return the generated token as the response
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(token)
//}
