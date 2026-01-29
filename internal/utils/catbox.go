package utils

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"os"
)

// Upload file to Catbox (simple implementation)
func UploadToCatbox(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("fileToUpload", filepath)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(fw, f)
	if err != nil {
		return "", err
	}
	w.WriteField("reqtype", "fileupload")
	w.Close()

	resp, err := http.Post("https://catbox.moe/user/api.php", w.FormDataContentType(), &b)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("catbox upload failed")
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	return string(bodyBytes), nil
}