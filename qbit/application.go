package qbit

import (
	"encoding/json"
	"io"
)

func (client *Client) GetVersion() (string, error) {
	ep := "/api/v2/app/version"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func (client *Client) GetWebAPIVersion() (string, error) {
	ep := "/api/v2/app/webapiVersion"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (client *Client) GetBuildInfo() (map[string]interface{}, error) {
	ep := "/api/v2/app/buildInfo"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var buildInfo map[string]interface{}
	err = json.Unmarshal(body, &buildInfo)
	if err != nil {
		return nil, err
	}
	return buildInfo, nil
}

func (client *Client) GetPreferences() (Preferences, error) {
	ep := "/api/v2/app/preferences"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return Preferences{}, err
	}

	var prefs Preferences
	if err := json.NewDecoder(resp.Body).Decode(&prefs); err != nil {
		return prefs, FailedToDecodeResponse(err)
	}

	return prefs, nil
}
func (client *Client) SetPreferences(pref map[string]string) error {
	ep := "/api/v2/app/setPreferences"
	jsonData, err := json.Marshal(pref)
	if err != nil {
		return err
	}
	jsonString := string(jsonData)
	opts := map[string]string{
		"json": jsonString,
	}
	_, err = client.PostFormData(ep, opts)
	if err != nil {
		return err
	}
	return nil
}
