package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func rebuildbase64(w http.ResponseWriter, r *http.Request) {

	//log about request
	log.Println("method:", r.Method)
	log.Printf("%v\n", r.URL)
	log.Printf("%v\n", r.RemoteAddr)
	log.Printf("%v\n", r.Host)
	log.Printf("%v\n", r.Header)

	//m max 5 MB file name we can change ut
	r.ParseMultipartForm(5 << 20)

	//myfileparam is the name of file in post request body

	log.Printf("%v\n", r.Header.Get("Content-Type"))

	base64enc := r.Body

	defer base64enc.Close()

	cont, err := ioutil.ReadAll(base64enc)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, "malformed request", http.StatusBadRequest)

		return
	}

	// there some err variable shadowing
	//var buf []byte
	var mp map[string]json.RawMessage

	err = json.Unmarshal(cont, &mp)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, "malformed json request", http.StatusBadRequest)

		return
	}

	var str string
	err = json.Unmarshal(mp["Base64"], &str)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, "malformed json request", http.StatusBadRequest)

		return
	}

	buf, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, "malformed base64 encoding", http.StatusBadRequest)
		return

	}

	log.Printf("%v\n", r.Header.Get("Content-Type"))

	log.Printf("%v\n", http.DetectContentType(buf))

	addgwheader(w, temp)

	s, e := w.Write(buf)
	if e != nil {
		log.Println(e)
		return
	}
	log.Println(s)

	// so  here we can use either open file or  ioutil.write file
	/*
		fmt.Fprintf(w, "%v\n", handler.Header)
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	*/

}
