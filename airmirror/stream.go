package airmirror

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//AirMirror airplay manager
type AirMirror struct {
	tv *AirTV
}

//NewClient create instance
func NewClient(tv *AirTV) *AirMirror {
	am := &AirMirror{tv: tv}
	return am
}

//StreamReq start a mirror stream
func (a *AirMirror) StreamReq() error {
	log.Printf(a.tv.Address)

	streamInfo, err := os.Open("airmirror/post_stream.plist")
	if err != nil {
		return err
	}
	info, err := streamInfo.Stat()

	header := http.Header{
		"X-Apple-Device-ID": {"0xa4d1d2800b68"},
		"Content-Length":    {fmt.Sprintf("%d", info.Size())},
	}
	resp, err := a.doHTTPReq(http.MethodGet, "stream.xml", nil, header)
	if err != nil {
		return err
	}
	log.Printf("resp:%+v", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("body:%s", string(body))

	resp, err = a.doHTTPReq(http.MethodPost, "stream", streamInfo, header)
	if err != nil {
		return err
	}
	log.Printf("resp:%+v", resp)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	log.Printf("body:%s", string(body))
	return err
}

func (a *AirMirror) doHTTPReq(method, path string, body io.ReadSeeker, header http.Header) (*http.Response, error) {

	url := "http://" + a.tv.Address + "/" + path
	log.Printf("doHttpReq url:%s", url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = header
	log.Printf("header:%+v", header)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
