package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var seatingJson = []byte{}

//the struct below might have to be called selection
type Event struct {
	Sport     string
	Lat        string
	Long      string
	//ask if location is a string, if it is a string we need to refer to the coordinates 
	//that already fall under the name of the location
	Time       string
	Notes      string
	Location string

}
var Output []Event


// the goal of the main function is to generate a list of "Events" which 

func main() {
	sportsList:= [...]string{"Boy's Waterpolo", "Girl's Waterpolo", "Football", "Baseball", "Swimming", "Boy's Basketball", "Girl's Basketball", "Boy's Tennis", "Girl's Tennis","Ultimate Frisbee","Girl's Soccer", "Boy's Soccer", "Girl's Squash", "Boy's Squash", "Girl's Volleyball", "Boy's Volleyball","Cross Country", "Track"}
	locations:= [...]string{"Old Gym","New Gym","Pool","Upper Field","Lower Field","Tennis Courts"}
	var latitude = string 
	var longitude = string 
	var time = int
	var sport string
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	time = (rand.Intn(max-min) + min)
	
	for count := 0: count < 4 {
		min := 0
		max := 4
		var locationPicker = [(rand.Intn(max - min) + min)]
	
		if location[locationPicker] == "New Gym" {
			latitude = "34.40540045"
			longitude = "-119.47687865"
		}
		if location[locationPicker] == "Old Gym" {
			latitude = "34.40541902"
			longitude = "-119.47892551"
		}
		if location[locationPicker] == "Pool" {
			latitude = "34.40656022"
			longitude = "-119.47836095"
		}
		if location[locationPicker] == "Upper Field" {
		latitude = "34.40620172"
		longitude = "-119.47591854"
		}
		if location[locationPicker] == "Lower Field" {
			latitude = "34.40649877"
			longitude = "-119.47729934"
		}
		if location[locationPicker] == "Tennis Courts" {
			latitude = "34.40659497"
			longitude = "-119.47889747"
		}
		max := 14
		var sport = (rand.Intn(max-min) + min)
		var notes = "Notes"
		Output[count].Sport = sport
		Output[count].Lat = latitude
		Output[count].Long = longitude
		Output[count].Time = time
		Output[count].Notes = notes
	}
	fmt.Println(Output)
	
}
func handler(w http.ResponseWriter, r *http.Request) {

    // Let's print the info
    fmt.Println("Incoming Request: ")
    fmt.Println("Method: ", r.Method, " ", r.URL)

    // Collect all the header keys
    headerKeys := make([]string, len(r.Header))
    i := 0
    for k := range r.Header {
        headerKeys[i] = k
        i++
    }
    // Show Client Headers
    for _, line := range headerKeys {
        fmt.Println("  > ", line, ":", r.Header.Get(line))
    }

	// Answer the Client request
	eventjson, _ = json.Marshal(Output)
	eventjson := Output
//	jsonData, err = json.Marshal(//this probably has to be whatever function we need to change into JSON)
	fmt.Fprintf(w, string(eventjson))
	
	// MacFarlane's basic emtpy handler function
}

//strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}



