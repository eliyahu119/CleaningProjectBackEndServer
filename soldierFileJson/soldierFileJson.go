package soldierFileJson

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	Soldiers        = "soldiers"
	Commanders      = "Commanders"
	Days            = "days"
	Date            = "date"
	SoldierInLine   = "soldierInLine"
	CommanderInLine = "commanderInLine"
	MaxBorderSize   = "maxBorderSize"
	Max_range       = "max_range"
	WeekendDays     = "weekendDays"
	NumberOfBlocks  = "numberOfBlocks"
)

//Soldeir_file the Struct for the json init
type Soldeir_file struct {
	Soldiers        []string `json:"soldiers"`
	Commanders      []string `json:"Commanders"`
	Days            []string `json:"days"`
	Date            string   `json:"date"`
	SoldierInLine   int      `json:"soldierInLine"`
	CommanderInLine int      `json:"commanderInLine"`
	MaxBorderSize   int      `json:"maxBorderSize"`
	Max_range       int      `json:"max_range"`
	WeekendDays     int      `json:"weekendDays"`
	NumberOfBlocks  int      `json:"numberOfBlocks"`
}

//CreateTheJSON creates the JsonStruct for the  json file initializaion of the backEnd
func CreateTheJSON(Form url.Values) (Soldeir_file, error) {
	soldierInLineJS, err := strconv.Atoi(Form[SoldierInLine][0])
	if err != nil {
		return Soldeir_file{}, err
	}
	commanderInLineJS, err := strconv.Atoi(Form[CommanderInLine][0])
	if err != nil {
		return Soldeir_file{}, err
	}
	maxBorderSizeJS, err := strconv.Atoi(Form[MaxBorderSize][0])
	if err != nil {
		return Soldeir_file{}, err
	}
	max_rangeJS, err := strconv.Atoi(Form[Max_range][0])
	if err != nil {
		return Soldeir_file{}, err
	}
	weekendDaysJS, err := strconv.Atoi(Form[WeekendDays][0])
	if err != nil {
		return Soldeir_file{}, err
	}
	numberOfBlocksJS, err := strconv.Atoi(Form[NumberOfBlocks][0])
	if err != nil {
		return Soldeir_file{}, err
	}

	jsonFile := Soldeir_file{
		Soldiers:        strings.Split(Form[Soldiers][0], ","),
		Commanders:      strings.Split(Form[Commanders][0], ","),
		Days:            []string{"ראשון", "שני", "שלישי", "רביעי", "חמישי", "שישי", "שבת"},
		Date:            Form[Date][0],
		SoldierInLine:   soldierInLineJS,
		CommanderInLine: commanderInLineJS,
		MaxBorderSize:   maxBorderSizeJS,
		Max_range:       max_rangeJS,
		WeekendDays:     weekendDaysJS,
		NumberOfBlocks:  numberOfBlocksJS,
	}
	return jsonFile, nil
}
