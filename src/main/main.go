package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jmoiron/jsonq"
)

type Type struct {
	// Name of the type
	Name string `json:"name"`
	// The effective basedata, damage multiplize 2x
	EffectiveAgainst []string `json:"effectiveAgainst"`
	// The weak basedata that against, damage multiplize 0.5x
	WeakAgainst []string `json:"weakAgainst"`
}

type Pokemon struct {
	Number         string   `json:"Number"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Weaknesses     []string `json:"Weaknesses"`
	FastAttackS    []string `json:"Fast Attack(s)"`
	Weight         string   `json:"Weight"`
	Height         string   `json:"Height"`
	Candy          struct {
		Name     string `json:"Name"`
		FamilyID int    `json:"FamilyID"`
	} `json:"Candy"`
	NextEvolutionRequirements struct {
		Amount int    `json:"Amount"`
		Family int    `json:"Family"`
		Name   string `json:"Name"`
	} `json:"Next Evolution Requirements,omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	SpecialAttacks      []string `json:"Special Attack(s)"`
	BaseAttack          int      `json:"BaseAttack"`
	BaseDefense         int      `json:"BaseDefense"`
	BaseStamina         int      `json:"BaseStamina"`
	CaptureRate         float64  `json:"CaptureRate"`
	FleeRate            float64  `json:"FleeRate"`
	BuddyDistanceNeeded int      `json:"BuddyDistanceNeeded"`
}

// Move is an attack information. The
type Move struct {
	// The ID of the move
	ID int `json:"id"`
	// Name of the attack
	Name string `json:"name"`
	// Type of attack
	Type string `json:"type"`
	// The damage that enemy will take
	Damage int `json:"damage"`
	// Energy requirement of the attack
	Energy int `json:"energy"`
	// Dps is Damage Per Second
	Dps float64 `json:"dps"`
	// The duration
	Duration int `json:"duration"`
}

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}
type BaseDataArr []BaseData

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/list url:", r.URL)
	fmt.Fprint(w, "The List Handler\n")
	fmt.Fprintln(w, r)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/get url:", r.URL)
	fmt.Fprint(w, "The Get Handler\n")
}

func typeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/types url:", r.URL)
	fmt.Fprint(w, "All of the pokemon types\n")
	m := getData()

	n := m["types"].([]interface{})
	//types := make([]*Type, len(n))

	for i := range n {

		name := n[i].(map[string]interface{})["name"].(string)
		ea := n[i].(map[string]interface{})["effectiveAgainst"].([]interface{})
		wa := n[i].(map[string]interface{})["weakAgainst"].([]interface{})
		fmt.Fprintln(w, "\nName:", name)
		fmt.Fprintln(w, "Effective Against:", ea)
		fmt.Fprintln(w, "Weak Against:", wa)

		/*//types[i] = &Type{name, ea, wa}
		//fmt.Fprintln(w, moves[i])
		//fmt.Fprintln(w, "\nName:", types[i].Name)
		fmt.Fprintln(w, "\nName:", name)
		//fmt.Fprintln(w, "Effective Against:", types[i].EffectiveAgainst)
		fmt.Fprintln(w, "Effective Against:", ea)
		//fmt.Fprintln(w, "Weak Against:", types[i].WeakAgainst)
		fmt.Fprintln(w, "Weak Against:", wa)*/

	}
}
func returnSingleType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["type"]
	m := getData()

	n := m["types"].([]interface{})
	//types := make([]*Type, len(n))

	for i := range n {

		name := n[i].(map[string]interface{})["name"].(string)
		ea := n[i].(map[string]interface{})["effectiveAgainst"].([]interface{})
		wa := n[i].(map[string]interface{})["weakAgainst"].([]interface{})
		if name == key {
			fmt.Fprintln(w, "\nName:", name)
			fmt.Fprintln(w, "Effective Against:", ea)
			fmt.Fprintln(w, "Weak Against:", wa)
		}
	}
	fmt.Println("Key: " + key)
}
func pokemonHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/pokemons url:", r.URL)
	fmt.Fprint(w, "All pokemons\n")
	m := getData()

	n := m["pokemons"].([]interface{})

	//pokemons := make([]*Pokemon, len(n))

	for i := range n {
		number := n[i].(map[string]interface{})["Number"].(string)
		name := n[i].(map[string]interface{})["Name"].(string)
		classification := n[i].(map[string]interface{})["Classification"].(string)
		type1 := n[i].(map[string]interface{})["Type I"].([]interface{})
		type2 := n[i].(map[string]interface{})["Type II"]
		weaknesses := n[i].(map[string]interface{})["Weaknesses"].([]interface{})
		fastattacks := n[i].(map[string]interface{})["Fast Attack(s)"].([]interface{})
		weight := n[i].(map[string]interface{})["Weight"].(string)
		height := n[i].(map[string]interface{})["Height"].(string)

		//pokemons[i] = &Pokemon{number, name, classification, type1, type2, weaknesses, fastattacks, weight, height}
		//fmt.Fprintln(w, pokemons[i])
		//fmt.Fprintln(w, "\nName:", types[i].Name)
		fmt.Fprintln(w, "\nName:", name)
		fmt.Fprintln(w, "Number:", number)
		fmt.Fprintln(w, "Classification:", classification)
		fmt.Fprintln(w, "Type I:", type1)
		fmt.Fprintln(w, "Type II:", type2)
		fmt.Fprintln(w, "Weaknesses:", weaknesses)
		fmt.Fprintln(w, "Fast Attack(s):", fastattacks)
		fmt.Fprintln(w, "Weight:", weight)
		fmt.Fprintln(w, "Height:", height)

	}

}
func moveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/moves url:", r.URL)
	fmt.Fprint(w, "All moves\n")
	m := getData()

	n := m["moves"].([]interface{})
	moves := make([]*Move, len(n))

	for i := range n {
		id := n[i].(map[string]interface{})["id"].(float64)
		name := n[i].(map[string]interface{})["name"].(string)
		tip := n[i].(map[string]interface{})["type"].(string)
		damage := n[i].(map[string]interface{})["damage"].(float64)
		energy := n[i].(map[string]interface{})["energy"].(float64)
		dps := n[i].(map[string]interface{})["dps"].(float64)
		duration := n[i].(map[string]interface{})["duration"].(float64)

		moves[i] = &Move{int(id), name, tip, int(damage), int(energy), dps, int(duration)}
		//fmt.Fprintln(w, moves[i])
		fmt.Fprintln(w, "\nName:", moves[i].Name)
		fmt.Fprintln(w, "ID:", moves[i].ID)
		fmt.Fprintln(w, "Type:", moves[i].Type)
		fmt.Fprintln(w, "Damage:", moves[i].Damage)
		fmt.Fprintln(w, "Energy:", moves[i].Energy)
		fmt.Fprintln(w, "Dps:", moves[i].Dps)
		fmt.Fprintln(w, "Duration:", moves[i].Duration)

	}
}
func otherwise(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Pokedex\n")

	/*b, m := readData()
	fmt.Fprintln(w, b.Pokemons[0].Name)

	fmt.Println(len(m))
	for k, v := range m {
		x := b.Pokemons[0].Name
		fmt.Fprintln(w, x)
		fmt.Println("\nk:", k)
		fmt.Println("v:", v)
	}*/

}
func readData() (BaseData, map[string]interface{}) {
	log.Println("getData called")
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Print("Error:", err)
	}
	basedata := BaseData{}
	m := make(map[string]interface{})
	err = json.Unmarshal([]byte(content), &basedata)
	err = json.Unmarshal([]byte(content), &m)
	if err != nil {
		fmt.Print("Error:", err)
	}
	return basedata, m
}

func getData() map[string]interface{} {
	log.Println("getData called")
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Print("Error:", err)
	}
	//var basedata BaseData
	m := make(map[string]interface{})
	//err = json.Unmarshal([]byte(content), &basedata)
	err = json.Unmarshal([]byte(content), &m)
	if err != nil {
		fmt.Print("Error:", err)
	}
	return m
}
func decodeData() *jsonq.JsonQuery {
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Print("Error:", err)
	}
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(content)))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	jq.String("types")
	return jq

}

func parseMap(aMap map[string]interface{}, w http.ResponseWriter, r *http.Request) {

	for key, value := range aMap {

		switch concreteVal := value.(type) {
		case map[string]interface{}:
			fmt.Fprintln(w, key)
			parseMap(value.(map[string]interface{}), w, r)
		case []interface{}:
			fmt.Fprintln(w, key)
			parseArray(value.([]interface{}), w, r)
		default:
			fmt.Fprintln(w, key, ":", concreteVal)

		}
	}

}

func parseArray(anArray []interface{}, w http.ResponseWriter, r *http.Request) {

	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			//fmt.Println("Index:", i)

			parseMap(val.(map[string]interface{}), w, r)
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}), w, r)
		default:
			fmt.Fprintln(w, "-", concreteVal)
			//fmt.Fprintln(w, "Index", i, ":", concreteVal)

		}
	}
}

func main() {
	//TODO: read data.json to a BaseData
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/list", listHandler)
	myRouter.HandleFunc("/get", getHandler)
	myRouter.HandleFunc("/types", typeHandler)
	myRouter.HandleFunc("/pokemons", pokemonHandler)
	myRouter.HandleFunc("/moves", moveHandler)
	myRouter.HandleFunc("/types/{type}", returnSingleType)
	//TODO: add more
	myRouter.HandleFunc("/", otherwise)
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
