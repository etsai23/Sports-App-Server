
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


Sport := [...]string{
 "Baseball",
 "Boys Lacrosse",
 "Girls Lacrosse",
 "Boys Tennis",
 "Girls Tennis",
 "Boys Volleyball",
 "Girls Volleyball",
 "Swimming",
 "Track and Field",
 "Boys Waterpolo",
 "Girls Waterpolo",
 "Football",
 "Cross Country",
 "Boys Basketball",
 "Girls Basketball",
 "Boys Soccer",
 "Girls Soccer",
 "Squash",
 "Ultimate Frisbee"

}


//the struct below might have to be called selection
type Event struct {
	Sport     string
	Lat        string
	Long      string
	//ask if location is a string, if it is a string we need to refer to the coordinates 
	//that already fall under                                                                                  the name of the location
	Time       string
	Notes      string

}







jsonData, err = json.Marshal(//this probably has to be whatever function we need to change into JSON)

