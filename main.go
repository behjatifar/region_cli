package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var provinces = make(map[string][]Agency)

type Agency struct {
	Id            int
	Name          string
	Address       string
	Phone         string
	RegisterDate  string
	EmployeeCount int
}
type User struct {
	region string
}

// global user
var user User

func main() {
	// region flag
	region := flag.String("region", "", "choosing region")
	flag.Parse()
	command := flag.String("command", "", "for set command")
	flag.Parse()

	if len(*region) > 1 {
		setRegion(*region)
	}

	for {

		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}
}
func runCommand(command string) {
	switch command {
	case "help":
		{

		}
	case "region":
		{
			region := scanInput("enter your region")
			if len(provinces[region]) > 0 {
				user = User{region: region}
				fmt.Printf(" region %s Set successfully \n", user.region)

			} else {
				fmt.Printf("Region %s not available .. ! \n", region)
			}

		}

	case "list":
		{
			fmt.Println("---------------------------------")
			fmt.Println("list of agencies in your region")
			for i := 0; i < len(provinces[user.region]); i++ {
				fmt.Println("ID:", provinces[user.region][i].Id, provinces[user.region][i].Name, "\t")
			}
			fmt.Println("---------------------------------")
		}

	case "get":
		{

			strid := scanInput("Please Enter Agency ID also you can See Agency Names + ID in \"list\" command ")
			id, err := strconv.Atoi(strid)
			if err != nil {
				fmt.Println("❌ Invalid ID. Please enter a number.")
				return
			}
			// fmt.Println(provinces[user.region])
			for _, v := range provinces[user.region] {
				if id == v.Id {
					fmt.Println(v.Name, v.Address, v.Phone, v.EmployeeCount)
				}
			}
		}

	case "edit":
		{

		}
	case "add":
		{
			if len(user.region) > 0 {

				// time format
				t := time.Now()
				var time string = fmt.Sprintf("%d-%02d-%02d %02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
				var IdStoreStructre int = len(provinces[user.region]) + 1
				provinces[user.region] = append(
					provinces[user.region], Agency{
						Id:            IdStoreStructre,
						Name:          scanInput("enter the name"),
						Address:       scanInput("enter the address"),
						Phone:         scanInput("enter the phone-number"),
						EmployeeCount: 5,
						RegisterDate:  time,
					},
				)
				saveProvinceData()
				fmt.Println(provinces[user.region])
			} else {
				fmt.Println("for adding Agency you have to select a region first")
			}
		}
	case "exit":
		{
			os.Exit(0)
		}
	}

}

func saveProvinceData() {
	data, err := json.MarshalIndent(provinces, "", "")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("provinces.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

func loadProvinceData(provinces *map[string][]Agency) {
	data, err := os.ReadFile("provinces.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &provinces)
	if err != nil {
		panic(err)
	}
	fmt.Println("data loaded successfully✅")

}
func init() {
	loadProvinceData(&provinces)
}
func setRegion(regName string) {
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("please enter your region")
	// scanner.Scan()
	// var region string = scanner.Text()
	fmt.Println("reg", provinces["tehran"])
	if len(provinces[regName]) > 0 {
		user = User{region: regName}
		fmt.Printf(" region %s Set successfully \n", user.region)

	} else {
		fmt.Printf("Region %s not available .. ! \n", regName)
	}
}
func scanInput(scanTxt string) (result string) {
	scanner := bufio.NewScanner(os.Stdin)
	// scan user input
	fmt.Println(scanTxt)
	scanner.Scan()

	result = scanner.Text()

	return result
}
