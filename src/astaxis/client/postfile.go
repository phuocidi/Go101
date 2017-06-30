/*
Package client implements a impersonate real web user for astaxis
func main() {
    target_url := "http://localhost:9090/upload"
    filename := "./astaxie.pdf"
    postFile(filename, target_url)
}
*/
package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func PostFile(filename string, targetURL string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// IMPORTANT
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("[client.PostFile error]: writting to buffer", err)
		return err
	}

	// open file handler
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("[client.PostFile error]: open file", err)
		return err
	}

	// io copy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("[client.PostFile error]: copy file", err)
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetURL, contentType, bodyBuf)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}
