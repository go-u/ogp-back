package users

import (
	"context"
	"log"
	"net/http"
	"server/initialize/db"
	"server/initialize/firebaseauth"
	"server/tools"
)

func DeleteAcount(r *http.Request) error {
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return err
	}

	// start transaction
	tx, err := db.Connection.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	// delete from user table
	_, err = user.Delete(context.Background(), tx)
	if err != nil {
		_ = tx.Rollback() // rollback
		return err
	}

	// delete from firebase
	err = firebaseauth.Client.DeleteUser(context.Background(), user.UID)
	if err != nil {
		log.Printf("firebase: error deleting user: %v\n", err)
		return err
	}
	// commit tx
	err = tx.Commit()
	return err
}
