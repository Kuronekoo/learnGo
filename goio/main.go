package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"git.garena.com/shopee/bg-logistics/go/gocommon/logger"
	"git.garena.com/shopee/bg-logistics/tianlu/wms-v2/apps/constant"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"time"
)

func main() {

}
func combineToOneZipAndUnzip() {
	var cOut []byte
	var err error
	var stderr bytes.Buffer
	var cmd *exec.Cmd
	country := "CN"
	dateStr := time.Now().Format("20060102")
	fileFolder := fmt.Sprintf("/tmp/ff/%s/%s_%d/", country, dateStr, time.Now().Unix())
	zipFileName := fmt.Sprintf("%s_%s_full.zip", country, dateStr)
	csvFileName := fmt.Sprintf("%s_%s.csv", country, dateStr)
	command1 := fmt.Sprintf(`cd %s && cat * >> %s`, fileFolder, zipFileName)
	command2 := fmt.Sprintf(`cd %s && tar -xzvf %s && mv *.csv %s`, fileFolder, zipFileName, csvFileName)
	if !fileExists(fileFolder + zipFileName) {
		fmt.Println(command1)
		cmd = exec.Command("/bin/sh", "-c", command1)
		cmd.Stderr = &stderr
		cOut, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error() + stderr.String())
		}
		fmt.Println(string(cOut))
	}
	if !fileExists(fileFolder + csvFileName) {
		fmt.Println(command2)
		cmd = exec.Command("/bin/sh", "-c", command2)
		cmd.Stderr = &stderr
		cOut, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error() + stderr.String())
		}
		fmt.Println(string(cOut))
	}
}

// 判断所给路径文件/文件夹是否存在
func fileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func downloadAndUnzip(downloadUrls []string) error {
	for _, url := range downloadUrls {
		fileByte, err := OriginGet(url, nil, 30)
		if err != nil {
			return err
		}
		r, err := zip.NewReader(bytes.NewReader(fileByte), int64(len(fileByte)))
		if err != nil {
			return err
		}
		for _, f := range r.File {
			ioFile, err := f.Open()
			if err != nil {
				return err
			}
			s := bufio.NewScanner(ioFile)
			if !s.Scan() {
				continue
			}
			for s.Scan() {
				data := s.Text()
				fmt.Println(data)
			}
			err = ioFile.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func OriginGet(url string, headerMap map[string]string, timeOut int64) ([]byte, error) {
	//创建client
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeOut),
	}
	//组装请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	originAddHeader(req, headerMap)
	//请求
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp == nil {
		return nil, errors.New("http resp empty for request ")
	}
	//读取字节流
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	statusCode := resp.StatusCode
	logger.LogInfof("request url : %v , res code : %v ", url, statusCode)
	if statusCode != http.StatusOK {
		errcode := constant.ErrHttp5xx
		if 300 <= statusCode && statusCode < 400 {
			errcode = constant.ErrHttp3xx
		} else if 400 <= statusCode && statusCode < 500 {
			errcode = constant.ErrHttp4xx
		}
		logger.LogErrorf("request url : %v  fail, status %v", url, statusCode)
		return nil, errors.New(fmt.Sprintf("request fail, status %v", errcode))
	}
	return body, nil

}

func OriginGetFileNameAndDownload(url string, headerMap map[string]string, timeOut int64) (string, []byte, error) {
	//创建client
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeOut),
	}
	//组装请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", nil, err
	}
	originAddHeader(req, headerMap)
	//请求
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp == nil {
		return "", nil, errors.New("http resp empty for request ")
	}
	contentDisposition := resp.Header.Get("Content-Disposition")
	//读取字节流
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	statusCode := resp.StatusCode
	logger.LogInfof("download from url : %v , res code : %v ", url, statusCode)
	if statusCode != http.StatusOK {
		errcode := constant.ErrHttp5xx
		if 300 <= statusCode && statusCode < 400 {
			errcode = constant.ErrHttp3xx
		} else if 400 <= statusCode && statusCode < 500 {
			errcode = constant.ErrHttp4xx
		}
		logger.LogErrorf("download from : %v  fail, status %v", url, statusCode)
		return "", nil, errors.New(fmt.Sprintf("request fail, status %v", errcode))
	}
	if len(contentDisposition) > 0 {
		reg := regexp.MustCompile(`(filename=")(.*)(")`)
		result := reg.FindAllStringSubmatch(contentDisposition, -1)
		fileName := result[0][2]
		return fileName, body, nil
	}
	return "", body, nil

}

func originAddHeader(req *http.Request, headerMap map[string]string) {
	for key, value := range headerMap {
		req.Header.Add(key, value)
	}
}
