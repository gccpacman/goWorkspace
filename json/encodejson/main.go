package main

import (
	"encoding/json"
	"fmt"
	// "github.com/antonholmquist/jason"
	"io/ioutil"
	"net/http"
	"os"
)

/*
	所有字段必须大写, 否则无法Unmarshal
*/
type Match struct {
	Comp_short_chn   string
	Match_season     string
	Match_round      int64
	Match_mdate      string
	Club_name_home   string
	Club_name_away   string
	Match_score_home int64
	Match_score_away int64
}

type MatchList struct {
	Matchlist []Match
}

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	resp, err := http.Get("...")
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
		fmt.Printf("%s\n", string(contents))
		fmt.Println()
		// contentsBytes := []byte(contents)
		// fmt.Println(contentsBytes)

		var ss MatchList
		// contStr := `	{"matchlist":[{"match_mdate":"2014-07-24","match_season":"2014","match_round":60,"club_name_home":"河南建业","match_score_away":1,"comp_short_chn":"中杯","club_name_away":"广州恒大","match_score_home":2},{"match_mdate":"2014-07-15","match_season":"2014","match_round":30,"club_name_home":"广州恒大","match_score_away":0,"comp_short_chn":"中杯","club_name_away":"广东日之泉","match_score_home":1}]}`
		// json.Unmarshal([]byte(contStr), &ss)
		json.Unmarshal([]byte(contents), &ss)
		fmt.Println(ss)
		// fmt.Println(ss.matchlist[0].club_name_home)
		for _, m := range ss.Matchlist {
			fmt.Println(m.Match_mdate, m.Comp_short_chn, m.Match_season, "赛季, 轮次", m.Match_round)
			fmt.Println(m.Club_name_home, m.Match_score_home, "-", m.Match_score_away, m.Club_name_away)
		}

		var s Serverslice
		str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
		fmt.Println(str)
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s)
		for _, server := range s.Servers {
			fmt.Println("ServerName =", server.ServerName)
			fmt.Println("ServerIP =", server.ServerIP)
		}

		// objMap, err :=

		// j, err1 := jason.NewObjectFromBytes(contentsBytes)

		// fmt.Printf("%v-> %s (err: %s)\n", contents, contentsBytes, err1)

		// matchlist, err := j.GetObjectArray("matchlist")
		// if err != nil {
		// 	fmt.Println("%s", err)
		// 	os.Exit(1)
		// }
		// for k, match := range matchlist {

		// 	fmt.Println(k, match)
		// 	club_name_home, err := match.GetString("club_name_home")
		// 	if err != nil {
		// 		fmt.Println("%s", err)
		// 		continue
		// 	}
		// 	fmt.Println("club_name_home=", club_name_home)

		// 	match_score_home, err := match.GetInt64("match_score_home")
		// 	if err != nil {
		// 		fmt.Println("%s", err)
		// 		continue
		// 	}
		// 	fmt.Println("match_score_home=", match_score_home)

		// }

	}
	// fmt.Println("resp = ", resp, "(err = ", err, ")")

}
