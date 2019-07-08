package test

import (
	"net/http"
	"fmt"
	"io"
	"log"
	"github.com/enroban/go-url-pattern/pattern"
)

func PatternTest1()  {
	h3 := func(params map[string]string, w http.ResponseWriter, r *http.Request) {
		fmt.Println("params:",params)
		abc:=r.URL.Query().Get("abc")
		fmt.Println("abc:",abc)
		io.WriteString(w, "abc is:"+abc)
	}

	pattern.PatternsFunctionContainer["/hello/{id}/world/{id2}"]=h3

	http.HandleFunc("/", pattern.UrlMatch)


	log.Fatal(http.ListenAndServe(":8081", nil))
}

func PatternTest2()  {
	fmt.Println(pattern.Matchpattern("/abcd/{id1}/abc/{id2}","/abcd/ad1/abc/ad23sd"))
}
