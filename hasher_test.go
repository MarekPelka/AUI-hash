package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

type urlHashHandler func(w http.ResponseWriter, r *http.Request)

var (
	TEST_STRING = "String_for_testing"
	MD5_RESPONSE = "1f9be8d2262152abbf9c595fe8651ce9"
	SHA1_RESPONSE = "49f0edf87144e8aef8fcf43753cbd7a2497998b2"
	SHA224_RESPONSE = "82022c87bb14169295b5b13688404f013f4c39011c204ae358bff579"
	SHA256_RESPONSE = "77c307a66057925a284f6fe6346b5a89bd11e93be3a39e0da43b37fdf05d61d6"
	SHA384_RESPONSE = "40aa6fb476e83ac0a82aac3484da942a5fa417bdf376f115298cfd28b9d0093cd282fa678d8549f3624108c0a27fb7bb"
	SHA512_RESPONSE = "f9e6aa9514902a0362c64c9849b41bab1525d4d1732e8807de8a380015996eb6ab57e5a613845add6524f4cdd2dc5c9b8ac86343c1977eb8ae2fe150b8697771"
)

func HttpWraper(t *testing.T, function urlHashHandler, url string, response string)  {
	req, err := http.NewRequest("GET", url+TEST_STRING, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(function)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := response
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMD5(t *testing.T) {
	HttpWraper(t, md5Handler, MD5Url, MD5_RESPONSE)
}

func TestSHA1(t *testing.T) {
	HttpWraper(t, sha1Handler, SHA1Url, SHA1_RESPONSE)
}

func TestSHA224(t *testing.T) {
	HttpWraper(t, sha224Handler, SHA224Url, SHA224_RESPONSE)
}

func TestSHA256(t *testing.T) {
	HttpWraper(t, sha256Handler, SHA256Url, SHA256_RESPONSE)
}

func TestSHA384(t *testing.T) {
	HttpWraper(t, sha384Handler, SHA384Url, SHA384_RESPONSE)
}

func TestSHA512(t *testing.T) {
	HttpWraper(t, sha512Handler, SHA512Url, SHA512_RESPONSE)
}

