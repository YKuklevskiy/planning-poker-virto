package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// handleUserAPIKeys handles getting user API keys
// @Summary Get API Keys
// @Description get list of API keys for the user
// @Tags apikey
// @Produce  json
// @Param userId path string true "the user ID to get API keys for"
// @Success 200 object standardJsonResponse{data=[]model.APIKey}
// @Failure 403 object standardJsonResponse{}
// @Failure 500 object standardJsonResponse{}
// @Security ApiKeyAuth
// @Router /users/{userId}/apikeys [get]
func (a *api) handleUserAPIKeys() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		UserID := vars["userId"]

		APIKeys, keysErr := a.db.GetUserApiKeys(UserID)
		if keysErr != nil {
			a.Failure(w, r, http.StatusInternalServerError, keysErr)
			return
		}

		a.Success(w, r, http.StatusOK, APIKeys, nil)
	}
}

type apikeyGenerateRequestBody struct {
	Name string `json:"name"`
}

// handleAPIKeyGenerate handles generating an API key for a user
// @Summary Generate API Key
// @Description Generates an API key for the user
// @Tags apikey
// @Produce  json
// @Param userId path string true "the user ID to generate API key for"
// @Param key body apikeyGenerateRequestBody true "new api key object"
// @Success 200 object standardJsonResponse{data=model.APIKey}
// @Failure 403 object standardJsonResponse{}
// @Failure 500 object standardJsonResponse{}
// @Security ApiKeyAuth
// @Router /users/{userId}/apikeys [post]
func (a *api) handleAPIKeyGenerate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		UserID := vars["userId"]

		body, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			a.Failure(w, r, http.StatusBadRequest, Errorf(EINVALID, bodyErr.Error()))
			return
		}

		var k = apikeyGenerateRequestBody{}
		jsonErr := json.Unmarshal(body, &k)
		if jsonErr != nil {
			a.Failure(w, r, http.StatusBadRequest, Errorf(EINVALID, jsonErr.Error()))
			return
		}

		APIKeys, keysErr := a.db.GetUserApiKeys(UserID)
		if keysErr != nil {
			a.Failure(w, r, http.StatusInternalServerError, keysErr)
			return
		}

		if len(APIKeys) == a.config.UserAPIKeyLimit {
			a.Failure(w, r, http.StatusForbidden, Errorf(EINVALID, "USER_APIKEY_LIMIT_REACHED"))
			return
		}

		APIKey, keyErr := a.db.GenerateApiKey(UserID, k.Name)
		if keyErr != nil {
			a.Failure(w, r, http.StatusInternalServerError, keyErr)
			return
		}

		a.Success(w, r, http.StatusOK, APIKey, nil)
	}
}

type apikeyUpdateRequestBody struct {
	Active bool `json:"active"`
}

// handleUserAPIKeyUpdate handles updating a users API key
// @Summary Update API Key
// @Description Updates the API key of the user
// @Tags apikey
// @Produce  json
// @Param userId path string true "the user ID"
// @Param keyID path string true "the API Key ID to update"
// @Param key body apikeyUpdateRequestBody true "api key object to update"
// @Success 200 object standardJsonResponse{data=[]model.APIKey}
// @Failure 403 object standardJsonResponse{}
// @Failure 500 object standardJsonResponse{}
// @Security ApiKeyAuth
// @Router /users/{userId}/apikeys/{keyID} [put]
func (a *api) handleUserAPIKeyUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		UserID := vars["userId"]
		APK := vars["keyID"]

		body, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			a.Failure(w, r, http.StatusBadRequest, Errorf(EINVALID, bodyErr.Error()))
			return
		}

		var k = apikeyUpdateRequestBody{}
		jsonErr := json.Unmarshal(body, &k)
		if jsonErr != nil {
			a.Failure(w, r, http.StatusBadRequest, Errorf(EINVALID, jsonErr.Error()))
			return
		}

		APIKeys, keysErr := a.db.UpdateUserApiKey(UserID, APK, k.Active)
		if keysErr != nil {
			a.Failure(w, r, http.StatusInternalServerError, keysErr)
			return
		}

		a.Success(w, r, http.StatusOK, APIKeys, nil)
	}
}

// handleUserAPIKeyDelete handles deleting a users API key
// @Summary Delete API Key
// @Description Deletes the API key
// @Tags apikey
// @Produce  json
// @Param userId path int false "the user ID"
// @Param keyID path int false "the API Key ID to update"
// @Success 200 object standardJsonResponse{data=[]model.APIKey}
// @Failure 403 object standardJsonResponse{}
// @Failure 500 object standardJsonResponse{}
// @Security ApiKeyAuth
// @Router /users/{userId}/apikeys/{keyID} [delete]
func (a *api) handleUserAPIKeyDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		UserID := vars["userId"]
		APK := vars["keyID"]

		APIKeys, keysErr := a.db.DeleteUserApiKey(UserID, APK)
		if keysErr != nil {
			a.Failure(w, r, http.StatusInternalServerError, keysErr)
			return
		}

		a.Success(w, r, http.StatusOK, APIKeys, nil)
	}
}
