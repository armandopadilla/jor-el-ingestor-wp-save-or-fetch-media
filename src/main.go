package main

import (
	"encoding/json"
	"fmt"
)

/**
* Main driver
**/
func main() {

	// Fetch data from SQS
	data := fetchDataFromWPPostQueue()

	// For each one check if we've seen it already.
	for i := 0; i < len(data); i++ {
		// UnMarshal
		message := Message{}

		err := json.Unmarshal([]byte(*data[i].Body), &message)

		if err != nil {
			fmt.Println("Error", err)
			break
		}

		// Get the ID
		//id := message.ID
		title := message.Title

		// Check DB
		// post := getPostFromDB(id)

		// Is present and dates match?
		//if post == 0 {
		// Save to S3
		fileName := fmt.Sprintf("%s.json", title)
		saveToS3(fileName, *data[i].Body)

		// Check if there are attachments.
		// if attachements - place each media into sqs
		//if message.Media.FeaturedMedia != 0 {
		//    saveToMediaSQS()
		//}

		// Save to DB
		//saveToDB()
		// }
	}
}

/*
{
    dataLakeLocation: '' // Where the data lives for this post. Used to update the media payload.
}
*/
