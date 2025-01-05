package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Barang struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Barang string             `bson:"nama_barang,omitempty" json:"nama_barang,omitempty"`
	Kategori    string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Merek       string             `bson:"merek,omitempty" json:"merek,omitempty"`
	Tahun       int                `bson:"tahun,omitempty" json:"tahun,omitempty"`
	Harga       float64            `bson:"harga,omitempty" json:"harga,omitempty"`
	Stok        int                `bson:"stok,omitempty" json:"stok,omitempty"`
	Gambar      string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
}

type DetailBarang struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Barang      Barang             `bson:"barang,omitempty" json:"barang,omitempty"`
	Warna       string             `bson:"warna,omitempty" json:"warna,omitempty"`
	Berat       float64            `bson:"berat,omitempty" json:"berat,omitempty"`
	Dimensi     string             `bson:"dimensi,omitempty" json:"dimensi,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	TanggalMasuk string            `bson:"tanggal_masuk,omitempty" json:"tanggal_masuk,omitempty"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}
