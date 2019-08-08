package main

/*
Media struct {
FeaturedMedia int `json:"featuredMedia"`
WpMediaLink []struct {
Href string `json:"href"`
} `json:"wpMediaLink"`
} `json:"media"`
*/

// Message - Message structure
type Message struct {
	ID    int `json:ID`
	Media struct {
		FeaturedMedia int `json:FeaturedMedia`
		WPMediaLink   []struct {
			Href string `json:Href`
		}
	} `json:Media`
	Date  string `json:Date`
	Title string `json:Title`
}
