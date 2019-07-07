package pattern

import (
	"strings"
	"regexp"
	"net/http"
	"log"
)

//最大匹配深度，默认10
var MaxPatternLength = 10

//匹配模板http处理函数的容器，匹配成功以后就会自动调用注册的函数进行处理
var PatternsFunctionContainer = make(map[string]func(params map[string]string, w http.ResponseWriter, r *http.Request))

//uri处理，将类似于/a/b/{id1}/c/{id2}/d的传入之后，处理出来获取到id1,id2
func GetUrlParamsByPattern(pattern string, uri string) (map[string]string) {
	params := make(map[string]string)
	uriSplited := strings.Split(uri, "/")
	patternSplited := strings.Split(pattern, "/")
	r, _ := regexp.Compile("^\\{[0-9a-zA-Z]+\\}$")
	r2, _ := regexp.Compile("\\{|\\}")
	for k, v := range patternSplited {
		if r.MatchString(v) {
			keys := r2.Split(v, 3)
			params[keys[1]] = uriSplited[k]
		}
	}
	return params
}

//判断模板与路径是否匹配
func matchpattern(pattern string, uri string) bool {
	r, _ := regexp.Compile("\\{[0-9a-zA-Z]+\\}")
	patternsplited := r.Split(pattern, MaxPatternLength)
	patternsplitedsize := len(patternsplited)
	reg := ""
	reg += "^"
	for x, v := range patternsplited {
		reg += v
		if x+1 != patternsplitedsize {
			//默认不支持中文模板参数，若要允许则使用这个
			//reg += "[0-9a-zA-Z-_\u4e00-\u9fa5]+"
			reg += "[0-9a-zA-Z-_]+"
		}
	}
	reg += "$"
	r2, _ := regexp.Compile(reg)
	return r2.MatchString(uri)
}

//url匹配处理返回url匹配以后模板代表的参数，还有函数用以后期处理
func UrlMatch(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	for k, v := range PatternsFunctionContainer {
		if matchpattern(k, url) {
			params := GetUrlParamsByPattern(k, url)
			v(params,w,r)
			return
		}
	}
	log.Println("Pattern not found")
}
