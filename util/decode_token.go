package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pattanunNP/WishyWishyBackend/models"
)

func Decode(Token string) (*models.LineProfile, error) {

	endpoint := fmt.Sprintf("https://api.line.me/oauth2/v2.1/verify?access_token=%s", Token[1:])

	// fmt.Println(endpoint)

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(err)

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	if res.StatusCode == 200 {
		endpoint := fmt.Sprintf("https://api.line.me/v2/profile")
		client := http.Client{}
		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			fmt.Println(err)
		}

		token := fmt.Sprintf("Bearer %s", Token[1:])
		req.Header.Add("Authorization", token)

		res, err := client.Do(req)
		if err != nil {

			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
		var profile models.LineProfile
		err = json.Unmarshal(body, &profile)
		if err != nil {
			fmt.Println(err)
		}
		return &profile, nil
	} else if res.StatusCode == 400 {
		return nil, fmt.Errorf("Invalid token")

	}
	return nil, fmt.Errorf("Unknown error")

}
