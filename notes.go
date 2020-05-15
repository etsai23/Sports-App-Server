
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

var Sports [18]string {
	Sports[0] = "Baseball"
	Sports[1] = "Boys Lacrosse"
	Sports[2] = "Girls Lacrosse"
	Sports[3] = "Boys Tennis"
	Sports[4] = "Girls Tennis"
	Sports[5] = "Boys Volleyball"
	Sports[6] = "Girls Volleyball"
	Sports[7] = "Swimming"
	Sports[8] = "Track kand Field"
	Sports[9] = "Boys Waterpolo"
	Sports[10] = "Girls Waterpolo"
	Sports[11] = "Football"
	Sports[12] = "Cross Country"
	Sports[13] = "Boys Basketball"
	Sports[14] = "Girls Basketball"
	Sports[15] = "Boys Soccer"
	Sports[16] = "Girls Soccer"
	Sports[17] = "Squash"
	Sports[18] = "Ultimate Frisbee" 
}

//the struct below might have to be called selection
type Event struct {
	SportsName     string
	Lat        string
	Long      string
	//ask if location is a string, if it is a string we need to refer to the coordinates 
	//that already fall under                                                                                  the name of the location
	Time       string
	Notes      string

}

var game = []Sports{}
var eventData = []Event{}
var sportJson = []byte{}



func handlerBySport(w http.ResponseWriter, r *http.Request) {
	// info
	fmt.Println("Handler for Sport - Incoming Request: ")
	fmt.Println("Method: ", r.Method, " ", r.URL)

	fmt.Println("Sports", game)

	SportsName := strings.Split(r.URL.Path, "/")[2]
	//var error: error
	SportsName, _ = url.PathUnescape(game)

	var selectedSport = ""
	//var selectedSportLat = ""
	//var selectedSportLong = ""
	//var selectedSportTime = ""
	//var selectedSportNotes = ""



	for index, _ := range game {
		if game[index].Sports == SportsName {
			selectedSport = game[index].Sports
			break
		}
	}
		//seperate sport name from URL - ex: "Boys Basketball"
	//search through array "game" and find struct with with specific sport name
	
	
	var currentSportData = []string{}

	for index, _ := range game {
		if game[index].Sports== selectedSport {
			currentSportData = append(currentSportsData, eventData[index].SportsName) //append name of sport of game
			currentSportData = append(currentSportsData, eventData[index].Lat) //append Latittude of game
			currentSportData = append(currentSportsData, eventData[index].Long)// append Long. of game)
			currentSportData = append(currentSportsData, eventData[index].Time)//append Time of game)
			currentSportData = append(currentSportsData, eventData[index].Time)//append Notes of game)
		} 
		break
	}
	//For above, we need to find a way to store the lat, long, time and notes of a specfic sports event
	//into the currentSportData array. I think I did it correctly, but that's given that there's preexisting
	//data. I don't know how or where to create this/if we should do it manually



	//append each firstname and lastname into a new array, and print that using the bottom line of code
	sportJson, _ = json.Marshal(currentSportData)

	dataJson := "{\"Table\": \"" + selectedSport + "\", \"Names\": " + string(sportJson) + "}"

	// Answer the Client request
	fmt.Fprintf(w, dataJson)
}

func main {

	//need to code the logistics of each sports game here. 


	eventData = append(eventData, Event{
		SportsName: 
		Lat:
		Long:
		Time:
		Notes: 
	})
	//we gotta figure out how each of the variables under the Event struct are going to be defined. 
}

