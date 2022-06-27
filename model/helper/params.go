package helper

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadParams(r *http.Request, key string) (primitive.ObjectID, error) {
	params := mux.Vars(r)
	id := params[key]
	oid, err := primitive.ObjectIDFromHex(id)
	return oid, err
}