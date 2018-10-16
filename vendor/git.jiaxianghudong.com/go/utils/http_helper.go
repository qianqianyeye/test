package utils

import (
	"bytes"
	"strings"

	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get
func Get(apiUrl string, parm map[string]string, header map[string]string, isHttps bool) ([]byte, error) {

	if len(parm) > 0 {
		apiUrl = fmt.Sprintf("%s%s", apiUrl, "?")
		p := ""
		for k, v := range parm {
			if p == "" {
				p = fmt.Sprintf("%s=%s", k, v)
			} else {
				p = fmt.Sprintf("%s&%s=%s", p, k, v)
			}
		}
		apiUrl = fmt.Sprintf("%s%s", apiUrl, p)
	}

	client := &http.Client{}

	if isHttps {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	reqest, _ := http.NewRequest("GET", apiUrl, nil)

	for k, v := range header {
		reqest.Header.Set(k, v)
	}

	response, err := client.Do(reqest)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

// post
func Post(apiUrl string, data []byte, header map[string]string, isHttps bool) ([]byte, error) {

	client := &http.Client{}

	if isHttps {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	reqest, _ := http.NewRequest("POST", apiUrl, bytes.NewReader(data))

	for k, v := range header {
		reqest.Header.Set(k, v)
	}

	response, err := client.Do(reqest)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

// post
func PostMap(apiUrl string, parm map[string]string, header map[string]string, isHttps bool) ([]byte, error) {

	data := url.Values{}
	for k, v := range parm {
		data.Set(k, v)
	}

	reqParams := ioutil.NopCloser(strings.NewReader(data.Encode()))
	client := &http.Client{}

	if isHttps {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	reqest, _ := http.NewRequest("POST", apiUrl, reqParams)

	for k, v := range header {
		reqest.Header.Set(k, v)
	}

	response, err := client.Do(reqest)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

// 获取远程ip
func GetRemoteIP(r *http.Request) string {
	addr := r.Header.Get("Remote_addr")
	if addr == "" {
		addr = r.RemoteAddr
	}

	return strings.Split(addr, ":")[0]
}
