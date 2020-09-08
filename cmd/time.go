package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/spf13/cobra"
)

type obj struct {
	Message   string   `json:"message"`
	Curr_time string   `json:"curr_time"`
	Sys_time  string   `json:"sys_time"`
	Time      []string `json:"time"`
}

var url string
var station string
var line string
var test obj

func init() {
	//cobra.OnInitialize(initConfig)
	mtrtimeCmd.Flags().StringVarP(&station, "station", "s", "", "your MTR Station (required)")
	mtrtimeCmd.Flags().StringVarP(&line, "line", "l", "", "your MTR Line (required)")
	mtrtimeCmd.MarkFlagRequired("station")
	mtrtimeCmd.MarkFlagRequired("line")
	rootCmd.AddCommand(mtrtimeCmd)
}

var mtrtimeCmd = &cobra.Command{
	Use:   "time",
	Short: "Check out the MTR arrive time for the next 3 trains",
	Long:  `MTR Arrive time for the next 3 trains`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("hi %s, %s", station, line)
		//line=WRL&sta=YUL
		var result map[string]interface{}
		var msg, curr_time, sys_time, up, down interface{}
		//var uptime, downtime interface{}
		name := line + "-" + station

		url += "https://rt.data.gov.hk/v1/transport/mtr/getSchedule.php?" + "line=" + line + "&sta=" + station
		fmt.Printf("api: %s\n", url)
		fmt.Println(" ")

		response, err1 := http.Get(url)
		if err1 != nil {
			log.Println(err1)
			os.Exit(1)
		}

		data, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			log.Println(err2)
			os.Exit(1)
		}

		err3 := json.Unmarshal([]byte(data), &result)
		if err3 != nil {
			log.Println(err3)
			os.Exit(1)
		}
		//fmt.Println(result)
		//fmt.Println(" ")
		//fmt.Printf("%+v\n", result)

		for key, value := range result {
			// Each value is an interface{} type, that is type asserted as a string
			//fmt.Println(key, value)
			if key == "message" {
				msg = value
			}
		}

		birds := result["data"].(map[string]interface{})
		birds2 := birds[name].(map[string]interface{})

		var bool_up, bool_down = false, false
		for key, value := range birds2 {
			// Each value is an interface{} type, that is type asserted as a string
			if key == "curr_time" {
				curr_time = value
			} else if key == "sys_time" {
				sys_time = value
			} else if key == "UP" {
				up = value
				bool_up = true
			} else if key == "DOWN" {
				down = value
				bool_down = true
			}
		}

		//fmt.Println(birds[name])
		//fmt.Println(" ")

		fmt.Printf("Message: %s\n", msg)
		fmt.Printf("Current time: %s\n", curr_time)
		fmt.Printf("System time: %s\n\n", sys_time)

		if bool_up {
			fmt.Printf("Up time:\n")
			for i := 0; i < 3; i++ {
				value := reflect.ValueOf(up)
				//fmt.Println(value.Len())
				c := value.Index(i).Interface().(map[string]interface{})
				fmt.Println(c)
				//fmt.Printf("%T\n", c)
			}
		}

		fmt.Println(" ")

		if bool_down {
			fmt.Printf("Down time:\n")
			for i := 0; i < 3; i++ {
				value := reflect.ValueOf(down)
				c := value.Index(i).Interface().(map[string]interface{})
				fmt.Println(c)
			}
		}

		fmt.Printf("\nRaw JSON data:%s", string(data))
		//fmt.Println(string(data))
	},
}
