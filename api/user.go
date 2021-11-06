package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// handleUserProfile returns the users profile if it matches their session
// @Summary Get Profile
// @Description get a users profile
// @Tags user
// @Produce  json
// @Param id path int false "the user ID"
// @Success 200 object standardJsonResponse{data=model.User}
// @Failure 403 object standardJsonResponse{}
// @Failure 500 object standardJsonResponse{}
// @Router /users/{id} [get]
func (a *api) handleUserProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		UserID := vars["id"]

		UserCookieID := r.Context().Value(contextKeyUserID).(string)
		UserType := r.Context().Value(contextKeyUserType).(string)
		if UserID != UserCookieID && UserType != "adminUserType" {
			Failure(w, r, http.StatusForbidden, Errorf(EINVALID, "INVALID_USER"))
			return
		}

		User, UserErr := a.db.GetUser(UserID)
		if UserErr != nil {
			Failure(w, r, http.StatusInternalServerError, UserErr)
			return
		}

		Success(w, r, http.StatusOK, User, nil)
	}
}

// handleUserProfileUpdate attempts to update users profile
// @Summary Update Profile
// @Description Update a users profile
// @Tags user
// @Produce  json
// @Param id path int false "the user ID"
// @Success 200 object standardJsonResponse{data=model.User}
// @Failure 403 object standardJsonResponse{}
// @Failure 500 object standardJsonResponse{}
// @Router /users/{id} [put]
func (a *api) handleUserProfileUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		keyVal := getJSONRequestBody(r, w)
		UserName := keyVal["name"].(string)
		UserAvatar := keyVal["avatar"].(string)
		NotificationsEnabled, _ := keyVal["notificationsEnabled"].(bool)
		Country := keyVal["country"].(string)
		Locale := keyVal["locale"].(string)
		Company := keyVal["company"].(string)
		JobTitle := keyVal["jobTitle"].(string)

		UserID := vars["id"]
		UserCookieID := r.Context().Value(contextKeyUserID).(string)
		UserType := r.Context().Value(contextKeyUserType).(string)
		if UserID != UserCookieID && UserType != "adminUserType" {
			Failure(w, r, http.StatusForbidden, Errorf(EINVALID, "INVALID_USER"))
			return
		}

		updateErr := a.db.UpdateUserProfile(UserID, UserName, UserAvatar, NotificationsEnabled, Country, Locale, Company, JobTitle)
		if updateErr != nil {
			Failure(w, r, http.StatusInternalServerError, updateErr)
			return
		}

		user, UserErr := a.db.GetUser(UserID)
		if UserErr != nil {
			Failure(w, r, http.StatusInternalServerError, UserErr)
			return
		}

		Success(w, r, http.StatusOK, user, nil)
	}
}

// handleUserDelete attempts to delete a users account
// @Summary Delete User
// @Description Deletes a user
// @Tags user
// @Produce  json
// @Param id path int false "the user ID"
// @Success 200 object standardJsonResponse{}
// @Failure 403 object standardJsonResponse{}
// @Failure 500 object standardJsonResponse{}
// @Router /users/{id} [delete]
func (a *api) handleUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		UserID := vars["id"]
		UserCookieID := r.Context().Value(contextKeyUserID).(string)
		UserType := r.Context().Value(contextKeyUserType).(string)
		if UserID != UserCookieID && UserType != "adminUserType" {
			Failure(w, r, http.StatusForbidden, Errorf(EINVALID, "INVALID_USER"))
			return
		}

		User, UserErr := a.db.GetUser(UserID)
		if UserErr != nil {
			Failure(w, r, http.StatusInternalServerError, UserErr)
			return
		}

		updateErr := a.db.DeleteUser(UserID)
		if updateErr != nil {
			Failure(w, r, http.StatusInternalServerError, updateErr)
			return
		}

		a.email.SendDeleteConfirmation(User.UserName, User.UserEmail)

		a.clearUserCookies(w)

		Success(w, r, http.StatusOK, nil, nil)
	}
}

// handleGetActiveCountries gets a list of registered users countries
// @Summary Get Active Countries
// @Description get a list of users countries
// @Produce  json
// @Success 200 object standardJsonResponse{[]string}
// @Failure 500 object standardJsonResponse{}
// @Router /active-countries [get]
func (a *api) handleGetActiveCountries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		countries, err := a.db.GetActiveCountries()

		if err != nil {
			Failure(w, r, http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Cache-Control", "max-age=3600") // cache for 1 hour just to decrease load
		Success(w, r, http.StatusOK, countries, nil)
	}
}
