/*
Simple CLI to keep up to date with events in Formula 1
------------------------------------------------------
API: http://ergast.com/mrd/
Author: Ustym Hentosh
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	xj "github.com/basgys/goxml2json"
	"github.com/fatih/color"
)

func get_structed_info_from_api(url string) *bytes.Buffer {
	// sending and recieving Get request
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// reading body of the request
	bodyBytes_ds, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	// body -> json -> struct
	dictionary, err := xj.Convert(strings.NewReader(string(bodyBytes_ds)))
	if err != nil {
		panic(err)
	}
	return dictionary
}

func main() {

	if len(os.Args) != 2 {
		// ------------------ Geting Drivers Standing -----------------------------
		type DriverStandings struct {
			Root struct {
				StandingsTable struct {
					StandingsList struct {
						DriverStanding []struct {
							Position string `json:"-positionText"`
							Points   string `json:"-points"`
							Wins     string `json:"-wins"`
							Driver   struct {
								Name string `json:"FamilyName"`
							} `json:"Driver"`
							Team struct {
								Name string `json:"Name"`
							} `json:"Constructor"`
						} `json:"DriverStanding"`
					} `json:"StandingsList"`
				} `json:"StandingsTable"`
			} `json:"MRData"`
		}

		var dictionary_ds = get_structed_info_from_api("http://ergast.com/api/f1/current/driverStandings")
		var driver_standings DriverStandings
		sth := json.Unmarshal(dictionary_ds.Bytes(), &driver_standings)
		if sth != nil {
			fmt.Println("Error:", sth)
		}

		drivers_top := driver_standings.Root.StandingsTable.StandingsList.DriverStanding

		// Printing top 3 drivers
		color.Green("Drivers standings")
		for i := 0; i < 5; i++ {
			fmt.Printf(
				"%s. %s %s pt %s wins\n",
				drivers_top[i].Position,
				drivers_top[i].Driver.Name,
				drivers_top[i].Points,
				drivers_top[i].Wins,
			)
		}
		fmt.Println()

		// ------------------ Geting Teams Standing -----------------------------
		type TeamsStandings struct {
			Root struct {
				StandingsTable struct {
					StandingsList struct {
						TeamsStanding []struct {
							Team struct {
								Name        string `json:"Name"`
								Nationality string `json:"Nationality"`
							} `json:"Constructor"`
							Points   string `json:"-points"`
							Wins     string `json:"-wins"`
							Position string `json:"-positionText"`
						} `json:"ConstructorStanding"`
					} `json:"StandingsList"`
				} `json:"StandingsTable"`
			} `json:"MRData"`
		}

		dictionary_ts := get_structed_info_from_api("http://ergast.com/api/f1/current/constructorStandings")
		var teams_standings TeamsStandings
		sth_1 := json.Unmarshal(dictionary_ts.Bytes(), &teams_standings)
		if sth_1 != nil {
			fmt.Println("Error:", sth_1)
		}

		teams_top := teams_standings.Root.StandingsTable.StandingsList.TeamsStanding

		// Printing top 3 teams
		color.HiBlue("Constructors standings")
		for i := 0; i < 5; i++ {
			fmt.Printf(
				"%s. %s %s pt %s wins\n",
				teams_top[i].Position,
				teams_top[i].Team.Name,
				teams_top[i].Points,
				teams_top[i].Wins,
			)
		}
		fmt.Println()

		// ------------------------ Geting Recent Race -----------------------------
		type RecentRace struct {
			Root struct {
				RaceTable struct {
					Race struct {
						Round       string `json:"-round"`
						Name        string `json:"RaceName"`
						Date        string `json:"Date"`
						ResultsList struct {
							Result [3]struct {
								Position string `json:"-positionText"`
								Driver   struct {
									Name string `json:"FamilyName"`
								} `json:"Driver"`
							} `json:"Result"`
						} `json:"ResultsList"`
					} `json:"Race"`
				} `json:"RaceTable"`
			} `json:"MRData"`
		}

		dictionary_rr := get_structed_info_from_api("http://ergast.com/api/f1/current/last/results")
		var recent_race RecentRace
		sth_2 := json.Unmarshal(dictionary_rr.Bytes(), &recent_race)
		if sth_2 != nil {
			fmt.Println("Error:", sth_2)
		}
		last_race := recent_race.Root.RaceTable.Race

		// Last race info
		color.HiYellow("Recent Race")
		fmt.Printf(
			"%s %s\n",
			last_race.Date,
			last_race.Name,
		)
		for i := 0; i < 3; i++ {
			fmt.Printf(
				"%s. %s\n",
				last_race.ResultsList.Result[i].Position,
				last_race.ResultsList.Result[i].Driver.Name,
			)
		}
		fmt.Println()

		// ------------------------ Geting Schedule -----------------------------

		type Schedule struct {
			Root struct {
				RaceTable struct {
					Race []struct {
						Round string `json:"-round"`
						Name  string `json:"RaceName"`
						Date  string `json:"Date"`
					} `json:"Race"`
				} `json:"RaceTable"`
			} `json:"MRData"`
		}

		dictionary_sd := get_structed_info_from_api("http://ergast.com/api/f1/current")
		var schedule Schedule
		sth_3 := json.Unmarshal(dictionary_sd.Bytes(), &schedule)
		if sth_3 != nil {
			fmt.Println("Error:", sth_3)
		}

		intrace, err := strconv.Atoi(last_race.Round)
		if err != nil {
			panic(err)
		}
		next_race := schedule.Root.RaceTable.Race[intrace]

		// Next Race info
		color.HiMagenta("Next Race")
		fmt.Printf(
			"%s %s\n",
			next_race.Date,
			next_race.Name,
		)

		// Second arg
	} else {
		switch os.Args[1] {
		case "history":
			type History struct {
				Root struct {
					StandingsTable struct {
						StandingsList []struct {
							Season         string `json:"-season"`
							DriverStanding struct {
								Points string `json:"-points"`
								Wins   string `json:"-wins"`
								Driver struct {
									Name        string `json:"FamilyName"`
									Nationality string `json:"Nationality"`
								} `json:"Driver"`
								Constructor struct {
									Name        string `json:"Name"`
									Nationality string `json:"Nationality"`
								} `json:"Constructor"`
							} `json:"DriverStanding"`
						} `json:"StandingsList"`
					} `json:"StandingsTable"`
				} `json:"MRData"`
			}

			dictionary_h1 := get_structed_info_from_api("http://ergast.com/api/f1/driverStandings/1")
			dictionary_h2 := get_structed_info_from_api("http://ergast.com/api/f1/driverstandings/1?limit=30&offset=30")
			dictionary_h3 := get_structed_info_from_api("http://ergast.com/api/f1/driverstandings/1?limit=30&offset=60")
			var history_1 History
			var history_2 History
			var history_3 History
			json.Unmarshal(dictionary_h1.Bytes(), &history_1)
			json.Unmarshal(dictionary_h2.Bytes(), &history_2)
			json.Unmarshal(dictionary_h3.Bytes(), &history_3)

			all_history_1 := history_1.Root.StandingsTable.StandingsList
			all_history_2 := history_2.Root.StandingsTable.StandingsList
			all_history_3 := history_3.Root.StandingsTable.StandingsList

			// History info
			fmt.Print("-----------------\n")
			color.HiMagenta("All time Standing")
			fmt.Print("-----------------\n")
			for i := 0; i < time.Now().Year()-1952; i++ {
				switch {
				case i < 29:
					fmt.Printf(
						"%s: %s(%s) in %s %s pt %s wins \n",
						color.HiGreenString(all_history_1[i].Season),
						color.HiCyanString(all_history_1[i].DriverStanding.Driver.Name),
						all_history_1[i].DriverStanding.Driver.Nationality,
						color.HiRedString(all_history_1[i].DriverStanding.Constructor.Name),
						all_history_1[i].DriverStanding.Points,
						all_history_1[i].DriverStanding.Wins)
				case i < 58:
					fmt.Printf(
						"%s: %s(%s) in %s %s pt %s wins \n",
						color.HiGreenString(all_history_2[i-29].Season),
						color.HiCyanString(all_history_2[i-29].DriverStanding.Driver.Name),
						all_history_2[i-29].DriverStanding.Driver.Nationality,
						color.HiRedString(all_history_2[i-29].DriverStanding.Constructor.Name),
						all_history_2[i-29].DriverStanding.Points,
						all_history_2[i-29].DriverStanding.Wins)
				default:
					fmt.Printf(
						"%s: %s(%s) in %s %s pt %s wins \n",
						color.HiGreenString(all_history_3[i-58].Season),
						color.HiCyanString(all_history_3[i-58].DriverStanding.Driver.Name),
						all_history_3[i-58].DriverStanding.Driver.Nationality,
						color.HiRedString(all_history_3[i-58].DriverStanding.Constructor.Name),
						all_history_3[i-58].DriverStanding.Points,
						all_history_3[i-58].DriverStanding.Wins)
				}
			}
		case "ds":
			type DriverStandings struct {
				Root struct {
					StandingsTable struct {
						StandingsList struct {
							DriverStanding []struct {
								Position string `json:"-positionText"`
								Points   string `json:"-points"`
								Wins     string `json:"-wins"`
								Driver   struct {
									Name string `json:"FamilyName"`
								} `json:"Driver"`
								Team struct {
									Name string `json:"Name"`
								} `json:"Constructor"`
							} `json:"DriverStanding"`
						} `json:"StandingsList"`
					} `json:"StandingsTable"`
				} `json:"MRData"`
			}

			var dictionary_ds = get_structed_info_from_api("http://ergast.com/api/f1/current/driverStandings")
			var driver_standings DriverStandings
			sth := json.Unmarshal(dictionary_ds.Bytes(), &driver_standings)
			if sth != nil {
				fmt.Println("Error:", sth)
			}

			drivers_top := driver_standings.Root.StandingsTable.StandingsList.DriverStanding

			// Printing top drivers
			fmt.Print("-----------------\n")
			color.HiMagenta("Drivers standings")
			fmt.Print("-----------------\n")
			for i := 0; i < len(drivers_top); i++ {
				fmt.Printf(
					"%s. %s in %s %s pt %s wins\n",
					drivers_top[i].Position,
					color.HiCyanString(drivers_top[i].Driver.Name),
					color.HiRedString(drivers_top[i].Team.Name),
					drivers_top[i].Points,
					drivers_top[i].Wins,
				)
			}
			fmt.Println()

		case "cs":
			type TeamsStandings struct {
				Root struct {
					StandingsTable struct {
						StandingsList struct {
							TeamsStanding []struct {
								Team struct {
									Name        string `json:"Name"`
									Nationality string `json:"Nationality"`
								} `json:"Constructor"`
								Points   string `json:"-points"`
								Wins     string `json:"-wins"`
								Position string `json:"-positionText"`
							} `json:"ConstructorStanding"`
						} `json:"StandingsList"`
					} `json:"StandingsTable"`
				} `json:"MRData"`
			}

			dictionary_ts := get_structed_info_from_api("http://ergast.com/api/f1/current/constructorStandings")
			var teams_standings TeamsStandings
			sth_1 := json.Unmarshal(dictionary_ts.Bytes(), &teams_standings)
			if sth_1 != nil {
				fmt.Println("Error:", sth_1)
			}

			teams_top := teams_standings.Root.StandingsTable.StandingsList.TeamsStanding

			// Printing top teams
			fmt.Print("---------------------\n")
			color.HiMagenta("Constructors standings")
			fmt.Print("---------------------\n")
			for i := 0; i < len(teams_top); i++ {
				fmt.Printf(
					"%s. %s %s pt %s wins\n",
					teams_top[i].Position,
					color.HiRedString(teams_top[i].Team.Name),
					teams_top[i].Points,
					teams_top[i].Wins,
				)
			}
			fmt.Println()

		case "schedule":
			type Schedule struct {
				Root struct {
					RaceTable struct {
						Race []struct {
							Round   string `json:"-round"`
							Name    string `json:"RaceName"`
							Date    string `json:"Date"`
							Circuit struct {
								CircuitName string `json:"CircuitName"`
							} `json:"Circuit"`
						} `json:"Race"`
					} `json:"RaceTable"`
				} `json:"MRData"`
			}

			dictionary_sd := get_structed_info_from_api("http://ergast.com/api/f1/current")
			var schedule Schedule
			sth_3 := json.Unmarshal(dictionary_sd.Bytes(), &schedule)
			if sth_3 != nil {
				fmt.Println("Error:", sth_3)
			}

			race := schedule.Root.RaceTable.Race

			// Shedule
			fmt.Print("--------\n")
			color.HiMagenta("Schedule")
			fmt.Print("--------\n")
			for i := 0; i < len(race); i++ {
				if time.Now().Format("2006-01-02") > race[i].Date {
					fmt.Printf(
						"%s %s | %s\n",
						color.RedString(race[i].Date),
						race[i].Name,
						color.HiWhiteString(race[i].Circuit.CircuitName),
					)
				} else {
					fmt.Printf(
						"%s %s | %s\n",
						color.HiGreenString(race[i].Date),
						race[i].Name,
						color.HiWhiteString(race[i].Circuit.CircuitName),
					)
				}
			}
		}
	}
}
