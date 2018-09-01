// Copyright © 2018 Ricardo Mattos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Coin struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	VolumeUsd24h     string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvaliableSupply  string `json:"avaliable_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

const apiUrl string = "https://api.coinmarketcap.com/v1/ticker/"

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "Get the actual value of Crypto Coins",
	Long:  `Get the actual value of Crypto Coins just passing his name as parameter.`,
	Run: func(cmd *cobra.Command, args []string) {

		var coinObj = []Coin{}

		if all, _ := cmd.Flags().GetBool("all"); all {
			resp, err := http.Get(apiUrl)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer resp.Body.Close()

			err = json.NewDecoder(resp.Body).Decode(&coinObj)
			if err != nil {
				fmt.Println("coult not parse json obj")
				return
			}
		} else if coinName, _ := cmd.Flags().GetString("name"); coinName != "" {
			resp, err := http.Get(fmt.Sprintf("%s/%s", apiUrl, coinName))
			if err != nil {
				fmt.Println(err)
				return
			}

			defer resp.Body.Close()

			err = json.NewDecoder(resp.Body).Decode(&coinObj)
			if err != nil {
				fmt.Println("coult not parse json obj")
				return
			}
		}

		for idx, element := range coinObj {
			fmt.Printf("\n#%d\nName:%s\nPrice Usd:%s\nPrice BTC:%s\n\n", idx, element.Name, element.PriceUsd, element.PriceBtc)
		}
	},
}

func init() {
	rootCmd.AddCommand(coinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	coinCmd.Flags().StringP("name", "n", "", "Actual value of <coin_name> crypto coin")
	coinCmd.Flags().Bool("all", false, "Actual values of all crypto coins")
}
