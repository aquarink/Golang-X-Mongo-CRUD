package controller

import (
	"juriback2/model"
	"juriback2/service"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type initKontenController struct {
	kontenServiceInterface service.KonytenServiceInterface
}

func KontenController(kontenService service.KonytenServiceInterface) *initKontenController {
	return &initKontenController{kontenService}
}

// Controller

func (handler *initKontenController) Landing(w http.ResponseWriter, r *http.Request) {
	_, err := handler.kontenServiceInterface.SerKontenSemua()
	if err != nil {
		log.Println("ERROR Landing : " + err.Error())
		http.Error(w, "ERROR Landing", http.StatusInternalServerError)
		return
	}

	// log.Println(data)
	http.Error(w, "Berhasil Landing", http.StatusOK)
}

func (handler *initKontenController) ById(w http.ResponseWriter, r *http.Request) {
	data, err := handler.kontenServiceInterface.SerKontenCariById("61c41ef2b7d546e54a4438c8")
	if err != nil {
		log.Println("ERROR ById : " + err.Error())
		http.Error(w, "ERROR ById", http.StatusInternalServerError)
		return
	}

	log.Println(data)
	http.Error(w, "Berhasil ById", http.StatusOK)
}

func (handler *initKontenController) Tambah(w http.ResponseWriter, r *http.Request) {

	judul := "8u  usdh234567h @#$%^&Ihv hg hdfhdf"
	isi := "sd;sah h dsh hd fskd9- 38434873 jdfdfh"

	min := 1000000000
	max := 9999999999
	ran := strconv.Itoa(rand.Intn(max-min) + min)

	var re = regexp.MustCompile("[^a-z0-9]+")
	slug := strings.Trim(re.ReplaceAllString(strings.ToLower(judul), "-"), "-") + "-" + ran

	var konten model.KontenInput

	konten.Kode = ran
	konten.Tipe = "artikel"
	konten.Judul = judul
	konten.Isi = isi
	konten.Thumb = "/var/www/html"
	konten.Slug = slug
	konten.Tag = "aaaa,sss,dddd"

	data, err := handler.kontenServiceInterface.SerKontenInsert(konten)
	if err != nil {
		log.Println("ERROR Tambah : " + err.Error())
		http.Error(w, "ERROR Tambah", http.StatusInternalServerError)
		return
	}

	log.Println(data)
	http.Error(w, "Berhasil Tambah", http.StatusOK)
}

func (handler *initKontenController) Ubah(w http.ResponseWriter, r *http.Request) {

	judul := "Ubah data judul"
	isi := "datas isi diubah"

	var konten model.KontenUpdate

	konten.Judul = judul
	konten.Isi = isi
	konten.Thumb = "/var/www/html"
	konten.Tag = "aaaa,sss,dddd"

	data, err := handler.kontenServiceInterface.SerKontenUpdate("4567446831", konten)
	if err != nil {
		log.Println("ERROR Ubah : " + err.Error())
		http.Error(w, "ERROR Ubah", http.StatusInternalServerError)
		return
	}

	log.Println(data)
	http.Error(w, "Berhasil Ubah", http.StatusOK)
}

func (handler *initKontenController) Views(w http.ResponseWriter, r *http.Request) {

	slug := "8u-usdh234567h-ihv-hg-hdfhdf-4567446831"

	dt, err := handler.kontenServiceInterface.SerKontenCariBySlug(slug)
	if err != nil {
		log.Println("ERROR SerKontenCariBySlug : " + err.Error())
		http.Error(w, "ERROR SerKontenCariBySlug", http.StatusInternalServerError)
		return
	}

	var vw model.KontenUpdateView
	vw.View = dt.View + 1

	data, err := handler.kontenServiceInterface.SerKontenTambahView(slug, vw)
	if err != nil {
		log.Println("ERROR Views : " + err.Error())
		http.Error(w, "ERROR Views", http.StatusInternalServerError)
		return
	}

	log.Println(data)
	http.Error(w, "Berhasil Views", http.StatusOK)
}
