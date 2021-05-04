package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type User struct {
	first  string
	middle string
	last   string
	born   int
}

func main() {
	fmt.Println("Hello World!!")

	// addUser()
	// updataUser()
	deleteUser()
	readUser()

}

func readUser() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			if err == iterator.Done {
				break
			}
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}

func addUser() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	// Create Client
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// データ追加

	// user1 := User{
	// 	first: "Ada",
	// 	last: "Lovelance",
	// 	born: 1815,
	// }

	// user2 := User{
	// 	first:  "Ada",
	// 	middle: "Mathison",
	// 	last:   "Lovelace",
	// 	born:   1815,
	// }

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func updataUser() {
	// 初期化
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// データ更新
	_, updateError := client.Collection("users").Doc("user2").Set(ctx, map[string]interface{}{
		"first":  "Yeah",
		"middle": "deleted items",
	}, firestore.MergeAll)
	if updateError != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	// 切断
	defer client.Close()
}

func deleteUser() {
	// 初期化
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// フィールド削除
	_, errorDelete := client.Collection("users").Doc("user2").Update(ctx, []firestore.Update{
		{
			Path:  "middle",
			Value: firestore.Delete,
		},
	})
	if errorDelete != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}
