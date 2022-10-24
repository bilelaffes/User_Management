package Users

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {

	body, err := ioutil.ReadFile("DataSet.json")
	if err != nil {
		t.Error("Error to read file")
	}

	resp, err := http.Post("http://localhost:6000/add/users", "application/json", bytes.NewBuffer(body))

	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		jsonStr := string(body)
		t.Log("Response: ", jsonStr)
	} else {
		t.Log("Get failed with error: ", resp.Status)
	}
}
