package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/alligator"
	"github.com/lindsaygelle/nook/character/alpaca"
	"github.com/lindsaygelle/nook/character/anteater"
	"github.com/lindsaygelle/nook/character/axolotl"
	"github.com/lindsaygelle/nook/character/bear"
	"github.com/lindsaygelle/nook/character/bearcub"
	"github.com/lindsaygelle/nook/character/beaver"
	"github.com/lindsaygelle/nook/character/bird"
	"github.com/lindsaygelle/nook/character/boar"
	"github.com/lindsaygelle/nook/character/bull"
	"github.com/lindsaygelle/nook/character/camel"
	"github.com/lindsaygelle/nook/character/cat"
	"github.com/lindsaygelle/nook/character/chameleon"
	"github.com/lindsaygelle/nook/character/chicken"
	"github.com/lindsaygelle/nook/character/cow"
	"github.com/lindsaygelle/nook/character/deer"
	"github.com/lindsaygelle/nook/character/dodo"
	"github.com/lindsaygelle/nook/character/dog"
	"github.com/lindsaygelle/nook/character/duck"
	"github.com/lindsaygelle/nook/character/eagle"
	"github.com/lindsaygelle/nook/character/elephant"
	"github.com/lindsaygelle/nook/character/fox"
	"github.com/lindsaygelle/nook/character/frillneckedlizard"
	"github.com/lindsaygelle/nook/character/frog"
	"github.com/lindsaygelle/nook/character/furseal"
	"github.com/lindsaygelle/nook/character/giraffe"
	"github.com/lindsaygelle/nook/character/goat"
	"github.com/lindsaygelle/nook/character/gorilla"
	"github.com/lindsaygelle/nook/character/gyroid"
	"github.com/lindsaygelle/nook/character/hamster"
	"github.com/lindsaygelle/nook/character/hedgehog"
	"github.com/lindsaygelle/nook/character/hippo"
	"github.com/lindsaygelle/nook/character/horse"
	"github.com/lindsaygelle/nook/character/kangaroo"
	"github.com/lindsaygelle/nook/character/koala"
	"github.com/lindsaygelle/nook/character/lion"
	"github.com/lindsaygelle/nook/character/mole"
	"github.com/lindsaygelle/nook/character/monkey"
	"github.com/lindsaygelle/nook/character/mouse"
	"github.com/lindsaygelle/nook/character/octopus"
	"github.com/lindsaygelle/nook/character/ostrich"
	"github.com/lindsaygelle/nook/character/otter"
	"github.com/lindsaygelle/nook/character/owl"
	"github.com/lindsaygelle/nook/character/panther"
	"github.com/lindsaygelle/nook/character/peacock"
	"github.com/lindsaygelle/nook/character/pelican"
	"github.com/lindsaygelle/nook/character/penguin"
	"github.com/lindsaygelle/nook/character/pig"
	"github.com/lindsaygelle/nook/character/pigeon"
	"github.com/lindsaygelle/nook/character/rabbit"
	"github.com/lindsaygelle/nook/character/raccoon"
	"github.com/lindsaygelle/nook/character/reindeer"
	"github.com/lindsaygelle/nook/character/rhinoceros"
	"github.com/lindsaygelle/nook/character/seagull"
	"github.com/lindsaygelle/nook/character/sheep"
	"github.com/lindsaygelle/nook/character/skunk"
	"github.com/lindsaygelle/nook/character/sloth"
	"github.com/lindsaygelle/nook/character/squirrel"
	"github.com/lindsaygelle/nook/character/tapir"
	"github.com/lindsaygelle/nook/character/tiger"
	"github.com/lindsaygelle/nook/character/tortoise"
	"github.com/lindsaygelle/nook/character/turkey"
	"github.com/lindsaygelle/nook/character/turtle"
	"github.com/lindsaygelle/nook/character/walrus"
	"github.com/lindsaygelle/nook/character/wolf"
	"github.com/lindsaygelle/w3g"
	"golang.org/x/text/language"
)

const (
	categoryResident = "Resident"
	categoryVillager = "Villager"
)

var (
	contentHeaders = http.Header{
		w3g.ContentLanguage: {language.AmericanEnglish.String()},
		w3g.ContentType:     {"application/json;charset=utf-8"}}
)

type characterLink struct {
	Animal    string `json:"animal"`
	Character string `json:"character"`
	Link      string `json:"link"`
	Special   bool   `json:"special"`
}

type character struct {
	Animal        string `json:"animal"`
	Category      string `json:"character"`
	Birthday      uint8  `json:"birthday"`
	BirthdayMonth uint8  `json:"birthday_month"`
	Gender        string `json:"gender"`
	Language      string `json:"language"`
	Name          string `json:"name"`
	Ok            bool   `json:"ok"`
	Special       bool   `json:"special"`
}

type resident struct {
	character
}

type villager struct {
	character

	Phrase string `json:"phrase"`
}

func getAnimalKey(v nook.Animal) string {
	return string(v.Key)
}

func getAnimalName(l language.Tag, v nook.Animal) (string, bool) {
	name, ok := v.Name.Get(l)
	return name.Value, ok
}

func getCharacterKey(v nook.Character) string {
	return string(v.Key)
}

func getCharacterLink(v nook.Character) string {
	return ("/" + strings.ToLower(getCharacterPath(v)))
}

func getCharacterName(l language.Tag, v nook.Character) (string, bool) {
	name, ok := v.Name.Get(l)
	return name.Value, ok
}

func getCharacterPath(v nook.Character) string {
	return strings.Join([]string{getAnimalKey(v.Animal), getCharacterKey(v)}, "/")
}

func getGenderName(l language.Tag, v nook.Gender) (string, bool) {
	name, ok := v.Name.Get(l)
	return name.Value, ok
}

func getVillagerPhrase(l language.Tag, v nook.Villager) (string, bool) {
	name, ok := v.Phrase.Get(l)
	return name.Value, ok
}

// handle adds the character links to the default Go HTTP server.
// It panics on the condition it cannot parse the slice of characterLinks.
func handle(v []characterLink) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, r, &b)
	})
}

// handleResident adds resident route to the default Go HTTP server.
// It panics on the condition it cannot parse the resident instance.
func handleResident(v nook.Resident) {
	pattern := getCharacterLink(v.Character)
	resident := newResident(language.AmericanEnglish, v)
	b, err := json.Marshal(resident)
	if err != nil {
		panic(err)
	}
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, r, &b)
	})
}

// handleVillager adds village route to the default Go HTTP server.
// It panics on the condition it cannot parse the villager instance.
func handleVillager(v nook.Villager) {
	pattern := getCharacterLink(v.Character)
	villager := newVillager(language.AmericanEnglish, v)
	b, err := json.Marshal(villager)
	if err != nil {
		panic(err)
	}
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, r, &b)
	})
}

// newCharacterLink creates a new characterLinke instance.
func newCharacterLink(v nook.Character) characterLink {
	return characterLink{
		Animal:    string(v.Animal.Key),
		Character: string(v.Key),
		Link:      getCharacterLink(v),
		Special:   v.Special}
}

// newCharacter returns a new character instance. It creates the character based on the language.Tag provided.
func newCharacter(l language.Tag, v nook.Character) character {
	animalValue, animalOK := getAnimalName(l, v.Animal)
	characterValue, characterOK := getCharacterName(l, v)
	genderValue, genderOK := getGenderName(l, v.Gender)
	return character{
		Animal:        animalValue,
		Birthday:      v.Birthday.Day,
		BirthdayMonth: uint8(v.Birthday.Month),
		Gender:        characterValue,
		Language:      l.String(),
		Name:          genderValue,
		Ok:            (animalOK && characterOK && genderOK),
		Special:       v.Special}
}

// newResident returns a new resident instance. It creates the resident based on the language.Tag provided.
func newResident(l language.Tag, v nook.Resident) resident {
	character := newCharacter(l, v.Character)
	resident := resident{character: character}
	resident.Category = categoryResident
	return resident
}

// newVillager returns a new villager instance. It creates the villager based on the language.Tag provided.
func newVillager(l language.Tag, v nook.Villager) villager {
	character := newCharacter(l, v.Character)
	villager := villager{character: character}
	phraseValue, ok := getVillagerPhrase(l, v)
	villager.Category = categoryVillager
	villager.Ok = (character.Ok && ok)
	villager.Phrase = phraseValue
	return villager
}

// process processes a collection of interfaces ant attempts to build out a root path for the HTTP server.
func process(v ...interface{}) {
	characterLinks := make([]characterLink, 0)
	for _, v := range v {
		characterLink := processContent(v)
		characterLinks = append(characterLinks, characterLink...)
	}
	handle(characterLinks)
}

// processContent processes the interface type given to the function. It inspects the type and determins whether
// the current argument is a instance of nook.Residents or nook.Villagers. On unknown type it panics.
func processContent(v interface{}) []characterLink {
	switch t := v.(type) {
	case nook.Residents:
		return processResidents(t)
	case nook.Villagers:
		return processVillagers(t)
	default:
		panic(t)
	}
}

// processResident processes an individual nook.Resident.
func processResident(v nook.Resident) {
	handleResident(v)
}

// processResidnets process all nook.Residents.
func processResidents(v nook.Residents) []characterLink {
	characterLinks := make([]characterLink, 0)
	v.Each(func(k nook.Key, v nook.Resident) {
		characterLinks = append(characterLinks, newCharacterLink(v.Character))
		processResident(v)
	})
	return characterLinks
}

// processVillager processes an indivdual nook.Villager.
func processVillager(v nook.Villager) {
	handleVillager(v)
}

// processVillagers processes all nook.Villagers.
func processVillagers(v nook.Villagers) []characterLink {
	characterLinks := make([]characterLink, 0)
	v.Each(func(k nook.Key, v nook.Villager) {
		characterLinks = append(characterLinks, newCharacterLink(v.Character))
		processVillager(v)
	})
	return characterLinks
}

// writeResponse writes the common headers to http.ResponseWriter.
func writeResponse(w http.ResponseWriter, r *http.Request, b *[]byte) {
	for key, values := range contentHeaders {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(*b)
}

func main() {
	process(
		alligator.Villagers,
		alpaca.Residents,
		anteater.Villagers,
		axolotl.Residents,
		bear.Villagers,
		bearcub.Villagers,
		beaver.Residents,
		bird.Residents,
		bird.Villagers,
		boar.Residents,
		bull.Villagers,
		camel.Residents,
		cat.Residents,
		cat.Villagers,
		chameleon.Residents,
		chicken.Villagers,
		cow.Villagers,
		deer.Villagers,
		dodo.Residents,
		dog.Residents,
		dog.Villagers,
		duck.Villagers,
		eagle.Villagers,
		elephant.Villagers,
		fox.Residents,
		frillneckedlizard.Residents,
		frog.Villagers,
		furseal.Residents,
		giraffe.Residents,
		goat.Villagers,
		gorilla.Villagers,
		gyroid.Residents,
		hamster.Villagers,
		hedgehog.Residents,
		hippo.Villagers,
		horse.Villagers,
		kangaroo.Villagers,
		koala.Villagers,
		lion.Villagers,
		mole.Residents,
		monkey.Residents,
		monkey.Villagers,
		mouse.Villagers,
		octopus.Villagers,
		ostrich.Villagers,
		otter.Residents,
		owl.Residents,
		panther.Residents,
		peacock.Residents,
		pelican.Residents,
		penguin.Villagers,
		pig.Villagers,
		pigeon.Residents,
		rabbit.Residents,
		rabbit.Villagers,
		raccoon.Residents,
		reindeer.Residents,
		rhinoceros.Villagers,
		seagull.Residents,
		sheep.Villagers,
		skunk.Residents,
		sloth.Residents,
		squirrel.Residents,
		squirrel.Villagers,
		tapir.Residents,
		tiger.Villagers,
		tortoise.Residents,
		turkey.Residents,
		turtle.Residents,
		walrus.Residents,
		wolf.Villagers)

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}
