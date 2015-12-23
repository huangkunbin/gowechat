package lib

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	
)

type MusicResponseBody struct {
	Status string  `json:"status"`
	Code   int     `json:"code"`
	Data   AllData `json:"data"`
}

type AllData struct {
	Data MusicData `json:"data"`
}
type MusicData struct {
	Count      int         `json:"count"`
	TotalPages int         `json:"totalPages"`
	S          string      `json:"s"`
	NowPage    string      `json:"nowPage"`
	List       []MusicList `json:"list"`
}
type MusicList struct {
	SongId    int    `json:"songId"`
	SongName  string `json:"songName"`
	UserName  string `json:"userName"`
	AlbumPic  string `json:"albumPic"`
	AlbumName string `json:"albumName"`
	SongUrl   string `json:"songUrl"`
}

func GetMusic(str string) (MusicResponseBody, error){

	//str := "红豆"
	url := "http://a.apix.cn/geekery/music/query?s=" + str + "&limit=2&p=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("apix-key", "84d3e0847c714d39725ef1dd6c9fb5db")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	var config MusicResponseBody
	
	err := json.Unmarshal([]byte(string(body)), &config)
	
	return config,err
	
	/*if err := json.Unmarshal([]byte(string(body)), &config); err == nil {
		fmt.Println(config)
		fmt.Println(config.Data.Data.List)
	} else {
		fmt.Println(err)
	}*/

}