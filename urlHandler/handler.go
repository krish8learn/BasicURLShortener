package urlHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

/*Handler for home page*/
func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home Page")
	}
}

/*
we will perform operation by map
*/
var PathURLs = map[string]string{
	"/krishgit": "https://github.com/krish8learn",
	"/urlRepo":  "https://github.com/krish8learn/BasicURLShortener",
}

/*
this will return http.HandlerFunc which also implements http.Handler
*/
func MapHandler(mapUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if originalUrl, cond := mapUrls[r.URL.Path]; cond {
			//given url path match with map key, proper url return
			http.Redirect(w, r, originalUrl, http.StatusFound)
		}
		//given url does not match, http.Handler call , get back to home
		fallback.ServeHTTP(w, r)
	}
}

//struct for yaml and json
type Mapper struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

/*
we will perform by YAML
*/
var YamlUrls = `
- path: /krishgit
  url: https://github.com/krish8learn
- path: /urlRepo
  url: https://github.com/krish8learn/BasicURLShortener
`

func YamlHandler(fallback http.Handler) (http.HandlerFunc, error) {
	// parse the yaml data into array
	var array []Mapper
	err := yaml.Unmarshal([]byte(YamlUrls), &array)
	if err != nil {
		return nil, err
	}
	//convert yaml struct to map
	mapYaml := make(map[string]string)
	for _, value := range array {
		mapYaml[value.Path] = value.Url
	}
	//use mapHandler
	return MapHandler(mapYaml, fallback), nil
}

/*
we will perform by json
*/

var jsonURL = `[
	{
		"path":  "/krishgit",
		"url": "https://github.com/krish8learn"	
	},
	{
		"path":  "/urlRepo",
		"url": "https://github.com/krish8learn/BasicURLShortener"	
	}
]`

func JsonHandler(fallback http.Handler) (http.HandlerFunc, error) {
	// parse the yaml data into array
	var array []Mapper
	err := json.Unmarshal([]byte(jsonURL), &array)
	if err != nil {
		return nil, err
	}
	//convert yaml struct to map
	mapJson := make(map[string]string)
	for _, value := range array {
		mapJson[value.Path] = value.Url
	}
	//use mapHandler
	return MapHandler(mapJson, fallback), nil
}
