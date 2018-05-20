package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

var (
	MD5Url    = "/md5/"
	SHA1Url   = "/sha1/"
	SHA224Url = "/sha224/"
	SHA256Url = "/sha256/"
	SHA384Url = "/sha384/"
	SHA512Url = "/sha512/"
	ALLUrl    = "/all/"
)

func encode(b []byte) string {
	return hex.EncodeToString(b[:])
}

func md5Handler(w http.ResponseWriter, r *http.Request) {
	h := md5.Sum([]byte(r.URL.Path[len(MD5Url):]))
	fmt.Fprintf(w, "%s", encode(h[:]))
}

func sha1Handler(w http.ResponseWriter, r *http.Request) {
	h := sha1.Sum([]byte(r.URL.Path[len(SHA1Url):]))
	fmt.Fprintf(w, "%s", encode(h[:]))
}

func sha224Handler(w http.ResponseWriter, r *http.Request) {
	h := sha256.Sum224([]byte(r.URL.Path[len(SHA224Url):]))
	fmt.Fprintf(w, "%s", encode(h[:]))
}

func sha256Handler(w http.ResponseWriter, r *http.Request) {
	h := sha256.Sum256([]byte(r.URL.Path[len(SHA256Url):]))
	fmt.Fprintf(w, "%s", encode(h[:]))
}

func sha384Handler(w http.ResponseWriter, r *http.Request) {
	h := sha512.Sum384([]byte(r.URL.Path[len(SHA384Url):]))
	fmt.Fprintf(w, "%s", encode(h[:]))
}

func sha512Handler(w http.ResponseWriter, r *http.Request) {
	h := sha512.Sum512([]byte(r.URL.Path[len(SHA512Url):]))
	fmt.Fprintf(w, "%s", encode(h[:]))
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	wordTohash := []byte(r.URL.Path[len(ALLUrl):])
	m5 := md5.Sum(wordTohash)
	s1 := sha1.Sum(wordTohash)
	s224 := sha256.Sum224(wordTohash)
	s256 := sha256.Sum256(wordTohash)
	s384 := sha512.Sum384(wordTohash)
	s512 := sha512.Sum512(wordTohash)
	fmt.Fprintf(w, "%s\n%s\n%s\n%s\n%s\n%s\n", encode(m5[:]), encode(s1[:]), encode(s224[:]), encode(s256[:]), encode(s384[:]), encode(s512[:]))
}

func main() {
	http.HandleFunc(MD5Url, md5Handler)
	http.HandleFunc(SHA1Url, sha1Handler)
	http.HandleFunc(SHA224Url, sha224Handler)
	http.HandleFunc(SHA256Url, sha256Handler)
	http.HandleFunc(SHA384Url, sha384Handler)
	http.HandleFunc(SHA512Url, sha512Handler)
	http.HandleFunc(ALLUrl, allHandler)

	log.Fatal(http.ListenAndServe(":8008", nil))
}
