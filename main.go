package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	// tentukan file
	filename := "faces.jpeg"

	// tentukan haar cascade file
	xmlFile := "data/haarcascades/haarcascade_frontalface_alt2.xml"

	// buat window baru
	window := gocv.NewWindow("Faces Detector")
	defer window.Close()

	// baca file
	img := gocv.IMRead(filename, gocv.IMReadColor)

	// buat warna
	red := color.RGBA{255, 0, 0, 0}

	// load classifier untuk mengenali wajah
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	// kondisikan file haar cascade
	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	// jika gambar matriknya kosong
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return
	}

	// deteksi wajah
	rects := classifier.DetectMultiScale(img)

	// tampilkan jumalah wajah yang terdeteksi
	fmt.Printf("found %d faces\n", len(rects))

	// gambar persegi panjang di sekitar setiap wajah pada gambar asli,
	// bersama dengan teks yang mengidentifikasi sebagai "Wajah {count}"
	count := 1
	for _, r := range rects {
		// buat persegi
		gocv.Rectangle(&img, r, red, 2)

		// buat dan tampilkan teks
		text := fmt.Sprintf("Wajah %d", count)
		pt := image.Pt(r.Min.X, r.Min.Y-5)
		gocv.PutText(&img, text, pt, gocv.FontHersheyPlain, 1.1, red, 1)

		count++
	}

	for {
		// tampilkan image
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
