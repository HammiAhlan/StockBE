package _714220035

import (
	"fmt"
	"testing"

	"github.com/irgifauzi/back-bola/model"
	"github.com/irgifauzi/back-bola/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertDetailBarang(t *testing.T) {
	nama_barang := "Laptop Gaming"
	kategori := "Elektronik"
	merek := "ASUS ROG"
	tahun := 2023
	harga := 25000000.0
	stok := 10
	gambar := "https://example.com/laptop-gaming.jpg"

	barang := model.Barang{
		Nama_Barang: nama_barang,
		Kategori:    kategori,
		Merek:       merek,
		Tahun:       tahun,
		Harga:       harga,
		Stok:        stok,
		Gambar:      gambar,
	}

	warna := "Hitam"
	berat := 2.5
	dimensi := "35x25x2 cm"
	deskripsi := "Laptop gaming dengan performa tinggi untuk kebutuhan profesional."
	tanggal_masuk := "2025-01-01"

	insertedID, err := module.InsertDetailBarang(module.MongoConn, "detail_barang", barang, warna, berat, dimensi, deskripsi, tanggal_masuk)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
		return
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetDetailBarangFromID(t *testing.T) {
	id := "668e488005df6fa1b2719599"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	detailBarang, err := module.GetDetailBarangFromID(objectID, module.MongoConn, "detail_barang")
	if err != nil {
		t.Fatalf("error calling GetDetailBarangFromID: %v", err)
	}
	fmt.Println(detailBarang)
}

func TestGetAllDetailBarang(t *testing.T) {
	data := module.GetAllDetailBarang(module.MongoConn, "detail_barang")
	fmt.Println(data)
}

func TestDeleteDetailBarangByID(t *testing.T) {
	id := "668e488005df6fa1b2719599"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteDetailBarangByID(objectID, module.MongoConn, "detail_barang")
	if err != nil {
		t.Fatalf("error calling DeleteDetailBarangByID: %v", err)
	}

	_, err = module.GetDetailBarangFromID(objectID, module.MongoConn, "detail_barang")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
