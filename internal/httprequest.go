package internal

import (
	"errors"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strings"

	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/perflogger"
)

// func HttpGet(url string) (string, error) {
// 	logger := perflogger.GetInstance()
// 	logger.Start()

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	result, err := ioutil.ReadAll(resp.Body)

// 	logger.StopAndLog("GET", url)

// 	return string(result), err
// }

func HttpGet(url string) (string, error) {
	var proxyURL = config.ProxyURL
	logger := perflogger.GetInstance()
	logger.Start()

	var resp *http.Response
	var err error
	if proxyURL != "" {
		proxy, err := u.Parse(proxyURL)
		if err != nil {
			return "", err
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
		resp, err = client.Get(url)
	} else {
		resp, err = http.Get(url)
	}
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", errors.New("response error")
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("GET", url)
	return string(result), err
}

// func HttpPost(url string, body string) (string, error) {
// 	logger := perflogger.GetInstance()
// 	logger.Start()

// 	resp, err := http.Post(url, "application/json", strings.NewReader(body))
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	result, err := ioutil.ReadAll(resp.Body)

// 	logger.StopAndLog("POST", url)

// 	return string(result), err
// }

func HttpPost(url string, body string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	var resp *http.Response
	var err error
	var proxyURL = config.ProxyURL
	if proxyURL != "" {
		proxy, err := u.Parse(proxyURL)
		if err != nil {
			return "", err
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
		resp, err = client.Post(url, "application/json", strings.NewReader(body))
	} else {
		resp, err = http.Post(url, "application/json", strings.NewReader(body))
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("POST", url)

	return string(result), err
}
