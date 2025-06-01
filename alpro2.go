package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"strings"
)

// Menyimpan data donatur
type donatur struct {
	id          int     // id unik donatur
	nama        string  // nama donatur
	kotaAsal    string  // kota asal donatur
	jumlah      float64 // jumlah uang yang di donasiin
	sudahDonasi bool    // status donasi (true jika sudah donasi)
	tujuanDonasi string // tempat donasi yang akan di sumbangkan 
}

const MAX_DONATUR int = 1000 // maksimal jumlah donatur

type dataDonatur [MAX_DONATUR]donatur

func main() {
	/* I.S. Program dijalankan dan belum ada interaksi dari pengguna.
	   F.S. Menyajikan menu utama dan memanggil fungsi sesuai pilihan pengguna*/

	var D dataDonatur
	var jumlahDonatur int
	var pilihan int
	
	jumlahDonatur = 4
	
	D[0] = donatur{id: 123, nama:"Anto", kotaAsal:"Palembang", jumlah: 5000, sudahDonasi: true, tujuanDonasi: "Palestina"}
	D[1] = donatur{id: 456, nama:"Budi", kotaAsal:"Medan", jumlah: 15000, sudahDonasi: true, tujuanDonasi: "Yayasan Kanker"}
	D[2] = donatur{id: 789, nama:"Dimas", kotaAsal:"Jakarta", jumlah: 10000, sudahDonasi: false, tujuanDonasi: "Panti Asuhan"}
	D[3] = donatur{id: 147, nama:"Jaki", kotaAsal:"Bandung", jumlah: 20000, sudahDonasi: true, tujuanDonasi: "Palestina"}
	
	for pilihan != 9 {
		menu()
		fmt.Print("Pilih (1-9): ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			tambahDonatur(&D, &jumlahDonatur)
		case 2:
			lihatDonatur(D, jumlahDonatur)
		case 3:
			editDonatur(&D, &jumlahDonatur)
		case 4:
			hapusDonatur(&D, &jumlahDonatur)
		case 5:
			urutData(&D, &jumlahDonatur)
		case 6:
			cariDonaturJumlah(D, jumlahDonatur)
		case 7:
			cariDonaturID(D, jumlahDonatur)
		case 8:
			menuEkstrim(D, jumlahDonatur)
		case 9:
			tampilkanTerimakasih()
		default:
			tampilkanError("Pilihan tidak valid! Masukkan angka 1-9")
		}
		if pilihan != 9 {
			pauseProgram()
		}
	}
}

func menu() {
	clear()
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                         SELAMAT DATANG                                                                         ║")
	fmt.Println("║                                                                   DI APLIKASI REKAP DATA DONASI                                                                ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  Created by:                                                                                                                                                   ║")
	fmt.Println("║  • Achmad Rafi Dwiyandar                                                                                                                                       ║")
	fmt.Println("║  • Sarah Nur Aqilah Tanjung                                                                                                                                    ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                          MENU UTAMA                                                                            ║")
	fmt.Println("╠═══╦════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1 ║ TAMBAH DONATUR                                                                                                                                             ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 2 ║ LIHAT DONATUR                                                                                                                                              ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 3 ║ EDIT DONATUR                                                                                                                                               ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 4 ║ HAPUS DONATUR                                                                                                                                              ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 5 ║ URUT DATA                                                                                                                                                  ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 6 ║ CARI DONATUR BERDASARKAN JUMLAH DONASI                                                                                                                     ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 7 ║ CARI DONATUR BERDASARKAN ID                                                                                                                                ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 8 ║ MENU DONASI EKSTRIM                                                                                                                                        ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 9 ║ EXIT                                                                                                                                                       ║")
	fmt.Println("╚═══╩════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func tambahDonatur(D *dataDonatur, jumlahDonatur *int) {
	/* I.S. Terdefinisi id donatur, nama donatur, kota asal, jumlah donasi, dan status donasi
	   F.S. Donatur baru ditambahkan ke array jika belum penuh */

	// header utama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                        TAMBAH DONATUR BARU                                                                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")

	var donaturBaru donatur

	// cek apakah array sudah penuh
	if *jumlahDonatur >= MAX_DONATUR {
		tampilkanError("Data donatur sudah penuh!")
	}

	// cek ID duplikat
	fmt.Print("Masukkan ID Donatur: ")
	fmt.Scanln(&donaturBaru.id)

	for i := 0; i < *jumlahDonatur; i++ {
		if D[i].id == donaturBaru.id {
			tampilkanError("ID donatur sudah ada! Silakan gunakan ID yang lain.")
			fmt.Print("Masukkan ID Donatur: ")
			fmt.Scanln(&donaturBaru.id)
		}
	}
	
	donaturBaru.nama = scanLineInput("Masukkan Nama Donatur: ")
	donaturBaru.kotaAsal = scanLineInput("Masukkan Kota Asal: ")

	fmt.Print("Masukkan Jumlah Donasi: ")
	fmt.Scanln(&donaturBaru.jumlah)

	fmt.Print("Masukkan Status Donasi (true/false): ")
	fmt.Scanln(&donaturBaru.sudahDonasi)

	donaturBaru.tujuanDonasi = scanLineInput("Masukkan Tujuan Donasi: ")
	
	D[*jumlahDonatur] = donaturBaru
	*jumlahDonatur++
	tampilkanSuccess("Donatur berhasil ditambahkan!")
}

func lihatDonatur(D dataDonatur, jumlahDonatur int) {
	/* I.S. pengguna memilih nomor 2 pada bagian menu
	   F.S. Menampilkan id donatur, nama donatur, kota asal, jumlah donasi, dan status donasi di program */

	// header utama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                   DAFTAR DONATUR                                                                               ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")

	// cek jika belum ada donatur
	if jumlahDonatur == 0 {
		tampilkanError("Belum ada donatur")
	}

	// header tabel
	fmt.Println("╔════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦══════════════════════════════╦═══════════════╦══════════════════════╗")
	fmt.Println("║   ID   ║                  NAMA                     ║                KOTA                 ║            JUMLAH            ║    STATUS     ║       TUJUAN         ║")
	fmt.Println("╠════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬══════════════════════════════╬═══════════════╬══════════════════════╣")

	// data donatur
	for i := 0; i < jumlahDonatur; i++ {
		fmt.Printf("║%-7d ║%-43s║%-37s║ Rp%-27.0f║%-15s║%-22s║\n", 
			D[i].id, 
			truncateString(D[i].nama, 43), 
			truncateString(D[i].kotaAsal, 37), 
			D[i].jumlah, 
			statusDonasi(D[i].sudahDonasi), 
			truncateString(D[i].tujuanDonasi, 22))
	}
	fmt.Println("╚════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩══════════════════════════════╩═══════════════╩══════════════════════╝")

	// tampilkan info donasi
	infoDonasi(D, jumlahDonatur)
}

func statusDonasi(sudah bool) string {
	/* I.S. Terdefinisi nilai boolean status donasi dari fungsi tambahDonatur
	   F.S. Mengembalikan string status donasi di program */

	// pengecekan hasil status donasi
	if sudah == true {
		return "Sudah Donasi"
	}
	return "Belum Donasi"
}

func editDonatur(D *dataDonatur, jumlahDonatur *int) {
	// header utama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                          EDIT DONATUR                                                                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
	var idCari int
	var ditemukan bool = false
	var indeks int = -1
	var dataDiperbarui bool = false
	
	if *jumlahDonatur == 0 {
		tampilkanError("Belum ada donatur")
	}
	
	fmt.Print("Masukkan ID Donatur yang akan diedit: ")
	fmt.Scanln(&idCari)
	
	for i := 0; i < *jumlahDonatur && !ditemukan; i++ {
		if D[i].id == idCari {
			ditemukan = true
			indeks = i
		}
	}
	
	if ditemukan {
		fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                                                   DONATUR DITEMUKAN                                                                        ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
		
		// tabel data donatur yang ditemukan
		fmt.Println("╔════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦══════════════════════════╦═══════════════╦══════════════════════╗")
		fmt.Printf("║ %-6s ║ %-41s ║ %-35s ║ %-24s ║ %-13s ║ %-20s ║\n", "ID", "NAMA", "KOTA", "JUMLAH", "STATUS", "TUJUAN")
		fmt.Println("╠════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬══════════════════════════╬═══════════════╬══════════════════════╣")
		
		// format data
		nama := truncateString(D[indeks].nama, 41)
		kota := truncateString(D[indeks].kotaAsal, 35)
		tujuan := truncateString(D[indeks].tujuanDonasi, 20)
		jumlahStr := formatRupiah(D[indeks].jumlah)
		status := statusDonasi(D[indeks].sudahDonasi)
		
		fmt.Printf("║ %-6d ║ %-41s ║ %-35s ║ %24s ║ %-13s ║ %-20s ║\n", D[indeks].id, nama, kota, jumlahStr, status, tujuan)
		fmt.Println("╚════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩══════════════════════════╩═══════════════╩══════════════════════╝")
		
		var pilihan int
		for pilihan != 7 {
			menuEdit()
			fmt.Print("Pilih (1-7): ")
			fmt.Scanln(&pilihan)
			switch pilihan {
			case 1:
				D[indeks].nama = scanLineInput("Masukkan Nama Baru: ")
				tampilkanSuccess("Nama berhasil diperbarui!")
				dataDiperbarui = true
			case 2:
				D[indeks].kotaAsal = scanLineInput("Masukkan Kota Asal Baru: ")
				tampilkanSuccess("Kota asal berhasil diperbarui!")
				dataDiperbarui = true 
			case 3:
				fmt.Print("Masukkan Jumlah Donasi Baru: ")
				fmt.Scanln(&D[indeks].jumlah)
				tampilkanSuccess("Jumlah donasi berhasil diperbarui!")
				dataDiperbarui = true
			case 4:
				fmt.Print("Masukkan Status Donasi Baru (true/false): ")
				fmt.Scanln(&D[indeks].sudahDonasi)
				tampilkanSuccess("Status donasi berhasil diperbarui!")
				dataDiperbarui = true
			case 5:
				D[indeks].tujuanDonasi = scanLineInput("Masukkan Tujuan Donasi Baru: ")
				tampilkanSuccess("Tujuan donasi berhasil diperbarui!")
				dataDiperbarui = true
			case 6:
				D[indeks].nama = scanLineInput("Masukkan Nama Baru: ")
				D[indeks].kotaAsal = scanLineInput("Masukkan Kota Asal Baru: ")
				fmt.Print("Masukkan Jumlah Donasi Baru: ")
				fmt.Scanln(&D[indeks].jumlah)
				fmt.Print("Masukkan Status Donasi Baru (true/false): ")
				fmt.Scanln(&D[indeks].sudahDonasi)
				D[indeks].tujuanDonasi = scanLineInput("Masukkan Tujuan Donasi Baru: ")
				tampilkanSuccess("Semua data berhasil diperbarui!")
				dataDiperbarui = true
			case 7:
				fmt.Println("Kembali ke menu utama...")
			default:
				tampilkanError("Pilihan tidak valid!")
			}
		}
		
		if dataDiperbarui {
			fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                                                                          DATA TERBARU                                                                          ║")
			fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
			
			// tabel data donatur yang sudah diperbarui
			fmt.Println("\n╔════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦═════════════════════════════╦═══════════════╦═══════════════════════╗")
			fmt.Printf("║ %-6s ║ %-41s ║ %-35s ║ %-27s ║ %-13s ║ %-21s ║\n", "ID", "NAMA", "KOTA", "JUMLAH", "STATUS", "TUJUAN")
			fmt.Println("╠════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬═════════════════════════════╬═══════════════╬═══════════════════════╣")
			
			// format
			namaNew := truncateString(D[indeks].nama, 41)
			kotaNew := truncateString(D[indeks].kotaAsal, 35)
			tujuanNew := truncateString(D[indeks].tujuanDonasi, 21)
			jumlahStrNew := formatRupiah(D[indeks].jumlah)
			statusNew := statusDonasi(D[indeks].sudahDonasi)
			
			fmt.Printf("║ %-6d ║ %-41s ║ %-35s ║ %27s ║ %-13s ║ %-21s ║\n", 
				D[indeks].id, 
				namaNew, 
				kotaNew, 
				jumlahStrNew, 
				statusNew, 
				tujuanNew)
			fmt.Println("╚════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩═════════════════════════════╩═══════════════╩═══════════════════════╝")
		}
	} else {
		tampilkanError("Donatur dengan ID tersebut tidak ditemukan!")
	}
}

func menuEdit() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                          EDIT DATA                                                                             ║")
	fmt.Println("╠═══╦════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1 ║ Edit Nama Donatur                                                                                                                                          ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 2 ║ Edit Kota Asal                                                                                                                                             ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 3 ║ Edit Jumlah Donasi                                                                                                                                         ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 4 ║ Edit Status Donasi                                                                                                                                         ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 5 ║ Edit Tujuan Donasi                                                                                                                                         ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 6 ║ Edit Semua Data                                                                                                                                            ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 7 ║ Kembali                                                                                                                                                    ║")
	fmt.Println("╚═══╩════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

func hapusDonatur(D *dataDonatur, jumlahDonatur *int) {
	/* I.S. Id donatur, nama donatur, jumlah donasi, dan status donasi dihapus
	   F.S. Data yang tidak diperlukan, tidak tersedia */

	// header utama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                        HAPUS DONATUR                                                                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")

	var idHapus int
	var ditemukan bool = false
	var indeks int = -1

	// cek jika belum ada donatur
	if *jumlahDonatur == 0 {
		tampilkanError("Belum ada donatur")
	}

	// meminta ID donatur yang akan dihapus
	fmt.Print("Masukkan ID Donatur yang akan dihapus: ")
	fmt.Scanln(&idHapus)

	// mencari donatur dengan sequential search
	for i := 0; i < *jumlahDonatur && !ditemukan; i++ {
		if D[i].id == idHapus {
			ditemukan = true
			indeks = i
		}
	}

	// jika donatur ditemukan, konfirmasi penghapusan
	if ditemukan {
		var konfirmasi string

		fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                                                       DONATUR DITEMUKAN                                                                        ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
		fmt.Printf("ID: %d | Nama: %s | Kota: %s | Jumlah: Rp%.0f | Status: %s | Tujuan: %s\n", 
			D[indeks].id, D[indeks].nama, D[indeks].kotaAsal, D[indeks].jumlah, 
			statusDonasi(D[indeks].sudahDonasi), D[indeks].tujuanDonasi)

		fmt.Print("Apakah Anda yakin ingin menghapus data ini? (y/n): ")
		fmt.Scanln(&konfirmasi)

		if konfirmasi == "y" || konfirmasi == "Y" {
			// menghapus dengan menggeser semua elemen setelahnya
			for i := indeks; i < *jumlahDonatur-1; i++ {
				D[i] = D[i+1]
			}
			// mengurangi jumlah donatur
			*jumlahDonatur--
			tampilkanSuccess("Data donatur berhasil dihapus!")
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	} else {
		tampilkanError("Donatur dengan ID tersebut tidak ditemukan!")
	}
}

func urutData(D *dataDonatur, jumlahDonatur *int) {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                           URUT DATA                                                                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")

	var pilihan int

	if *jumlahDonatur == 0 {
		tampilkanError("Belum ada donatur")
	}

	for pilihan != 5 {
		menuMengurut()
		fmt.Print("Pilih (1-5): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			urutJumlahInsertionSortAsce(D, *jumlahDonatur)
			tampilkanSuccess("Data jumlah donasi berhasil diurutkan dari terkecil ke terbesar")
			lihatDonatur(*D, *jumlahDonatur)
		case 2:
			urutJumlahInsertionSortDesce(D, *jumlahDonatur)
			tampilkanSuccess("Data jumlah donasi berhasil diurutkan dari terbesar ke terkecil")
			lihatDonatur(*D, *jumlahDonatur)
		case 3:
			urutIDSelectionSortAsce(D, *jumlahDonatur)
			tampilkanSuccess("Data ID donatur berhasil diurutkan dari terkecil ke terbesar")
			lihatDonatur(*D, *jumlahDonatur)
		case 4:
			urutIDSelectionSortDesce(D, *jumlahDonatur)
			tampilkanSuccess("Data ID donatur berhasil diurutkan dari terbesar ke terkecil")
			lihatDonatur(*D, *jumlahDonatur)
		case 5:
			fmt.Println("Kembali ke menu utama")
		default:
			tampilkanError("Pilihan tidak valid!")
		}
		if pilihan != 5 {
			pauseProgram()
		}
	}
}

func menuMengurut() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                            URUT DATA                                                                           ║")
	fmt.Println("╠═══╦════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1 ║ Urut Jumlah Donasi (Terkecil → Terbesar)                                                                                                                   ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 2 ║ Urut Jumlah Donasi (Terbesar → Terkecil)                                                                                                                   ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 3 ║ Urut ID Donatur (Terkecil → Terbesar)                                                                                                                      ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 4 ║ Urut ID Donatur (Terbesar → Terkecil)                                                                                                                      ║")
	fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 5 ║ Kembali ke Menu Utama                                                                                                                                      ║")
	fmt.Println("╚═══╩════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// insertion sort berdasarkan jumlah donasi (ascending)
func urutJumlahInsertionSortAsce(D *dataDonatur, jumlahDonatur int) {
	var i, j int
	var temp donatur

	for i = 1; i < jumlahDonatur; i++ {
		temp = D[i]
		j = i - 1
		
		for j >= 0 && D[j].jumlah > temp.jumlah {
			D[j+1] = D[j]
			j--
		}
		D[j+1] = temp
	}
}

// insertion sort berdasarkan jumlah donasi (descending)
func urutJumlahInsertionSortDesce(D *dataDonatur, jumlahDonatur int) {
	var i, j int
	var temp donatur

	for i = 1; i < jumlahDonatur; i++ {
		temp = D[i]
		j = i - 1
		
		for j >= 0 && D[j].jumlah < temp.jumlah {
			D[j+1] = D[j]
			j--
		}
		D[j+1] = temp
	}
}

// selection sort berdasarkan ID (ascending)
func urutIDSelectionSortAsce(D *dataDonatur, jumlahDonatur int) {
	var i, idx, pass int
	var temp donatur

	pass = 1
	for pass < jumlahDonatur {
		idx = pass - 1
		i = pass
		for i < jumlahDonatur {
			if D[i].id < D[idx].id {
				idx = i
			}
			i++
		}
		temp = D[pass-1]
		D[pass-1] = D[idx]
		D[idx] = temp
		pass++
	}
}

// selection sort berdasarkan ID (descending)
func urutIDSelectionSortDesce(D *dataDonatur, jumlahDonatur int) {
	var i, idx, pass int
	var temp donatur

	pass = 1
	for pass < jumlahDonatur {
		idx = pass - 1
		i = pass
		for i < jumlahDonatur {
			if D[i].id > D[idx].id {
				idx = i
			}
			i++
		}
		temp = D[pass-1]
		D[pass-1] = D[idx]
		D[idx] = temp
		pass++
	}
}

// Sequential search berdasarkan jumlah donasi
func cariDonaturJumlah(D dataDonatur, jumlahDonatur int) {
	//header uatama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                             CARI DONATUR BERDASARKAN JUMLAH DONASI                                                             ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
	
	var jumlahCari float64
	var i int
	var ditemukan bool = false
	
	if jumlahDonatur == 0 {
		tampilkanError("Belum ada donatur")
	}
	
	fmt.Print("Masukkan jumlah donasi yang dicari: ")
	fmt.Scanln(&jumlahCari)
	
	fmt.Printf("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                                                                         HASIL PENCARIAN                                                                        ║\n")
	fmt.Printf("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝\n")
	
	// Sequential search
	for i = 0; i < jumlahDonatur; i++ {
		if D[i].jumlah == jumlahCari {
			if !ditemukan {
				fmt.Println("╔═══════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦═════════════════════════╦═══════════════╦════════════════════════╗")
				fmt.Println("║     ID    ║                   NAMA                    ║                KOTA                 ║         JUMLAH          ║    STATUS     ║         TUJUAN         ║")
				fmt.Println("╠═══════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬═════════════════════════╬═══════════════╬════════════════════════╣")
			}
			
			// format
			jumlahStr := fmt.Sprintf("Rp%.0f", D[i].jumlah)
			
			fmt.Printf("║%-11d║%-43s║%-37s║%-25s║%-15s║%-24s║\n",
				D[i].id,
				truncateString(D[i].nama, 43),
				truncateString(D[i].kotaAsal, 37),
				jumlahStr,
				statusDonasi(D[i].sudahDonasi),
				truncateString(D[i].tujuanDonasi, 24))
			ditemukan = true
		}
	}
	
	if ditemukan {
		fmt.Println("╚═══════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩═════════════════════════╩═══════════════╩════════════════════════╝")
	} else {
		tampilkanError(fmt.Sprintf("Donatur dengan jumlah donasi Rp%.0f tidak ditemukan", jumlahCari))
	}
}

// Fungsi helper untuk memformat angka dengan ribuan (opsional)
func formatRupiah(jumlah float64) string {
	str := fmt.Sprintf("%.0f", jumlah)
	
	if len(str) > 3 {
		var result string
		for i := 0; i < len(str); i++ {
			if i > 0 && (len(str)-i)%3 == 0 {
				result += "."
			}
			result += string(str[i])
		}
		return "Rp" + result
	}
	return "Rp" + str
}

// Binary search berdasarkan ID (data harus sudah terurut)
func cariDonaturID(D dataDonatur, jumlahDonatur int) {
	// header utama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                   CARI DONATUR BERDASARKAN ID                                                                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")

	var idCari int

	if jumlahDonatur == 0 {
		tampilkanError("Belum ada donatur")
	}

	fmt.Print("Masukkan ID donatur yang dicari: ")
	fmt.Scanln(&idCari)

	// Salin data untuk pengurutan sementara
	var tempData dataDonatur
	for i := 0; i < jumlahDonatur; i++ {
		tempData[i] = D[i]
	}

	// Urutkan data berdasarkan ID terlebih dahulu menggunakan selection sort
	urutIDSelectionSortAsce(&tempData, jumlahDonatur)

	// Binary search
	indeks := binarySearchID(tempData, jumlahDonatur, idCari)

	fmt.Printf("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                                                                         HASIL PENCARIAN                                                                        ║\n")
	fmt.Printf("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝\n")

	if indeks != -1 {
		fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                                                               DONATUR DITEMUKAN                                                                ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
		
		// header tabel
		fmt.Println("╔═══════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦═════════════════════════╦═══════════════╦════════════════════════╗")
		fmt.Println("║    ID     ║                   NAMA                    ║                KOTA                 ║         JUMLAH          ║    STATUS     ║         TUJUAN         ║")
		fmt.Println("╠═══════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬═════════════════════════╬═══════════════╬════════════════════════╣")
	
		// format
		id := fmt.Sprintf("%d", tempData[indeks].id)
		nama := truncateString(tempData[indeks].nama, 41)
		kota := truncateString(tempData[indeks].kotaAsal, 35)
		jumlah := fmt.Sprintf("Rp%.0f", tempData[indeks].jumlah)
		status := statusDonasi(tempData[indeks].sudahDonasi)
		tujuan := truncateString(tempData[indeks].tujuanDonasi, 22)
		
		// pastikan setiap field memiliki lebar yang tepat
		fmt.Printf("║ %-9s ║ %-41s ║ %-35s ║ %-23s ║ %-13s ║ %-22s ║\n", 
			id, nama, kota, jumlah, status, tujuan)
		
		fmt.Println("╚═══════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩═════════════════════════╩═══════════════╩════════════════════════╝")
	} else {
		// donatur tidak ditemukan
		fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                                                    DONATUR TIDAK DITEMUKAN                                                                     ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
		fmt.Printf("Donatur dengan ID %d tidak ditemukan dalam database\n", idCari)
		fmt.Println("Pastikan ID yang dimasukkan sudah benar")
		
		// saran pencarian
		fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                                                      SARAN PENCARIAN                                                                           ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ • Periksa kembali ID yang dimasukkan                                                                                                                           ║")
		fmt.Println("║ • Gunakan menu 'LIHAT DONATUR' untuk melihat semua ID yang tersedia                                                                                            ║")
		fmt.Println("║ • Pastikan donatur sudah terdaftar dalam sistem                                                                                                                ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
	}
}

// Fungsi binary search untuk ID
func binarySearchID(D dataDonatur, n int, id int) int {
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2

		if D[mid].id == id {
			return mid
		} else if D[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // tidak ditemukan
}

// Menu untuk pencarian nilai ekstrim
func menuEkstrim(D dataDonatur, jumlahDonatur int) {
	// header utama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                       MENU DONASI EKSTRIM                                                                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")

	var pilihan int

	if jumlahDonatur == 0 {
		tampilkanError("Belum ada donatur")
	}

	for pilihan != 4 {
		fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                                                       DONASI EKSTRIM                                                                           ║")
		fmt.Println("╠═══╦════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1 ║ Tampilkan Donasi TERBESAR                                                                                                                                  ║")
		fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 2 ║ Tampilkan Donasi TERKECIL                                                                                                                                  ║")
		fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 3 ║ Tampilkan KEDUANYA                                                                                                                                         ║")
		fmt.Println("╠═══╬════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 4 ║ Kembali ke Menu Utama                                                                                                                                      ║")
		fmt.Println("╚═══╩════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
		fmt.Println()
		
		fmt.Print("Pilih (1-4): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanDonasiTerbesar(D, jumlahDonatur)
		case 2:
			tampilkanDonasiTerkecil(D, jumlahDonatur)
		case 3:
			tampilkanEkstrim(D, jumlahDonatur)
		case 4:
			fmt.Println("Kembali ke menu utama.")
		default:
			tampilkanError("Pilihan tidak valid! Masukkan angka 1-4")
		}
		
		if pilihan != 4 {
			pauseProgram()
			fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                                                                      MENU DONASI EKSTRIM                                                                       ║")
			fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
		}
	}
}

// Menampilkan donatur dengan donasi terbesar
func tampilkanDonasiTerbesar(D dataDonatur, jumlahDonatur int) {
	nilaiTerbesar := cariDonasiTerbesar(D, jumlahDonatur)
	
	// header dengan donasi terbesar
	fmt.Printf("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                                                              DONATUR DENGAN DONASI TERBESAR                                                                    ║\n")
	donasiFmt := formatRupiah(nilaiTerbesar)
	// padding manual untuk center alignment
	paddingKiri := (160 - len(donasiFmt)) / 2
	paddingKanan := 160 - len(donasiFmt) - paddingKiri
	fmt.Printf("║")
	for i := 0; i < paddingKiri; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%s", donasiFmt)
	for i := 0; i < paddingKanan; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	fmt.Printf("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝\n\n")
	
	// header tabel 
	fmt.Println("╔══════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦══════════════════════════╦═══════════════╦════════════════════════╗")
	fmt.Printf("║ %-8s ║ %-41s ║ %-35s ║ %-24s ║ %-13s ║ %-22s ║\n", "ID", "NAMA", "KOTA", "JUMLAH", "STATUS", "TUJUAN")
	fmt.Println("╠══════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬══════════════════════════╬═══════════════╬════════════════════════╣")
	
	// data donatur dengan donasi terbesar
	jumlahData := 0
	for i := 0; i < jumlahDonatur; i++ {
		if D[i].jumlah == nilaiTerbesar {
			// format data
			nama := truncateString(D[i].nama, 41)
			kota := truncateString(D[i].kotaAsal, 35)
			tujuan := truncateString(D[i].tujuanDonasi, 22)
			jumlahStr := formatRupiah(D[i].jumlah)
			status := statusDonasi(D[i].sudahDonasi)
			
			// format dengan spacing yang tepat sesuai header
			fmt.Printf("║ %-8d ║ %-41s ║ %-35s ║ %24s ║ %-13s ║ %-22s ║\n", 
				D[i].id, 
				nama, 
				kota, 
				jumlahStr, 
				status, 
				tujuan)
			jumlahData++
		}
	}
	fmt.Println("╚══════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩══════════════════════════╩═══════════════╩════════════════════════╝")

	// data statistik
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                         STATISTIK                                                                              ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	
	// format statistik 
	totalText := fmt.Sprintf("Total donatur dengan donasi terbesar: %d orang", jumlahData)
	totalSpaces := 160 - len(totalText)
	fmt.Printf("║ %s", totalText)
	for i := 0; i < totalSpaces-1; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	
	nilaiText := fmt.Sprintf("Nilai donasi terbesar: %s", formatRupiah(nilaiTerbesar))
	nilaiSpaces := 160 - len(nilaiText)
	fmt.Printf("║ %s", nilaiText)
	for i := 0; i < nilaiSpaces-1; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}
// Menampilkan donatur dengan donasi terkecil
func tampilkanDonasiTerkecil(D dataDonatur, jumlahDonatur int) {
	nilaiTerkecil := cariDonasiTerkecil(D, jumlahDonatur)
	
	// header dengan donasi terkecil
	fmt.Printf("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                                                              DONATUR DENGAN DONASI TERKECIL                                                                    ║\n")
	donasiFmt := formatRupiah(nilaiTerkecil)
	// padding manual untuk center alignment
	paddingKiri := (160 - len(donasiFmt)) / 2
	paddingKanan := 160 - len(donasiFmt) - paddingKiri
	fmt.Printf("║")
	for i := 0; i < paddingKiri; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%s", donasiFmt)
	for i := 0; i < paddingKanan; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	fmt.Printf("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝\n\n")
	
	// header tabel
	fmt.Println("╔══════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦══════════════════════════╦═══════════════╦════════════════════════╗")
	fmt.Printf("║ %-8s ║ %-41s ║ %-35s ║ %-24s ║ %-13s ║ %-22s ║\n", "ID", "NAMA", "KOTA", "JUMLAH", "STATUS", "TUJUAN")
	fmt.Println("╠══════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬══════════════════════════╬═══════════════╬════════════════════════╣")
	
	// data donatur dengan donasi terkecil
	jumlahData := 0
	for i := 0; i < jumlahDonatur; i++ {
		if D[i].jumlah == nilaiTerkecil {
			// Format data
			nama := truncateString(D[i].nama, 41)
			kota := truncateString(D[i].kotaAsal, 35)
			tujuan := truncateString(D[i].tujuanDonasi, 22)
			jumlahStr := formatRupiah(D[i].jumlah)
			status := statusDonasi(D[i].sudahDonasi)
			
			// Format dengan spacing
			fmt.Printf("║ %-8d ║ %-41s ║ %-35s ║ %24s ║ %-13s ║ %-22s ║\n", 
				D[i].id, 
				nama, 
				kota, 
				jumlahStr, 
				status, 
				tujuan)
			jumlahData++
		}
	}
	fmt.Println("╚══════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩══════════════════════════╩═══════════════╩════════════════════════╝")

	// data statistik
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                         STATISTIK                                                                              ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	
	// Format statistik dengan padding manual untuk 91 karakter
	totalText := fmt.Sprintf("Total donatur dengan donasi terbesar: %d orang", jumlahData)
	totalSpaces := 160 - len(totalText)
	fmt.Printf("║ %s", totalText)
	for i := 0; i < totalSpaces-1; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	
	nilaiText := fmt.Sprintf("Nilai donasi terkecil: %s", formatRupiah(nilaiTerkecil))
	nilaiSpaces := 160 - len(nilaiText)
	fmt.Printf("║ %s", nilaiText)
	for i := 0; i < nilaiSpaces-1; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// Pencarian nilai ekstrim (donasi terbesar dan terkecil) - tampilkan keduanya
func tampilkanEkstrim(D dataDonatur, jumlahDonatur int) {
	// header utama
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                        DONASI EKSTRIM                                                                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")

	nilaiTerbesar := cariDonasiTerbesar(D, jumlahDonatur)
	nilaiTerkecil := cariDonasiTerkecil(D, jumlahDonatur)

	// DONASI TERBESAR
	fmt.Printf("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                                                                   DONATUR DENGAN DONASI TERBESAR                                                               ║\n")
	donasiFmt := formatRupiah(nilaiTerbesar)
	// padding manual untuk center alignment
	paddingKiri := (160 - len(donasiFmt)) / 2
	paddingKanan := 160 - len(donasiFmt) - paddingKiri
	fmt.Printf("║")
	for i := 0; i < paddingKiri; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%s", donasiFmt)
	for i := 0; i < paddingKanan; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	fmt.Printf("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝\n\n")
	
	// header tabel
	fmt.Println("╔══════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦══════════════════════════╦═══════════════╦════════════════════════╗")
	fmt.Printf("║ %-8s ║ %-41s ║ %-35s ║ %-24s ║ %-13s ║ %-22s ║\n", "ID", "NAMA", "KOTA", "JUMLAH", "STATUS", "TUJUAN")
	fmt.Println("╠══════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬══════════════════════════╬═══════════════╬════════════════════════╣")
	
	// Data donatur dengan donasi terbesar
	jumlahDataBesar := 0
	for i := 0; i < jumlahDonatur; i++ {
		if D[i].jumlah == nilaiTerbesar {
			// format data
			nama := truncateString(D[i].nama, 41)
			kota := truncateString(D[i].kotaAsal, 35)
			tujuan := truncateString(D[i].tujuanDonasi, 22)
			jumlahStr := formatRupiah(D[i].jumlah)
			status := statusDonasi(D[i].sudahDonasi)
			
			// format dengan spacing
			fmt.Printf("║ %-8d ║ %-41s ║ %-35s ║ %24s ║ %-13s ║ %-22s ║\n", 
				D[i].id, 
				nama, 
				kota, 
				jumlahStr, 
				status, 
				tujuan)
			jumlahDataBesar++
		}
	}
	fmt.Println("╚══════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩══════════════════════════╩═══════════════╩════════════════════════╝")

	// DONASI TERKECIL
	fmt.Printf("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                                                                   DONATUR DENGAN DONASI TERKECIL                                                               ║\n")
	donasiFmtKecil := formatRupiah(nilaiTerkecil)
	// padding manual untuk center alignment
	paddingKiriKecil := (160 - len(donasiFmtKecil)) / 2
	paddingKananKecil := 160 - len(donasiFmtKecil) - paddingKiriKecil
	fmt.Printf("║")
	for j := 0; j < paddingKiriKecil; j++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%s", donasiFmtKecil)
	for j := 0; j < paddingKananKecil; j++ {
		fmt.Printf(" ")
	}
	fmt.Printf("║\n")
	fmt.Printf("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝\n\n")
	
	// header tabel
	fmt.Println("╔══════════╦═══════════════════════════════════════════╦═════════════════════════════════════╦══════════════════════════╦═══════════════╦════════════════════════╗")
	fmt.Printf("║ %-8s ║ %-41s ║ %-35s ║ %-24s ║ %-13s ║ %-22s ║\n", "ID", "NAMA", "KOTA", "JUMLAH", "STATUS", "TUJUAN")
	fmt.Println("╠══════════╬═══════════════════════════════════════════╬═════════════════════════════════════╬══════════════════════════╬═══════════════╬════════════════════════╣")
	
	// Data donatur dengan donasi terkecil
	jumlahDataKecil := 0
	for k := 0; k < jumlahDonatur; k++ {
		if D[k].jumlah == nilaiTerkecil {
			// Format data dengan padding yang konsisten
			nama := truncateString(D[k].nama, 41)
			kota := truncateString(D[k].kotaAsal, 35)
			tujuan := truncateString(D[k].tujuanDonasi, 22)
			jumlahStr := formatRupiah(D[k].jumlah)
			status := statusDonasi(D[k].sudahDonasi)
			
			// Format dengan spacing
			fmt.Printf("║ %-8d ║ %-41s ║ %-35s ║ %24s ║ %-13s ║ %-22s ║\n", 
				D[k].id, 
				nama, 
				kota, 
				jumlahStr, 
				status, 
				tujuan)
			jumlahDataKecil++
		}
	}
	fmt.Println("╚══════════╩═══════════════════════════════════════════╩═════════════════════════════════════╩══════════════════════════╩═══════════════╩════════════════════════╝")

	// STATISTIK GABUNGAN
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║%-160s║\n", "                                                STATISTIK EKSTRIM")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	
	// Format statistik
	textBesar := fmt.Sprintf("Donatur dengan donasi terbesar: %d orang", jumlahDataBesar)
	fmt.Printf("║ %-159s║\n", textBesar)
	
	textKecil := fmt.Sprintf("Donatur dengan donasi terkecil: %d orang", jumlahDataKecil)
	fmt.Printf("║ %-159s║\n", textKecil)
	
	selisih := nilaiTerbesar - nilaiTerkecil
	textSelisih := fmt.Sprintf("Selisih donasi: %s", formatRupiah(selisih))
	fmt.Printf("║ %-159s║\n", textSelisih)
	
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// Fungsi untuk mencari indeks donasi terbesar
func cariDonasiTerbesar(D dataDonatur, jumlahDonatur int) float64 {
	indeksTerbesar := D[0].jumlah

	for i := 1; i < jumlahDonatur; i++ {
		if D[i].jumlah > indeksTerbesar {
			indeksTerbesar = D[i].jumlah
		}
	}
	return indeksTerbesar
}

// Fungsi untuk mencari indeks donasi terkecil
func cariDonasiTerkecil(D dataDonatur, jumlahDonatur int) float64 {
	indeksTerkecil := D[0].jumlah

	for i := 1; i < jumlahDonatur; i++ {
		if D[i].jumlah < indeksTerkecil {
			indeksTerkecil = D[i].jumlah
		}
	}
	return indeksTerkecil
}

// procedure untuk menampilkan info donasi seperti total, rata-rata dll
func infoDonasi(D dataDonatur, jumlahDonatur int) {
	var totalDonasi float64 = 0
	var sudahDonasi int = 0
	var belumDonasi int = 0

	// hitung statistik donasi
	for i := 0; i < jumlahDonatur; i++ {
		totalDonasi += D[i].jumlah
		if D[i].sudahDonasi {
			sudahDonasi++
		} else {
			belumDonasi++
		}
	}

	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                          INFO DONASI                                                                           ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ Total Donatur        : %-4d orang                                                                                                                              ║\n", jumlahDonatur)
	fmt.Printf("║ Sudah Donasi         : %-4d orang                                                                                                                              ║\n", sudahDonasi)
	fmt.Printf("║ Belum Donasi         : %-4d orang                                                                                                                              ║\n", belumDonasi)
	fmt.Printf("║ Total Dana Terkumpul : Rp%-134.0f║\n", totalDonasi)

	if jumlahDonatur > 0 {
		fmt.Printf("║ Rata-rata Donasi     : Rp%-134.0f║\n", totalDonasi/float64(jumlahDonatur))
	}
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// procedure untuk membersihkan tampilan console terminal
func clear() {
	cmd := exec.Command("cmd", "/c", "cls") 
	cmd.Stdout = os.Stdout 
	cmd.Run() 
}

// procedure untuk menjeda program sampai pengguna menekan tombol enter
func pauseProgram() {
	fmt.Println("Tekan ENTER untuk melanjutkan.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// fungsi untuk user agar bisa menggunakan spasi saat pengimputan
func scanLineInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi helper untuk memotong string jika terlalu panjang
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// Fungsi untuk menampilkan pesan sukses
func tampilkanSuccess(pesan string) {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║%s║\n", centerString("✓ SUCCESS", 162))
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %s%s║\n", pesan, strings.Repeat(" ", 159-len(pesan)))
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// Fungsi untuk menampilkan pesan error
func tampilkanError(pesan string) {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║%s║\n", centerString("✗ ERROR", 162))
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %s%s║\n", pesan, strings.Repeat(" ", 159-len(pesan)))
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// Fungsi untuk menampilkan ucapan terima kasih saat exit
func tampilkanTerimakasih() {
	clear()
	fmt.Println("\n╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                     TERIMA KASIH                                                                              ║")
	fmt.Println("╠═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                             Terima kasih telah menggunakan                                                                    ║")
	fmt.Println("║                                                               APLIKASI REKAP DATA DONASI                                                                      ║")
	fmt.Println("║                                                                                                                                                               ║")
	fmt.Println("║                                                             Semoga donasi Anda bermanfaat                                                                     ║")
	fmt.Println("║                                                                kepada yang membutuhkan                                                                        ║")
	fmt.Println("╠═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  Created by:                                                                                                                                                  ║")
	fmt.Println("║  • Achmad Rafi Dwiyandar                                                                                                                                      ║")
	fmt.Println("║  • Sarah Nur Aqilah Tanjung                                                                                                                                   ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

// Fungsi helper untuk meratakan teks di tengah
func centerString(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}
	padding := (width - len(s)) / 2
	leftPad := strings.Repeat(" ", padding)
	rightPad := strings.Repeat(" ", width-len(s)-padding)
	return leftPad + s + rightPad
}