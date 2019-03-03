package main

import (
	"html/template"
	"net/http"
        "io/ioutil"
        "fmt"
        "os"
)

var (
	templates = template.Must(template.New("").
		ParseGlob("templates/*.html"))
)

func (cw *conferenceWebsiteServer) homeHandler(w http.ResponseWriter, r *http.Request) {
    
    var confDetailsSvcAddr, scheduleSvcAddr string

    mustMapEnv(&confDetailsSvcAddr, "CONFERENCE_DETAILS_SERVICE_ADDR")
    mustMapEnv(&scheduleSvcAddr, "SCHEDULE_SERVICE_ADDR")
    
    response, err := http.Get("http://" + scheduleSvcAddr + "/schedule")
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            fmt.Println(string(data))
            if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
	    	    "conference_name":   "DevNexus", 
	    	    "schedule":   string(data), 
	    }); err != nil {
        }
    }
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
	    panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
