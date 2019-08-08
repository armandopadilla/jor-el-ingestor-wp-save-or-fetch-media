package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//var db, err = sql.Open("mysql", "")

/**
 * Get the post from the DB, if we have one
 *
 **
func getPostFromDB(postID int) int {
	sql := fmt.Sprintf("SELECT id, modified_date FROM wp-posts WHERE id = %d", postID)
	results, err := db.Query(sql)

	if err != nil {
		fmt.Println("Error", err)
		return 0
	}

	return 0
}

/**
 * Save the post to the DB.  Only going to save, id, and modified date
 **
func insertPostToDB(postID int, modifiedDate string) {

}
*/
