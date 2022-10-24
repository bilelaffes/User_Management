package Users

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type login struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

var Token string

func TestLoginWithErrorPassword(t *testing.T) {

	var test_login login = login{
		Id:       "OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC",
		Password: "G4nzkSBFPcqVotYVFT"}

	client := http.Client{}
	body, _ := json.Marshal(test_login)
	req, err := http.NewRequest("POST", "http://localhost:6000/login", bytes.NewBuffer(body))
	req.Header = http.Header{
		"content-type": {"application/json"},
	}
	if err != nil {
		t.Error(err)
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
	t.Log("LoginWithErrorPassword: ", jsonStr)
}

func TestLoginWithGoodPassword(t *testing.T) {

	var test_login login = login{
		Id:       "OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC",
		Password: "G4nzkSBFPcqVotYVFTkx"}

	client := http.Client{}
	body, _ := json.Marshal(test_login)
	req, err := http.NewRequest("POST", "http://localhost:6000/login", bytes.NewBuffer(body))

	req.Header = http.Header{
		"content-type": {"application/json"},
	}

	if err != nil {
		t.Error(err)
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

	var token map[string]string

	if err := json.Unmarshal(body, &token); err != nil {
		t.Error(err.Error())
	}
	Token = token["token"]
	t.Log("LoginWithGoodPassword: ", jsonStr)
}
