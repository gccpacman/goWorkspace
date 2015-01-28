package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://60.195.249.97/evergrande/genhtml/matchldesc.action?id=311&s=2014&t=5829")
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("%s", err)
			os.Exit(1)
		}
		fmt.Println("%s\n", string(contents))

		contentsBytes := []byte(contents)
		fmt.Println(contentsBytes)
		j, err1 := jason.NewObjectFromBytes(contentsBytes)

		fmt.Printf("%v-> %s (err: %s)\n", contents, contentsBytes, err1)

		matchlist, err := j.GetObjectArray("matchlist")
		if err != nil {
			fmt.Println("%s", err)
			os.Exit(1)
		}
		for k, match := range matchlist {

			fmt.Println(k, match)
			club_name_home, err := match.GetString("club_name_home")
			if err != nil {
				fmt.Println("%s", err)
				continue
			}
			fmt.Println("club_name_home=", club_name_home)

			match_score_home, err := match.GetInt64("match_score_home")
			if err != nil {
				fmt.Println("%s", err)
				continue
			}
			fmt.Println("match_score_home=", match_score_home)

		}

	}
	// fmt.Println("resp = ", resp, "(err = ", err, ")")

}
