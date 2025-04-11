package qbit

import (
	"encoding/json"
	"fmt"
)

func (client *Client) List(opts map[string]string) ([]Torrent, error) {
	ep := "/api/v2/torrents/info"
	resp, err := client.Get(ep, opts)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var torrents []Torrent
	if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
		return nil, FailedToDecodeResponse(err)
	}
	return torrents, nil
}

func (client *Client) Properties(hash string) (TorrentGeneric, error) {
	ep := "/api/v2/torrents/properties"
	opts := map[string]string{
		"hash": hash,
	}
	resp, err := client.Get(ep, opts)
	var torrentGenInfo TorrentGeneric
	if err != nil {
		return torrentGenInfo, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&torrentGenInfo); err != nil {
		return torrentGenInfo, FailedToDecodeResponse(err)
	}
	return torrentGenInfo, nil
}

func (client *Client) Trackers(hash string) ([]Tracker, error) {
	ep := "/api/v2/torrents/trackers"
	opts := map[string]string{
		"hash": hash,
	}
	resp, err := client.Get(ep, opts)
	if err != nil {
		return nil, err
	}
	var trackers []Tracker
	if err := json.NewDecoder(resp.Body).Decode(&trackers); err != nil {
		return trackers, FailedToDecodeResponse(err)
	}
	return trackers, nil
}

func (client *Client) WebSeeds(hash string) ([]string, error) {
	ep := "/api/v2/torrents/webseeds"
	opts := map[string]string{"hash": hash}
	resp, err := client.Get(ep, opts)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var webseeds []struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&webseeds); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}
	urls := make([]string, len(webseeds))
	for i, item := range webseeds {
		urls[i] = item.URL
	}
	return urls, nil
}

func (client *Client) Files(hash string, indexes string) ([]File, error) {
	ep := "/api/v2/torrents/files"
	opts := map[string]string{
		"hash": hash,
	}
	if indexes != "" {
		opts["indexes"] = indexes
	}
	resp, err := client.Get(ep, opts)
	if err != nil {
		return nil, err
	}
	var files []File
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return files, FailedToDecodeResponse(err)
	}
	return files, nil
}

func (client *Client) PieceStates(hash string) ([]int, error) {
	ep := "/api/v2/torrents/pieceStates"
	opts := map[string]string{
		"hash": hash,
	}
	resp, err := client.Get(ep, opts)
	if err != nil {
		return nil, err
	}
	var pieceStates []int
	if err := json.NewDecoder(resp.Body).Decode(&pieceStates); err != nil {
		return pieceStates, FailedToDecodeResponse(err)
	}
	return pieceStates, nil
}

func (client *Client) DeleteTorrent(hashes string, deleteFiles bool) error {
	ep := "/api/v2/torrents/delete"
	deleteFiles_ := ""
	if deleteFiles {
		deleteFiles_ = "true"
	} else {
		deleteFiles_ = "false"
	}
	opts := map[string]string{
		"hashes":      hashes,
		"deleteFiles": deleteFiles_,
	}
	_, err := client.PostFormData(ep, opts)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) IncreasePrio(hashes string) error {
	ep := "/api/v2/torrents/increasePrio"
	opts := map[string]string{
		"hashes": hashes,
	}
	_, err := client.Get(ep, opts)
	return err
}

func (client *Client) DecreasePrio(h string) error {
	ep := "/api/v2/torrents/decreasePrio"
	opts := map[string]string{
		"hashes": h,
	}
	_, err := client.Get(ep, opts)
	return err
}

func (client *Client) TopPrio(h string) error {
	ep := "/api/v2/torrents/topPrio"
	opts := map[string]string{
		"hashes": h,
	}
	_, err := client.Get(ep, opts)
	return err
}

func (client *Client) BottomPrio(h string) error {
	ep := "/api/v2/torrents/bottomPrio"
	opts := map[string]string{
		"hashes": h,
	}
	_, err := client.Get(ep, opts)
	return err
}

func (client *Client) FilePrio(h, id string, priority int) error {
	ep := "/api/v2/torrents/filePrio"
	opts := map[string]string{
		"hash":     h,
		"id":       id,
		"priority": fmt.Sprintf("%d", priority),
	}
	_, err := client.Get(ep, opts)
	return err
}

func (client *Client) DownloadLimit(h string) (map[string]int, error) {
	ep := "/api/v2/torrents/downloadLimit"
	opts := map[string]string{
		"hashes": h,
	}
	resp, err := client.Post(ep, opts)
	if err != nil {
		return nil, err
	}

	var limits map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&limits); err != nil {
		return nil, FailedToDecodeResponse(err)
	}

	return limits, nil
}

func (client *Client) SetDownloadLimit(h string, limit int) error {
	ep := "/api/v2/torrents/setDownloadLimit"
	opts := map[string]string{
		"hashes": h,
		"limit":  fmt.Sprintf("%d", limit),
	}
	_, err := client.Post(ep, opts)
	return err
}

func (client *Client) SetShareLimits(h string, ratioLimit float64, seedingTimeLimit int, inactiveSeedingTimeLimit int) error {
	ep := "/api/v2/torrents/setShareLimits"
	opts := map[string]string{
		"hashes":                   h,
		"ratioLimit":               fmt.Sprintf("%f", ratioLimit),
		"seedingTimeLimit":         fmt.Sprintf("%d", seedingTimeLimit),
		"inactiveSeedingTimeLimit": fmt.Sprintf("%d", inactiveSeedingTimeLimit),
	}
	_, err := client.Post(ep, opts)
	return err
}

func (client *Client) UploadLimit(h string) (map[string]int, error) {
	ep := "/api/v2/torrents/uploadLimit"
	opts := map[string]string{
		"hashes": h,
	}
	resp, err := client.Post(ep, opts)
	if err != nil {
		return nil, err
	}

	var limits map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&limits); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return limits, nil
}

func (client *Client) SetLocation(h string, location string) error {
	ep := "/api/v2/torrents/setLocation"
	opts := map[string]string{
		"hashes":   h,
		"location": location,
	}
	_, err := client.Post(ep, opts)
	return err
}

func (client *Client) SetCategory(h, category string) error {
	ep := "/api/v2/torrents/setCategory"
	opts := map[string]string{
		"hashes":   h,
		"category": category,
	}
	_, err := client.Post(ep, opts)
	return err
}

func (client *Client) GetCategories() (map[string]struct {
	Name     string `json:"name"`
	SavePath string `json:"savePath"`
}, error) {
	ep := "/api/v2/torrents/categories"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return nil, err
	}

	var categories map[string]struct {
		Name     string `json:"name"`
		SavePath string `json:"savePath"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return categories, nil
}

func (client *Client) EditCategory(category, savePath string) error {
	ep := "/api/v2/torrents/editCategory"
	opts := map[string]string{
		"category": category,
		"savePath": savePath,
	}
	_, err := client.Post(ep, opts)
	return err
}

func (client *Client) RemoveCategories(categories string) error {
	ep := "/api/v2/torrents/removeCategories"
	opts := map[string]string{
		"categories": categories,
	}
	_, err := client.Post(ep, opts)
	return err
}

func (client *Client) AddTags(h, tags string) error {
	ep := "/api/v2/torrents/addTags"
	opts := map[string]string{
		"hashes": h,
		"tags":   tags,
	}
	_, err := client.Post(ep, opts)
	return err
}

func (client *Client) RemoveTags(h, tags string) error {
	ep := "/api/v2/torrents/removeTags"
	opts := map[string]string{
		"hashes": h,
		"tags":   tags,
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) GetTags() ([]string, error) {
	ep := "/api/v2/torrents/tags"
	resp, err := client.Get(ep, nil)
	if err != nil {
		return nil, err
	}

	var tags []string
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return tags, nil
}

func (client *Client) CreateTags(tags string) error {
	ep := "/api/v2/torrents/createTags"
	opts := map[string]string{
		"tags": tags,
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) DeleteTags(tags string) error {
	ep := "/api/v2/torrents/deleteTags"
	opts := map[string]string{
		"tags": tags,
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) ToggleFirstLastPiecePrio(hashes string) error {
	ep := "/api/v2/torrents/toggleFirstLastPiecePrio"
	opts := map[string]string{
		"hashes": hashes,
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) SetForceStart(hashes string, value bool) error {
	ep := "/api/v2/torrents/setForceStart"
	opts := map[string]string{
		"hashes": hashes,
		"value":  fmt.Sprintf("%v", value),
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) SetSuperSeeding(hashes string, value bool) error {
	ep := "/api/v2/torrents/setSuperSeeding"
	opts := map[string]string{
		"hashes": hashes,
		"value":  fmt.Sprintf("%v", value),
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) RenameFile(hash, oldPath, newPath string) error {
	ep := "/api/v2/torrents/renameFile"
	opts := map[string]string{
		"hash":    hash,
		"oldPath": oldPath,
		"newPath": newPath,
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) RenameFolder(hash, oldPath, newPath string) error {
	ep := "/api/v2/torrents/renameFolder"
	opts := map[string]string{
		"hash":    hash,
		"oldPath": oldPath,
		"newPath": newPath,
	}
	_, err := client.PostFormData(ep, opts)
	return err
}

func (client *Client) AddTorrentWithLink(torrentURLs []string, savePath, category string) error {
	formData := make(map[string]string)
	for _, torrentURL := range torrentURLs {
		formData["urls"] += torrentURL + "\n"
	}

	formData["savepath"] = savePath
	formData["category"] = category

	_, err := client.PostFormData("/api/v2/torrents/add", formData)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) AddTorrentWithFile(torrentFiles []string, savePath, category string) error {
	formData := make(map[string]string)
	formData["savepath"] = savePath
	formData["category"] = category

	for _, filePath := range torrentFiles {
		_, err := client.PostFileWithForm("/api/v2/torrents/add", formData, "torrents", filePath)
		return err
	}
	return nil
}
