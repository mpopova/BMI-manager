package main

import (
	"fmt"
	"log"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"bytes"
	"strconv"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
    	userName string
	loginUrl string
    	registerUrl string
	getProfileInfoUrl string
	getPersonalStatUrl string
    	getAverageBMIUrl string
    	calculateBMIUrl string
)

func init() {
	// log.Println(fmt.Sprintf("%s", "INIT"))
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/getProfileInfo", getProfileInfo).Methods("POST")
	router.HandleFunc("/getPersonalStat", getPersonalStat).Methods("POST")
	router.HandleFunc("/getAverageBMI", getAverageBMI).Methods("POST")
	router.HandleFunc("/calculateBMI", calculateBMI).Methods("POST")

	server = httptest.NewServer(router)
	loginUrl = fmt.Sprintf("%s/login", server.URL)
    	registerUrl = fmt.Sprintf("%s/register", server.URL)
    	getProfileInfoUrl = fmt.Sprintf("%s/getProfileInfo", server.URL)
	getPersonalStatUrl = fmt.Sprintf("%s/getPersonalStat", server.URL)
    	getAverageBMIUrl = fmt.Sprintf("%s/getAverageBMI", server.URL)
}

func TestLoginHandler(t *testing.T){
	log.Println("TestLoginHandler");

 	data := url.Values{}
    	data.Set("name", "test")
    	data.Add("password", "test")

	request, err := http.NewRequest("POST", loginUrl, bytes.NewBufferString(data.Encode()))

 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.Request.URL.Path != "/calculate" {
		t.Errorf("Unsuccessfull login !!!", res.StatusCode)
	}
}

// func TestCalculateBMI(t *testing.T){
//     log.Println("TestCalculateBMI");

//     data := url.Values{}
//     data.Add("BMI", "20")

//     w := httptest.NewRecorder()
//     setSession("test", "M", "25", w)

//     request, err := http.NewRequest("POST", calculateBMIUrl, bytes.NewBufferString(data.Encode()))
    
//     request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//     request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
//     res, err := http.DefaultClient.Do(request)

//     if err != nil {
//         t.Error(err)
//     }

//     if res.Request.URL.Path != "/calculate" {
//         t.Errorf("Unsuccessfull calculation !!!", res.StatusCode)
//     }
// }

func TestRegisterHandler(t *testing.T){
    log.Println("TestRegisterHandler");

    data := url.Values{}
    data.Set("name", "TestTestov")
    data.Add("password", "testPass")
    data.Add("gender", "F")
    data.Add("age", "22")

    request, err := http.NewRequest("POST", registerUrl, bytes.NewBufferString(data.Encode()))

    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err)
    }

    if res.Request.URL.Path != "/" {
        t.Errorf("Unsuccessfull registration !!!", res.StatusCode)
    }
}

func TestGetProfileInfoUrl(t *testing.T){
    log.Println("TestGetProfileInfoUrl")

    data := url.Values{}
    data.Set("name", "test")
    data.Add("password", "test")
    w := httptest.NewRecorder()
    setSession("test", "M", "25", w)
    request, err := http.NewRequest("POST", getProfileInfoUrl, bytes.NewBufferString(data.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    getProfileInfo(w, request)
    
    if err != nil {
        t.Error(err)
    }
    
    if w.Code != http.StatusOK {
        t.Errorf("No response from profile URL %v", http.StatusOK)
    }
}

func TestGetPersonalStat(t *testing.T){
    log.Println("TestGetPersonalStat")

    data := url.Values{}
    data.Set("name", "test")
    data.Add("password", "test")
    w := httptest.NewRecorder()
    setSession("test", "M", "25", w)
    request, err := http.NewRequest("POST", getPersonalStatUrl, bytes.NewBufferString(data.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    getPersonalStat(w, request)
    
    if err != nil {
        t.Error(err)
    }
    
    if w.Code != http.StatusOK {
        t.Errorf("No response from personal stats URL %v", http.StatusOK)
    }
}

func TestGetAverageBMI(t *testing.T){
    log.Println("TestGetAverageBMI")

    data := url.Values{}
    data.Set("name", "test")
    data.Add("password", "test")
    w := httptest.NewRecorder()
    setSession("test", "M", "25", w)
    request, err := http.NewRequest("POST", getAverageBMIUrl, bytes.NewBufferString(data.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    getPersonalStat(w, request)
    
    if err != nil {
        t.Error(err)
    }
    
    if w.Code != http.StatusOK {
        t.Errorf("No response from average BMI URL %v", http.StatusOK)
    }
}

func TestGetUserInfo(t *testing.T){
    log.Println("TestGetUserInfo")

    data := url.Values{}
    data.Set("name", "test")
    data.Add("password", "test")
    w := httptest.NewRecorder()
    
    request, err := http.NewRequest("POST", getAverageBMIUrl, bytes.NewBufferString(data.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    setSession("test", "M", "25", w)
    getUserInfo(request)
    
    if err != nil {
        t.Error(err)
    }
    
    if w.Code != http.StatusOK {
        t.Errorf("No user info %v", http.StatusOK)
    }
}
