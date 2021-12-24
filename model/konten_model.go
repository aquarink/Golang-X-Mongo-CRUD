package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type KontenFull struct {
	// type KontenPlain struct {
	Id      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Kode    string             `json:"kode"`
	Tipe    string             `json:"tipe"`
	Judul   string             `json:"judul"`
	Isi     string             `json:"isi"`
	Thumb   string             `json:"thumb"`
	Tanggal string             `json:"tanggal"`
	Slug    string             `json:"slug"`
	Tag     string             `json:"tag"`
	View    int                `json:"view"`
}

type KontenPlain struct {
	// type KontenFull struct {
	Id      string
	Kode    string
	Tipe    string
	Judul   string
	Isi     string
	Thumb   string
	Tanggal string
	Slug    string
	Tag     string
	View    int
}

type KontenInput struct {
	Kode    string `json:"kode"`
	Tipe    string `json:"tipe"`
	Judul   string `json:"judul"`
	Isi     string `json:"isi"`
	Thumb   string `json:"thumb"`
	Tanggal string `json:"tanggal"`
	Slug    string `json:"slug"`
	Tag     string `json:"tag"`
	View    int    `json:"view"`
}

type KontenUpdate struct {
	Judul string `json:"judul"`
	Isi   string `json:"isi"`
	Thumb string `json:"thumb"`
	Tag   string `json:"tag"`
}

type KontenUpdateView struct {
	View int `json:"view"`
}
