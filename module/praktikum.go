package module

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/irgifauzi/back-bola/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDetailBarang(db *mongo.Database, col string, barang model.Barang, warna string, berat float64, dimensi string, deskripsi string, tanggal_masuk string) (insertedID primitive.ObjectID, err error) {
	detailBarang := bson.M{
		"barang":       barang,
		"warna":        warna,
		"berat":        berat,
		"dimensi":      dimensi,
		"deskripsi":    deskripsi,
		"tanggal_masuk": tanggal_masuk,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), detailBarang)
	if err != nil {
		fmt.Printf("InsertDetailBarang: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllDetailBarang(db *mongo.Database, col string) (data []model.DetailBarang) {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllDetailBarang: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDetailBarangFromID(_id primitive.ObjectID, db *mongo.Database, col string) (detailBarang model.DetailBarang, errs error) {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&detailBarang)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return detailBarang, fmt.Errorf("no data found for ID %s", _id)
		}
		return detailBarang, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return detailBarang, nil
}

func UpdateDetailBarang(db *mongo.Database, col string, id primitive.ObjectID, barang model.Barang, warna string, berat float64, dimensi string, deskripsi string, tanggal_masuk string) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"barang":       barang,
			"warna":        warna,
			"berat":        berat,
			"dimensi":      dimensi,
			"deskripsi":    deskripsi,
			"tanggal_masuk": tanggal_masuk,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("UpdateDetailBarang: %v", err)
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}
	return nil
}

func DeleteDetailBarangByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
