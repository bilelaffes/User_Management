package Users

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type userUpdate struct {
	Age  json.Number `json:"age"`
	Name string      `json:"name"`
	Data string      `json:"data"`
}

type userUpdateWithErrorField struct {
	Filed string `json:"field"`
}

func TestUpdateUserWithoutToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost:6000/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	jsonStr := string(body)
	t.Log("UpdateUserWithoutTokenResponse: ", jsonStr)
}

func TestUpdateUserWithErrorToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost:6000/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
	if err != nil {
		t.Error(err)
	}

	req.Header = http.Header{
		"Authorization": {"token"},
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	jsonStr := string(body)
	t.Log("UpdateUserWithErrorTokenResponse: ", jsonStr)
}

func TestUpdateUser(t *testing.T) {

	var test_update userUpdate = userUpdate{
		Age:  "50",
		Name: "Linus Torvalds",
		Data: "Test_file_change"}
	body, _ := json.Marshal(test_update)

	client := http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost:6000/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}

	req.Header = http.Header{
		"Authorization": {Token},
		"content-type":  {"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	jsonStr := string(body)
	t.Log("UpdateUserResponse: ", jsonStr)
}

func TestUpdateUserWithFieldDoesNotExistToUser(t *testing.T) {

	var test_update userUpdateWithErrorField = userUpdateWithErrorField{
		Filed: "Test"}
	body, _ := json.Marshal(test_update)

	client := http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost:6000/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}

	req.Header = http.Header{
		"Authorization": {Token},
		"content-type":  {"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
}

func TestGetUserToValidateUpdate(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:6000/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
	if err != nil {
		t.Error(err)
	}

	req.Header = http.Header{
		"Authorization": {Token},
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	jsonStr := string(body)
	t.Log("GetUserToValidateUpdateResponse: ", jsonStr)
}
