package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func rebuildfile(w http.ResponseWriter, r *http.Request) {

	//log about request
	log.Println("method:", r.Method)
	log.Printf("%v\n", r.URL)
	log.Printf("%v\n", r.RemoteAddr)
	log.Printf("%v\n", r.Host)
	log.Printf("%v\n", r.Header)

	//m max 5 MB file name we can change ut
	r.ParseMultipartForm(5 << 20)

	//myfileparam is the name of file in post request body
	file, handler, err := r.FormFile("file")

	if err != nil {
		log.Println(err)
		return
	}

	//this only to parse post form to extract data for log
	if errp := r.ParseForm(); errp != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		log.Printf("Form[%q] = %q\n", k, v)
	}

	log.Printf("%v\n", handler.Filename)
	log.Printf("%v\n", handler.Size)

	defer file.Close()

	buf, er := ioutil.ReadAll(file)
	if er != nil {
		log.Println(er)
		return
	}
	log.Printf("%v\n", handler.Header.Get("Content-Type"))
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
