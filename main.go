package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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

var IdStoreStructre int = len(provinces[user.region]) + 1

func main() {
	provinces["tehran"] = []Agency{
		{Id: 1, Name: "markaz", Address: "xyz", Phone: "021-12345678", EmployeeCount: 1, RegisterDate: "2020-01-01"},
		{Id: 2, Name: "gharb", Address: "xyz", Phone: "021-12345678", EmployeeCount: 1, RegisterDate: "2020-01-01"},
	}
	// region flag
	region := flag.String("region", "", "choosing region")
	flag.Parse()
	command := flag.String("command", "get", "for set command")
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
				fmt.Println(provinces[user.region][i].Name, "\t")
			}
			fmt.Println("---------------------------------")
		}

	case "get":
		{
			var id int = 1
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
			// name , address , phone , employeeCount | auto compeleted fields Id , RegisterDate
			if len(user.region) > 0 {
				// var name, address, phone string
				// var employeeCount int scanInput()
				fmt.Println("Enter Name")
				// name = scanInput()
				fmt.Println("Enter Address")
				// address = scanInput()
				fmt.Println(time.Now().String())
				provinces[user.region] = []Agency{
					{Id: IdStoreStructre,
						Name:          scanInput("enter the name"),
						Address:       scanInput("enter the address"),
						Phone:         scanInput("enter the phone-number"),
						EmployeeCount: 5,
						RegisterDate:  time.Now().String()},
				}
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

func loadProvinceData(provinces map[string][]Agency) (pData map[string][]Agency) {
	data, err := os.ReadFile("provinces.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &provinces)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	return provinces
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
