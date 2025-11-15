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
			var getById *Agency
			if id > len(provinces[user.region]) {
				fmt.Println("❌ Invalid ID. Please enter a number.")
			} else {
				for _, v := range provinces[user.region] {
					if id == v.Id {
						// fmt.Println(v.Name, v.Address, v.Phone, v.EmployeeCount)
						// getById = &Agency{Id: v.Id, Name: v.Name, Address: v.Address, Phone: v.Phone, EmployeeCount: v.EmployeeCount, RegisterDate: v.RegisterDate}
						getById = &v

					}
				}
				fmt.Println(getById)
			}
		}

	case "edit":
		{
			strid := scanInput("enter id")
			id, err := strconv.Atoi(strid)
			if err != nil {
				fmt.Println("❌ Invalid ID. Please enter a number.")
				return
			}

			agency := getAgencyById(id)
			fmt.Println(agency)
			if agency != nil {
				for i := range provinces[user.region] {
					if provinces[user.region][i].Name == agency.Name {
						var (
							inputName    string
							inputAddress string
							inputPhone   string
						)

						inputName = scanInput("enter the name")
						inputAddress = scanInput("enter the address")
						inputPhone = scanInput("enter the phone-number")
						strid := scanInput("enter employeeCount")
						inputEmployeeCount, err := strconv.Atoi(strid)
						// check and validate EmployeeCount
						if err != nil {
							fmt.Println("Invalid INPUT employee count must be an Integer")
						}
						provinces[user.region][i].Name = inputName
						provinces[user.region][i].Address = inputAddress
						provinces[user.region][i].Phone = inputPhone
						provinces[user.region][i].EmployeeCount = inputEmployeeCount
						provinces[user.region][i].RegisterDate = getTime()
						fmt.Println("before edit", agency)
						fmt.Println("After edit", provinces[user.region][i])

						askSave := scanInput("Do you want to save edited data? 1:Yes 2:No")
						switch askSave {
						case "1":
							saveProvinceData()
							fmt.Println("New Edited Data Saved")
						case "2":

							fmt.Println("Changes are not saved")
							provinces[user.region][i] = *agency

						default:
							fmt.Println("Enter a valid Input Option")
						}
						break
					}
					fmt.Println("all", provinces[user.region])
					fmt.Println("edited", agency)
				}
			} else {
				fmt.Println("Wrong Input ID must be valid id See List Command for agency ID's ")
			}
		}
	case "add":
		{
			if len(user.region) > 0 {

				var IdStoreStructre int = len(provinces[user.region]) + 1
				var (
					inputName    string
					inputAddress string
					inputPhone   string
				)

				inputName = scanInput("enter the name")
				inputAddress = scanInput("enter the address")
				inputPhone = scanInput("enter the phone-number")
				strid := scanInput("enter employeeCount")
				inputEmployeeCount, err := strconv.Atoi(strid)
				// check and validate EmployeeCount
				if err != nil {
					fmt.Println("Invalid INPUT employee count must be an Integer")
				}
				provinces[user.region] = append(
					provinces[user.region], Agency{
						Id:            IdStoreStructre,
						Name:          inputName,
						Address:       inputAddress,
						Phone:         inputPhone,
						EmployeeCount: inputEmployeeCount,
						RegisterDate:  getTime(),
					},
				)

				saveProvinceData()
				fmt.Println(provinces[user.region])
			} else {
				fmt.Println("for adding Agency you have to select a region first")
			}
		}
	case "status":
		{
			var regionEmployeeCount int
			for i := 0; i < len(provinces[user.region]); i++ {

				regionEmployeeCount += provinces[user.region][i].EmployeeCount
			}

			fmt.Printf("Current Region Status | Active Agencies : %d  regionEmployee Count: %d \n ", len(provinces[user.region]), regionEmployeeCount)
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

func getAgencyById(id int) (agency *Agency) {
	if id > len(provinces[user.region]) {
		fmt.Println("❌ Invalid ID. Please enter a number.")
	} else {
		for _, v := range provinces[user.region] {
			if id == v.Id {
				// fmt.Println(v.Name, v.Address, v.Phone, v.EmployeeCount)
				// getById = &Agency{Id: v.Id, Name: v.Name, Address: v.Address, Phone: v.Phone, EmployeeCount: v.EmployeeCount, RegisterDate: v.RegisterDate}
				agency = &v

			}
		}
		return agency
	}
	return nil
}

func getTime() (datetime string) {
	// time format
	t := time.Now()
	datetime = fmt.Sprintf("%d-%02d-%02d %02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	return datetime
}
