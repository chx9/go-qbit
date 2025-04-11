package qbit

import (
	"encoding/json"
	"fmt"
)

func (client *Client) GetTransferInfo() (*TransferInfo, error) {
	ep := "/api/v2/transfer/info"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return nil, err
	}
	var transferInfo TransferInfo
	if err := json.NewDecoder(resp.Body).Decode(&transferInfo); err != nil {
		return nil, FailedToDecodeResponse(err)
	}
	return &transferInfo, nil
}

func (client *Client) GetSpeedLimitsMode() (int, error) {
	ep := "/api/v2/transfer/speedLimitsMode"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return 0, err
	}
    var mod struct {
        Enabled int `json:"enabled"`
    }
	if err := json.NewDecoder(resp.Body).Decode(&mod); err != nil {
		return 0, FailedToDecodeResponse(err)
	}
	return mod.Enabled, nil
}

func (client *Client) ToggleSpeedLimitsMode() error {
	ep := "/api/v2/transfer/toggleSpeedLimitsMode"
	_, err := client.PostFormData(ep, nil)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetGlobalDownloadLimit() (int, error) {
	ep := "/api/v2/transfer/downloadLimit"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return 0, err
	}
	var limit struct {
		Limit int `json:"limit"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&limit); err != nil {
		return 0, FailedToDecodeResponse(err)
	}
	return limit.Limit, nil
}

func (client *Client) SetGlobalDownloadLimit(limit int) error {
	ep := "/api/v2/transfer/setDownloadLimit"
	opts := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	_, err := client.PostFormData(ep, opts)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetGlobalUploadLimit() (int, error) {
	ep := "/api/v2/transfer/uploadLimit"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return 0, err
	}
	var limit Limit
	if err := json.NewDecoder(resp.Body).Decode(&limit); err != nil {
		return 0, FailedToDecodeResponse(err)
	}
	return limit.Limit, nil
}

func (client *Client) GetUploadLimit(limit int) error {
	ep := "/api/v2/transfer/setUploadLimit"
	opts := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}
	_, err := client.PostFormData(ep, opts)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) BanPeers(peers string) error {
	ep := "/api/v2/transfer/banPeers"
	opts := map[string]string{
		"peers": peers,
	}
	_, err := client.PostFormData(ep, opts)
	if err != nil {
		return err
	}
	return nil
}
