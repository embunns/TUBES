/*

RULES:
- Pelanggan dapat membeli/ mengganti sparepart maksimal 10 sparepart
- Pelanggan diasumsikan hanya datang maksimal sekali dalam sehari, didasarkan pada lamanya antrian dan lamanya waktu service
  (sehingga nama pelanggan dan waktu transaksi (tanggal, bulan, tahun) bisa dijadikan primary key)

SPESIFIKASI APLIKASI :
1. Aplikasi dapat melakukan penambahan data
	1.1. Penambahan nama, tanggal transaksi, bulan transaksi, tahun transaksi, dan sparepart yang dibeli
	1.2. Penambahan sparepart saja (karena sparepart yang dibeli bisa lebih dari 1)
	NOTES: TIDAK ADA PENAMBAHAN NAMA SAJA ATAU WAKTU TRANSAKSI SAJA KARENA DAPAT MENYEBABKAN VARIABEL LAIN BERNILAI NULL
2. Aplikasi dapat melakukan pengubahan data (yang diubah 1 transaksi)
3. Aplikasi dapat melakukan penghapusan data (penghapusan 1 transaksi)
	NOTES: PENGHAPUSAN DATA TIDAK DAPAT DILAKUKAN PADA SATU ATAU BEBERAPA VARIABEL SAJA KARENA DAPAT MENYEBABKAN VARIABEL BERNILAI NULL
4. Aplikasi dapat melakukan penghitungan tarif secara otomatis untuk tiap transaksi (ditampilkan di invoice)
5. Aplikasi dapat melakukan pencarian data
	5.1. Daftar pelanggan berdasarkan waktu transaksi
	5.2. Daftar pelanggan berdasarkan sparepart yang dibeli
6. Aplikasi dapat menampilkan data
	6.1. Menampilkan invoice milik pelanggan dalam sekali transaksi
	6.2. Menampilkan daftar sparepart terurut berdasarkan jumlah yang paling sering diganti

*/

package main

import "fmt"

const NMAX = 10000
const NMAXsparepart = 10

// Tipe bentukan yang akan digunakan untuk menghitung keseluruhan sparepat yang dibeli oleh seluruh pelanggan //
type sparepartglobal struct {
	nama string
	n    int
}

type sparepartcust struct {
	nama   [NMAXsparepart]string
	banyak int
}

type transaksi struct {
	tarif int
	waktu time
	nama  string
	spr   sparepartcust
}

type time struct {
	tanggal, bulan, tahun int
}

type tabtransaksi [NMAX]transaksi

type tabsparepart [NMAX]sparepartcust

type tabsparepartglobal [9]sparepartglobal

type pembelisparepart [NMAX]string

func main() {

	mulai()
	menu()
}

func mulai() {
	var blank string

	fmt.Println("===========================================================")
	fmt.Println("=                 APLIKASI SERVICE MOTOR                  =")
	fmt.Println("=                     Selamat Datang                      =")
	fmt.Println("===========================================================")

	fmt.Print("Tekan tombol apapun untuk lanjut...")
	fmt.Scanln(&blank)
	fmt.Println("\n\n")
}

func menu() {
	var n int
	var t tabtransaksi
	var iTransaksi, iSparepart int
	var p pembelisparepart
	// iTransaksi = iterasi data transaksi//
	var s tabsparepartglobal

	s[0].nama = "velg"
	s[1].nama = "kampas_rem"
	s[2].nama = "karburator"
	s[3].nama = "piston"
	s[4].nama = "gear_depan"
	s[5].nama = "conron"
	s[6].nama = "pullstart"
	s[7].nama = "ban_luar"
	s[8].nama = "ban_dalam"
	for i := 0; i < 9; i++ {
		s[i].n = 0
	}

	iTransaksi = 0
	for n != 6 {
		fmt.Println("===========================================================")
		fmt.Println("=                 APLIKASI SERVICE MOTOR                  =")
		fmt.Println("=                       Menu Utama                        =")
		fmt.Println("===========================================================")
		fmt.Println("=        1. Penambahan Data                               =")
		fmt.Println("=        2. Pengubahan Data                               =")
		fmt.Println("=        3. Penghapusan Data                              =")
		fmt.Println("=        4. Pencarian Data                                =")
		fmt.Println("=        5. Menampilkan Data                              =")
		fmt.Println("=        6. Keluar                                        =")
		fmt.Println("= ------------------------------------------------------- =")
		fmt.Print("         Masukkan pilihan anda : ")
		fmt.Scan(&n)
		fmt.Println("===========================================================")
		fmt.Print("\n\n")
		if n == 1 {
			header_tambah_data()
			tambah_data(&t, &iTransaksi, &s, &iSparepart)
		} else if n == 2 {
			ubah_data(&t, iTransaksi, &s, &iSparepart)
		} else if n == 3 {
			hapus_data(&t, &iTransaksi, &s)
		} else if n == 4 {
			cari_data(t, iTransaksi, iSparepart, s, &p)
		} else if n == 5 {
			tampil_data(s, &t, &iTransaksi)
		} else if n == 6 {

		} else {
			fmt.Println("   Pilihan Anda tidak valid. Silakan pilih opsi yang tersedia!")
		}
	}
	fmt.Println("")
	fmt.Println("===========================================================")
	fmt.Println("=              Terima kasih telah menggunakan             =")
	fmt.Println("=                  APLIKASI SERVICE MOTOR                 =")
	fmt.Println("===========================================================")
	fmt.Println("")

}

func header_tambah_data() {
	fmt.Println("===========================================================")
	fmt.Println("=                  APLIKASI SERVICE MOTOR                 =")
	fmt.Println("=                   Page Penambahan Data                  =")
	fmt.Println("===========================================================")
}

func tambah_data(transaction *tabtransaksi, iTransaksi *int, s *tabsparepartglobal, iSparepart *int) {

	var namasparepart, nama string
	var pilih, tgl, bln, thn int

	fmt.Println("=             1. Menambah seluruh data                    =")
	fmt.Println("=             2. Menambah sparepaart saja                 =")
	fmt.Println("-----------------------------------------------------------")
	fmt.Print("            Masukkan pilihan Anda : ")
	fmt.Scan(&pilih)

	if pilih == 1 {

		transaction[*iTransaksi].tarif = 0

		fmt.Print("  Masukkan nama customer (tanpa spsasi)   : ")
		fmt.Scan(&transaction[*iTransaksi].nama)
		// Satu customer diasumsikan maksimaml hanya melakukan sekali transaksi dalam sehari berdaasarkan estimasi waktu antrian dan lamanya service //
		fmt.Print("  Masukkan tanggal transaksi (contoh: 02) : ")
		fmt.Scan(&transaction[*iTransaksi].waktu.tanggal)
		fmt.Print("  Masukkan bulan transaksi (contoh: 12)   : ")
		fmt.Scan(&transaction[*iTransaksi].waktu.bulan)
		fmt.Print("  Masukkan tahun transaksi (contoh: 2023) : ")
		fmt.Scan(&transaction[*iTransaksi].waktu.tahun)

		fmt.Println("___________________________________________________________")
		fmt.Println("|                                                         |")
		fmt.Println("|    Daftar sparepart yang tersedia beserta hargannya :   |")
		fmt.Println("|             1. velg            Rp 200000                |")
		fmt.Println("|             2. kampas_rem      Rp 50000                 |")
		fmt.Println("|             3. karburator      Rp 350000                |")
		fmt.Println("|             4. piston          Rp 150000                |")
		fmt.Println("|             5. gear_depan      Rp 80000                 |")
		fmt.Println("|             6. conron          Rp 150000                |")
		fmt.Println("|             7. pullstart       Rp 230000                |")
		fmt.Println("|             8. ban_luar        Rp 170000                |")
		fmt.Println("|             9. ban_dalam       Rp 75000                 |")
		fmt.Println("|_________________________________________________________|")
		fmt.Println("  MASUKKAN NAMA-NAMA SPAREPART YANG DIBELI BERDASARKAN SPAREPART YANG TERSEDIA!")
		fmt.Println("  Masukkan 'SELESAI' tanpa tanda petik apabila semua sparepart sudah dimasukkan")
		fmt.Print("  Masukkan nama sparepart ke-1 : ")
		fmt.Scan(&namasparepart)

		i := 0
		for i < NMAXsparepart-1 && namasparepart != "SELESAI" {
			for namasparepart != "velg" && namasparepart != "kampas_rem" && namasparepart != "karburator" && namasparepart != "piston" && namasparepart != "gear_depan" && namasparepart != "conron" && namasparepart != "pullstart" && namasparepart != "ban_luar" && namasparepart != "ban_dalam" && namasparepart != "SELESAI" {
				fmt.Println("  Nama sparepart tidak valid. Masukkan nama sparepart sesuai jenis yang tersedia!")
				fmt.Print("  Masukkan nama sparepart ", "ke-", i+1, " : ")
				fmt.Scan(&namasparepart)
			}
			if namasparepart != "SELESAI" {
				transaction[*iTransaksi].spr.nama[i] = namasparepart
				transaction[*iTransaksi].tarif += hargajenissparepart(namasparepart)
				tambahSparepartGlobal(namasparepart, &*s)
			}
			i++
			if namasparepart != "SELESAI" {
				fmt.Print("  Masukkan nama sparepart ", "ke-", i+1, " : ")
			}

			if namasparepart == "velg" || namasparepart == "kampas_rem" || namasparepart == "karburator" || namasparepart == "piston" || namasparepart == "gear_depan" || namasparepart == "conron" || namasparepart == "pullstart" || namasparepart == "ban_luar" || namasparepart == "ban_dalam" {
				fmt.Scan(&namasparepart)
			}
			transaction[*iTransaksi].spr.banyak = i
		}

		fmt.Println()
		footer_tambah_data()
		*iTransaksi++
	} else if pilih == 2 {
		if *iTransaksi == 0 {
			fmt.Println("")
			fmt.Println("PENAMBAHAN SPAREPART GAGAL KARENA BELUM ADA DATA YANG TERSEDIA   ")
			fmt.Println("        SILAKAN LAKUKAN PENGISIAN DATA TERLEBIH DAHULU   ")
			fmt.Println("\n\n")
		} else {
			fmt.Println("")
			fmt.Println("  MASUKKAN VARIABEL BERIKUT DARI DATA YANG INGIN DITAMBAH")
			fmt.Print("  Nama customer      : ")
			fmt.Scan(&nama)
			fmt.Print("  Tanggal transaksi  : ")
			fmt.Scan(&tgl)
			fmt.Print("  Bulan transaksi    : ")
			fmt.Scan(&bln)
			fmt.Print("  Tahun transaksi    : ")
			fmt.Scan(&thn)
			j := 0
			found := -1
			for j < *iTransaksi && found == -1 {
				if nama == transaction[j].nama && tgl == transaction[j].waktu.tanggal && bln == transaction[j].waktu.bulan && thn == transaction[j].waktu.tahun {
					found = j
				}
				j++
			}
			if found != -1 {
				if transaction[found].spr.banyak == 10 {
					fmt.Println("TIDAK DAPAT MENAMBAHKAN DATA SPAREPART KARENA SPAREPART SUDAH MEMENUHI BATAS MAKSIMAL")
				} else if transaction[found].spr.banyak < 10 {
					fmt.Println("___________________________________________________________")
					fmt.Println("|                                                         |")
					fmt.Println("|    Daftar sparepart yang tersedia beserta hargannya :   |")
					fmt.Println("|             1. velg            Rp 200000                |")
					fmt.Println("|             2. kampas_rem      Rp 50000                 |")
					fmt.Println("|             3. karburator      Rp 350000                |")
					fmt.Println("|             4. piston          Rp 150000                |")
					fmt.Println("|             5. gear_depan      Rp 80000                 |")
					fmt.Println("|             6. conron          Rp 150000                |")
					fmt.Println("|             7. pullstart       Rp 230000                |")
					fmt.Println("|             8. ban_luar        Rp 170000                |")
					fmt.Println("|             9. ban_dalam       Rp 75000                 |")
					fmt.Println("|_________________________________________________________|")
					fmt.Println("  MASUKKAN NAMA-NAMA SPAREPART YANG DIBELI BERDASARKAN SPAREPART YANG TERSEDIA!")
					fmt.Println("  Masukkan 'SELESAI' tanpa tanda petik apabila semua sparepart sudah dimasukkan")
					fmt.Print("  Masukkan nama sparepart ke-", transaction[found].spr.banyak+1, " : ")
					fmt.Scan(&namasparepart)
					z := transaction[found].spr.banyak
					for z < NMAXsparepart-1 && namasparepart != "SELESAI" {
						for namasparepart != "velg" && namasparepart != "kampas_rem" && namasparepart != "karburator" && namasparepart != "piston" && namasparepart != "gear_depan" && namasparepart != "conron" && namasparepart != "pullstart" && namasparepart != "ban_luar" && namasparepart != "ban_dalam" && namasparepart != "SELESAI" {
							fmt.Println("  Nama sparepart tidak valid. Masukkan nama sparepart sesuai jenis yang tersedia!")
							fmt.Print("  Masukkan nama sparepart ", "ke-", z+1, " : ")
							fmt.Scan(&namasparepart)
						}
						if namasparepart != "SELESAI" {
							transaction[found].spr.nama[z] = namasparepart
							transaction[found].tarif += hargajenissparepart(namasparepart)
							tambahSparepartGlobal(namasparepart, &*s)
						}
						z++
						if namasparepart != "SELESAI" {
							fmt.Print("  Masukkan nama sparepart ", "ke-", z+1, " : ")
						}

						if namasparepart == "velg" || namasparepart == "kampas_rem" || namasparepart == "karburator" || namasparepart == "piston" || namasparepart == "gear_depan" || namasparepart == "conron" || namasparepart == "pullstart" || namasparepart == "ban_luar" || namasparepart == "ban_dalam" {
							fmt.Scan(&namasparepart)
						}
						transaction[*iTransaksi].spr.banyak = z

						if z == 9 && namasparepart != "SELESAI" {
							transaction[found].spr.nama[z] = namasparepart
							transaction[found].tarif += hargajenissparepart(namasparepart)
							tambahSparepartGlobal(namasparepart, &*s)
						}
					}

				}
			} else {
				fmt.Println(" ---{ DATA YANG ANDA CARI TIDAK DITEMUKAN }--")
				fmt.Println(" ---{   PENAMBAHAN DATA SPAREPART GAGAL   }--")
			}
		}
	}
}

func footer_tambah_data() {
	fmt.Println("                   PENAMBAHAN DATA BERHASIL                ")
	fmt.Println("===========================================================")
	fmt.Println("\n\n")
}

// Untuk menambah data pada sparepart global //
func tambahSparepartGlobal(nama string, s *tabsparepartglobal) {
	found := -1
	i := 0
	for i < 9 && found == -1 {
		if nama == s[i].nama {
			found = i
		}
		i++
	}
	s[found].n = s[found].n + 1
}

// Untuk menghapus data pada sparepart global //
func kurangSparepartGlobal(nama string, s *tabsparepartglobal) {
	found := -1
	i := 0
	for i < 9 && found == -1 {
		if nama == s[i].nama {
			found = i
		}
		i++
	}
	s[found].n--
}

func hargajenissparepart(jenis string) int {
	if jenis == "velg" {
		return 200000
	} else if jenis == "kampas_rem" {
		return 50000
	} else if jenis == "karburator" {
		return 350000
	} else if jenis == "piston" {
		return 150000
	} else if jenis == "gear_depan" {
		return 80000
	} else if jenis == "conron" {
		return 150000
	} else if jenis == "pullstart" {
		return 230000
	} else if jenis == "ban_luar" {
		return 170000
	} else if jenis == "ban_dalam" {
		return 75000
	}
	return 0
}

func ubah_data(transaction *tabtransaksi, iTransaksi int, s *tabsparepartglobal, iSparepart *int) {
	var tgl, bln, thn, banyaksparepartawal int
	var nama string
	var namasparepart string

	fmt.Println("===========================================================")
	fmt.Println("=                  APLIKASI SERVICE MOTOR                 =")
	fmt.Println("=                   Page Pengubahan Data                  =")
	fmt.Println("===========================================================")
	fmt.Print("  Masukkan nama customer yang datanya ingin diubah   : ")
	fmt.Scan(&nama)
	fmt.Print("  Masukkan tanggal dari data yang ingin diubah       : ")
	fmt.Scan(&tgl)
	fmt.Print("  Masukkan bulan dari data yang ingin diubah         : ")
	fmt.Scan(&bln)
	fmt.Print("  Masukkan tahun dari data yang ingin diubah         : ")
	fmt.Scan(&thn)

	//  Mencari index dari data yang ingin dicari berdasarkan waktu transaksi dan nama pelanggan //
	found := -1
	index := 0
	for index < iTransaksi && found == -1 {
		if nama == transaction[index].nama && tgl == transaction[index].waktu.tanggal && bln == transaction[index].waktu.bulan && thn == transaction[index].waktu.tahun {
			found = index
		}
		index++
	}

	// Diminta menginputkan data kembali apabila data yang ingin diubah tidak ditemukan (terus memasukkan data hingga data ditemukan)//
	for found == -1 {
		fmt.Println("")
		fmt.Println("  Data yang Anda cari tidak ditemukan. Silakan input kembali. ")
		fmt.Println("")
		fmt.Print("  Masukkan nama customer yang datanya ingin diubah   : ")
		fmt.Scan(&nama)
		fmt.Print("  Masukkan tanggal dari data yang ingin diubah       : ")
		fmt.Scan(&tgl)
		fmt.Print("  Masukkan bulan dari data yang ingin diubah         : ")
		fmt.Scan(&bln)
		fmt.Print("  Masukkan tahun dari data yang ingin diubah         : ")
		fmt.Scan(&thn)
		index = 0
		for index < iTransaksi && found == -1 {
			if tgl == transaction[index].waktu.tanggal && bln == transaction[index].waktu.bulan && thn == transaction[index].waktu.tahun {
				found = index
			}
			index++
		}
	}

	banyaksparepartawal = transaction[found].spr.banyak
	for i := 0; i < banyaksparepartawal; i++ {
		kurangSparepartGlobal(transaction[found].spr.nama[i], &*s)
	}

	transaction[found] = transaction[iTransaksi+1]
	transaction[found].tarif = 0

	fmt.Println("===========================================================")
	fmt.Println("                    INPUT DATA TERBARU")
	fmt.Print("  Masukkan nama customer terbaru          : ")
	fmt.Scan(&transaction[found].nama)
	// Satu customer diasumsikan maksimaml hanya melakukan sekali transaksi dalam sehari berdaasarkan estimasi waktu antrian dan lamanya service //
	fmt.Print("  Masukkan tanggal transaksi terbaru      : ")
	fmt.Scan(&transaction[found].waktu.tanggal)
	fmt.Print("  Masukkan bulan transaksi terbaru        : ")
	fmt.Scan(&transaction[found].waktu.bulan)
	fmt.Print("  Masukkan tahun transaksi terbaru        : ")
	fmt.Scan(&transaction[found].waktu.tahun)

	fmt.Println("___________________________________________________________")
	fmt.Println("|                                                         |")
	fmt.Println("|    Daftar sparepart yang tersedia beserta hargannya :   |")
	fmt.Println("|             1. velg            Rp 200000                |")
	fmt.Println("|             2. kampas_rem      Rp 50000                 |")
	fmt.Println("|             3. karburator      Rp 350000                |")
	fmt.Println("|             4. piston          Rp 150000                |")
	fmt.Println("|             5. gear_depan      Rp 80000                 |")
	fmt.Println("|             6. conron          Rp 150000                |")
	fmt.Println("|             7. pullstart       Rp 230000                |")
	fmt.Println("|             8. ban_luar        Rp 170000                |")
	fmt.Println("|             9. ban_dalam       Rp 75000                 |")
	fmt.Println("|_________________________________________________________|")
	fmt.Println("  MASUKKAN NAMA-NAMA SPAREPART YANG DIBELI BERDASARKAN SPAREPART YANG TERSEDIA!")
	fmt.Println("  Masukkan 'SELESAI' tanpa tanda petik apabila semua sparepart sudah dimasukkan")
	fmt.Print("  Masukkan nama sparepart ke-1 : ")
	fmt.Scan(&namasparepart)

	i := 0
	for i < NMAXsparepart-1 && namasparepart != "SELESAI" {
		for namasparepart != "velg" && namasparepart != "kampas_rem" && namasparepart != "karburator" && namasparepart != "piston" && namasparepart != "gear_depan" && namasparepart != "conron" && namasparepart != "pullstart" && namasparepart != "ban_luar" && namasparepart != "ban_dalam" && namasparepart != "SELESAI" {
			fmt.Println("  Nama sparepart tidak valid. Masukkan nama sparepart sesuai jenis yang tersedia!")
			fmt.Print("  Masukkan nama sparepart ", "ke-", i+1, " : ")
			fmt.Scan(&namasparepart)
		}
		if namasparepart != "SELESAI" {
			transaction[found].spr.nama[i] = namasparepart
			transaction[found].tarif += hargajenissparepart(namasparepart)
			tambahSparepartGlobal(namasparepart, &*s)
		}
		i++
		if namasparepart != "SELESAI" {
			fmt.Print("  Masukkan nama sparepart ", "ke-", i+1, " : ")
		}

		if namasparepart == "velg" || namasparepart == "kampas_rem" || namasparepart == "karburator" || namasparepart == "piston" || namasparepart == "gear_depan" || namasparepart == "conron" || namasparepart == "pullstart" || namasparepart == "ban_luar" || namasparepart == "ban_dalam" {
			fmt.Scan(&namasparepart)
		}
		transaction[found].spr.banyak = i
	}

	fmt.Println()

	fmt.Println("                   PENGUBAHAN DATA BERHASIL                ")
	fmt.Println("===========================================================")
	fmt.Println("\n\n")
}

func hapus_data(transaction *tabtransaksi, iTransaksi *int, s *tabsparepartglobal) {
	var tgl, bln, thn, banyaksparepartawal int
	var nama string
	fmt.Println("===========================================================")
	fmt.Println("=                  APLIKASI SERVICE MOTOR                 =")
	fmt.Println("=                   Page Penghapusan Data                 =")
	fmt.Println("===========================================================")
	fmt.Print("  Masukkan nama customer yang datanya ingin dihapus  : ")
	fmt.Scan(&nama)
	fmt.Print("  Masukkan tanggal dari data yang ingin dihapus      : ")
	fmt.Scan(&tgl)
	fmt.Print("  Masukkan bulan dari data yang ingin dihapus        : ")
	fmt.Scan(&bln)
	fmt.Print("  Masukkan tahun dari data yang ingin dihapus        : ")
	fmt.Scan(&thn)

	index := 0
	found := -1
	for index < *iTransaksi && found == -1 {
		if nama == transaction[index].nama && tgl == transaction[index].waktu.tanggal && bln == transaction[index].waktu.bulan && thn == transaction[index].waktu.tahun {
			found = index
		}
		index++
	}

	for found == -1 {
		fmt.Println("  Data yang diinputkan tidak valid. Masukkan data yang sesuai!")
		fmt.Print("  Masukkan nama customer yang datanya ingin dihapus: ")
		fmt.Scan(&nama)
		fmt.Print("  Masukkan tanggal dari data yang ingin dihapus: ")
		fmt.Scan(&tgl)
		fmt.Print("  Masukkan bulan dari data yang ingin dihapus: ")
		fmt.Scan(&bln)
		fmt.Print("  Masukkan tahun dari data yang ingin dihapus: ")
		fmt.Scan(&thn)
		index = 0
		for index < *iTransaksi && found == -1 {
			if tgl == transaction[index].waktu.tanggal && bln == transaction[index].waktu.bulan && thn == transaction[index].waktu.tahun {
				found = index
			}
			index++
		}
	}

	for index < *iTransaksi-1 {
		transaction[index] = transaction[index+1]
		index++
	}
	*iTransaksi--

	banyaksparepartawal = transaction[found].spr.banyak
	for i := 0; i < banyaksparepartawal; i++ {
		kurangSparepartGlobal(transaction[found].spr.nama[i], &*s)
	}

	transaction[found] = transaction[*iTransaksi+1]

	fmt.Println("                  PENGHAPUSAN DATA BERHASIL                ")
	fmt.Println("===========================================================")
	fmt.Println("\n\n")
}

func cari_data(transaction tabtransaksi, iTransaksi, iSparepart int, s tabsparepartglobal, n *pembelisparepart) {
	var pilihan, tgl, bln, thn int
	var spr string
	fmt.Println("===========================================================")
	fmt.Println("=                  APLIKASI SERVICE MOTOR                 =")
	fmt.Println("=                   Page Pencarian Data                   =")
	fmt.Println("===========================================================")
	fmt.Println("=             Pilihan Pencarian yang Tersedia:            =")
	fmt.Println("=  1. Daftar pelanggan berdasarkan waktu transaksi        =")
	fmt.Println("=  2. Daftar pelanggan berdasarkan sparepart yang dibeli  =")
	fmt.Println("-----------------------------------------------------------")
	fmt.Print("  Masukkan pilihan Anda : ")
	fmt.Scan(&pilihan)
	for pilihan != 1 && pilihan != 2 {
		fmt.Println("  Pilihan Anda tidak valid. Silakan pilih opsi yang tersedia!")
		fmt.Print("  Masukkan pilihan Anda : ")
		fmt.Scan(&pilihan)
	}

	if pilihan == 1 {
		fmt.Print("  Masukkan tanggal dari data yang ingin dicari   : ")
		fmt.Scan(&tgl)
		fmt.Print("  Masukkan bulan dari data yang ingin dicari     : ")
		fmt.Scan(&bln)
		fmt.Print("  Masukkan tahun dari data yang ingin dicari     : ")
		fmt.Scan(&thn)

		i := 0
		found := -1
		for i < iTransaksi && found == -1 {
			if transaction[i].waktu.tanggal == tgl && transaction[i].waktu.bulan == bln && transaction[i].waktu.tahun == thn {
				found = i
			}
			i++
		}

		no := 0
		if found != -1 {
			fmt.Println("")
			fmt.Println("===========================================================")
			fmt.Println("  Pelanggan yang melakukan transaksi pada ", tgl, "-", bln, "-", thn, " yaitu:")
			for i := 0; i < iTransaksi; i++ {
				if transaction[i].waktu.tanggal == tgl && transaction[i].waktu.bulan == bln && transaction[i].waktu.tahun == thn {
					fmt.Println("  ", no+1, ".", transaction[i].nama)
					no++
				}
			}
			fmt.Println("===========================================================")
			fmt.Println("\n\n")
		} else {
			fmt.Println("  Tidak ada pelanggan yang melakukan transaksi pada tanggal tersebut")
		}
	} else {
		fmt.Print("  Masukkan nama sparepart:")
		fmt.Scan(&spr)

		// Apabila input tidak valid maka akan dimintai input nama sparepart kembali//
		for spr != "velg" && spr != "kampas_rem" && spr != "karburator" && spr != "piston" && spr != "gear_depan" && spr != "conron" && spr != "pullstart" && spr != "ban_luar" && spr != "ban_dalam" {
			fmt.Println("  Nama sparepart tidak valid. Masukkan nama sparepart sesuai jenis yang tersedia!")
			fmt.Print("  Masukkan nama sparepart:")
			fmt.Scan(&spr)
		}

		// Menacari nama sparepart ada di indeks berapa di data sparepartglobal//
		i := 0
		found := -1
		for i < 9 && found == -1 {
			if s[i].nama == spr {
				found = i
			}
			i++
		}

		no := 1
		i = 0
		k := 0
		if s[found].n != 0 {
			fmt.Println("")
			fmt.Println("===========================================================")
			for i < iTransaksi {
				found = -1
				j := 0
				for j < transaction[i].spr.banyak && found == -1 {
					if transaction[i].spr.nama[j] == spr {
						if k == 0 {
							fmt.Println("  Daftar pelanggan yang membeli ", spr, " yaitu:")
							fmt.Println("  ", no, ".", transaction[i].nama)
							n[0] = transaction[i].nama
							k++
						} else {
							z := 0
							cek := false
							for z < k && !cek {
								cek = n[z] == transaction[i].nama
								z++
							}
							if !cek {
								fmt.Println("  Daftar pelanggan yang membeli ", spr, " yaitu:")
								fmt.Println("  ", no, ".", transaction[i].nama)
								n[k] = transaction[i].nama
								k++
							}

						}
						found = i
						no++
					}
					j++
				}
				i++
			}
			fmt.Println("")
			fmt.Println("===========================================================")
		} else {
			fmt.Println("  Tidak ada pelanggan yang membeli ", spr)

		}
	}
}

func tampil_data(sparepart tabsparepartglobal, transaction *tabtransaksi, iTransaksi *int) {
	var pass, i, idx int
	var s sparepartglobal
	//var temp transaksi
	var pilihan, pilihan2 int

	fmt.Println("===========================================================")
	fmt.Println("=                  APLIKASI SERVICE MOTOR                 =")
	fmt.Println("=                   Page Menampilkan Data                 =")
	fmt.Println("===========================================================")
	fmt.Println("=                  Pilihan yang tersedia :                =")
	fmt.Println("=                     1. Sparepart                        =")
	fmt.Println("=                     2. Invoice                          =")
	fmt.Println("=                     3. Kembali                          =")
	fmt.Println("===========================================================")
	fmt.Print("             Masukkan pilihan Anda : ")
	fmt.Scan(&pilihan)
	fmt.Println("\n\n")
	fmt.Println("===========================================================")
	for pilihan != 1 && pilihan != 2 && pilihan != 3 {
		fmt.Println("  Pilihan Anda tidak valid. Silakan pilih opsi yang tersedia!")
		fmt.Print("  Masukkan pilihan Anda : ")
		fmt.Scan(&pilihan)
	}
	if pilihan == 1 {
		fmt.Println("=                  Pilihan yang tersedia :                =")
		fmt.Println("=   1. Tampilkan sparepart dari yang sering diganti       =")
		fmt.Println("=   2. Tampilkan sparepart dari yang jarang diganti       =")
		fmt.Println("=   3. Kembali                                            =")
		fmt.Println("===========================================================")
		fmt.Print("  Masukkan pilihan Anda : ")
		fmt.Scan(&pilihan2)
		fmt.Println("===========================================================")
		fmt.Println("")
		for pilihan2 != 1 && pilihan2 != 2 && pilihan2 != 3 {
			fmt.Println("  Pilihan Anda tidak valid. Silakan pilih opsi yang tersedia!")
			fmt.Print("  Masukkan pilihan Anda : ")
			fmt.Scan(&pilihan2)
		}
		if pilihan2 == 1 {
			// INSERTION SORT  untuk menampilkan data sparepart secara terurut dari yang paling sering diganti (descending) //
			pass = 1
			for pass < 9 {
				i = pass
				s = sparepart[pass]
				for i > 0 && s.n > sparepart[i-1].n {
					sparepart[i] = sparepart[i-1]
					i--
				}
				sparepart[i] = s
				pass++
			}
			fmt.Println("")
			fmt.Println("  Berikut data terurut berdasarkan sparepart yang paling sering diganti : ")
			for i = 0; i < 9; i++ {
				fmt.Println("  ", i+1, ". ", sparepart[i].nama, " sebanyak ", sparepart[i].n)
			}
		} else if pilihan == 2 {
			// SELECTION SORT untuk menampilkan data sparepart secara terurut dari yang paling jarang diganti (ascending) //
			pass = 1
			for pass < 9 {
				i = pass
				idx = pass - 1
				for i < 9 {
					if sparepart[idx].n > sparepart[i].n {
						idx = i
					}
					i++
				}
				s = sparepart[pass-1]
				sparepart[pass-1] = sparepart[idx]
				sparepart[idx] = s
				pass++
			}
			fmt.Println("")
			fmt.Println("  Berikut data terurut berdasarkan sparepart yang paling jarang diganti : ")
			for i = 0; i < 9; i++ {
				fmt.Println("  ", i+1, ". ", sparepart[i].nama, " sebanyak ", sparepart[i].n)
			}
		}
		/*} else if pilihan == 2 {
			fmt.Println("= 	                Pilihan yang tersedia :                =")
			fmt.Println("=   1. Tampilkan Transaksi dari tarif terendah            =")
			fmt.Println("=   2. Tampilkan Transaksi dari tarif tertingi            =")
			fmt.Println("=   3. Kembali                                            =")
			fmt.Println("===========================================================")
			fmt.Print("  Masukkan pilihan Anda : ")
			fmt.Scan(&pilihan2)
			fmt.Println("===========================================================")
			for pilihan2 != 1 && pilihan2 != 2 && pilihan2 != 3 {
				fmt.Println("  Pilihan Anda tidak valid. Silakan pilih opsi yang tersedia!")
				fmt.Print("  Masukkan pilihan Anda : ")
				fmt.Scan(&pilihan2)
			}
			if pilihan2 == 1 {
				pass = 1
				for pass < *iTransaksi {
					i = pass
					temp = transaction[pass]
					for i > 0 && temp.tarif < transaction[i-1].tarif {
						transaction[i] = transaction[i-1]
						i--
					}
					transaction[i] = temp
					pass++
				}
				fmt.Println("  Berikut data terurut berdasarkan tarif terendah : ")
				for i = 0; i < *iTransaksi; i++ {
					fmt.Println("  ", i+1, ". ", transaction[i].nama, transaction[i].tarif, transaction[i].waktu.tanggal, transaction[i].waktu.bulan, transaction[i].waktu.tahun)
				}
			} else if pilihan2 == 2 {
				pass = 1
				for pass < *iTransaksi {
					i = pass
					temp = transaction[pass]
					for i > 0 && temp.tarif > transaction[i-1].tarif {
						transaction[i] = transaction[i-1]
						i--
					}
					transaction[i] = temp
					pass++
				}
				fmt.Println("  Berikut data terurut berdasarkan tarif tertinggi : ")
				for i = 0; i < *iTransaksi; i++ {
					fmt.Println("  ", i+1, ". ", transaction[i].nama, transaction[i].tarif, transaction[i].waktu.tanggal, transaction[i].waktu.bulan, transaction[i].waktu.tahun)
				}
			} else if pilihan2 == 3 {

			}
		} else if pilihan == 3 {
			fmt.Println("")
			fmt.Println("Data Sparepart Sebagai berikut : ")
			for i = 0; i < 9; i++ {
				fmt.Println("  ", i+1, ". ", sparepart[i].nama, " sebanyak ", sparepart[i].n)
			}
			fmt.Println("")
			fmt.Println("Data Transaksi Sebagai berikut : ")
			for i = 0; i < *iTransaksi; i++ {
				fmt.Println("  ", i+1, ". ", transaction[i].nama, " | Rp.", transaction[i].tarif, " | ", transaction[i].waktu.tanggal, "-", transaction[i].waktu.bulan, "-", transaction[i].waktu.tahun)
			}*/
	} else if pilihan == 2 {
		invoice(*transaction, *iTransaksi)
	}
	fmt.Println("")
}

func invoice(transaction tabtransaksi, iTransaksi int) {
	var nama string
	var tgl, bln, thn int
	fmt.Println("===========================================================")
	fmt.Println("=                  APLIKASI SERVICE MOTOR                 =")
	fmt.Println("=                 Page Menampilkan Invoice                =")
	fmt.Println("===========================================================")
	fmt.Println("   Masukkan data yang ingin ditampilkan invoice-nya")
	fmt.Print("   Nama pelanggan       : ")
	fmt.Scan(&nama)
	fmt.Print("   Tanggal transaksi    : ")
	fmt.Scan(&tgl)
	fmt.Print("   Bulan transaksi      : ")
	fmt.Scan(&bln)
	fmt.Print("   Tahun transaksi      : ")
	fmt.Scan(&thn)

	found := -1
	index := 0
	for index < iTransaksi && found == -1 {
		if nama == transaction[index].nama && tgl == transaction[index].waktu.tanggal && bln == transaction[index].waktu.bulan && thn == transaction[index].waktu.tahun {
			found = index
		}
		index++
	}

	for found == -1 {
		fmt.Println("  Data yang Anda cari tidak ditemukan. Silakan masukkan data kembali")
		fmt.Print("   Nama pelanggan       : ")
		fmt.Scan(&nama)
		fmt.Print("   Tanggal transaksi    : ")
		fmt.Scan(&tgl)
		fmt.Print("   Bulan transaksi      : ")
		fmt.Scan(&bln)
		fmt.Print("   Tahun transaksi      : ")
		fmt.Scan(&thn)
		index = 0
		for index < iTransaksi && found == -1 {
			if tgl == transaction[index].waktu.tanggal && bln == transaction[index].waktu.bulan && thn == transaction[index].waktu.tahun {
				found = index
			}
			index++
		}
	}

	fmt.Println("         ____________________________________________")
	fmt.Println("        |                                            |")
	fmt.Println("        |         TAGIHAN BIAYA SERVICE MOTOR        |")
	fmt.Printf("        |                %-1d / %-1d / %-18d | \n", transaction[found].waktu.tanggal, transaction[found].waktu.bulan, transaction[found].waktu.tahun)
	fmt.Printf("        |            Customer : %-20s | \n", transaction[found].nama)
	fmt.Println("        |                                            |")

	for i := 0; i < transaction[found].spr.banyak; i++ {
		nama = transaction[found].spr.nama[i]
		harga := hargajenissparepart(transaction[found].spr.nama[i])
		fmt.Printf("        |        %-16s Rp %-15d | \n", nama, harga)
	}

	fmt.Printf("        |             TOTAL :     Rp %-15d | \n", transaction[found].tarif)
	fmt.Println("        |____________________________________________|")
}
