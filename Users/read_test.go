package Users

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetListUsersWithoutToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:6000/users/list", nil)
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
	t.Log("GetListUsersWithoutTokenResponse: ", jsonStr)
}

func TestGetListUsersWithErrorToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:6000/users/list", nil)
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
	t.Log("GetListUsersWithErrorTokenResponse: ", jsonStr)
}

func TestGetListUsers(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:6000/users/list", nil)
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
	t.Log("GetListUsersResponse: ", jsonStr)
}

func TestGetUserWithoutToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:6000/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
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
	t.Log("GetUserWithoutTokenResponse: ", jsonStr)
}

func TestGetUserWithErrorToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:6000/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
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
	t.Log("TestGetUserWithErrorTokenResponse: ", jsonStr)
}

func TestGetUser(t *testing.T) {

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
	t.Log("TestGetUserResponse: ", jsonStr)
}
