package Users

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDeleteUserWithoutToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:6000/delete/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
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
	t.Log("DeleteUserWithoutTokenResponse: ", jsonStr)
}

func TestDeleteUserWithErrorToken(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:6000/delete/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
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
	t.Log("DeleteUserWithErrorTokenResponse: ", jsonStr)
}

func TestDeleteUser(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:6000/delete/user/OYjBPyWgtZea9b0ThORj8sbU9YtLhGBNh72gQkdMAJAcy1ZzqkZpVtzE6Ne6peu2E2Jp4TXU5JfYTLLpBGHmQqrOBuMVejRDUnJC", nil)
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
	t.Log("DeleteUserResponse: ", jsonStr)
}
