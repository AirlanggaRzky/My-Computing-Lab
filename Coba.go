package main

import "fmt"

type Tanggal struct {
	DD int
	MM int
	YY int
}

type Kendaraan struct {
	PlatNo        string
	NamaPemilik   string
	TahunProduksi int
	LastServis    Tanggal
}

type RiwayatServis struct {
	PlatNo           string
	WaktuServis      Tanggal
	JenisKerusakanID int
}

type JenisKerusakan struct {
	JenisKerusakanID int
	NamaKerusakan    string
}

var DataKendaraan [100]Kendaraan
var DataRiwayatServis [100]RiwayatServis
var DataJenisKerusakan [8]JenisKerusakan
var lastIndexKendaraan int = 0
var lastIndexRiwayat int = 0

// ===== HELPER =====
func TanggalKeInt(t Tanggal) int {
	return t.YY*10000 + t.MM*100 + t.DD
}

func CetakTanggal(t Tanggal) {
	fmt.Printf("%02d/%02d/%04d", t.DD, t.MM, t.YY)
}

func InputTanggal() Tanggal {
	var t Tanggal
	fmt.Print("Hari: ")
	fmt.Scan(&t.DD)
	fmt.Print("Bulan: ")
	fmt.Scan(&t.MM)
	fmt.Print("Tahun: ")
	fmt.Scan(&t.YY)
	return t
}

func CariIndeksKendaraan(plat string) int {
	var i int = 0
	for i < lastIndexKendaraan {
		if DataKendaraan[i].PlatNo == plat {
			return i
		}
		i = i + 1
	}
	return -1
}

func CariNamaKerusakan(id int) string {
	var i int = 0
	for i < len(DataJenisKerusakan) {
		if DataJenisKerusakan[i].JenisKerusakanID == id {
			return DataJenisKerusakan[i].NamaKerusakan
		}
		i = i + 1
	}
	return "Tidak diketahui"
}

func TampilDaftarJenisKerusakan() {
	fmt.Println("  Daftar Jenis Kerusakan:")
	var i int = 0
	for i < len(DataJenisKerusakan) {
		fmt.Printf("  [%d] %s\n", DataJenisKerusakan[i].JenisKerusakanID, DataJenisKerusakan[i].NamaKerusakan)
		i = i + 1
	}
}

func InisialisasiJenisKerusakan() {
	DataJenisKerusakan[0] = JenisKerusakan{1, "Ganti Oli"}
	DataJenisKerusakan[1] = JenisKerusakan{2, "Rem Blong"}
	DataJenisKerusakan[2] = JenisKerusakan{3, "Mesin Overheat"}
	DataJenisKerusakan[3] = JenisKerusakan{4, "Aki Lemah"}
	DataJenisKerusakan[4] = JenisKerusakan{5, "Ban Bocor"}
	DataJenisKerusakan[5] = JenisKerusakan{6, "Servis Rutin"}
	DataJenisKerusakan[6] = JenisKerusakan{7, "Kerusakan Transmisi"}
	DataJenisKerusakan[7] = JenisKerusakan{8, "Lainnya"}
}

// A. MANAJEMEN KENDARAAN

func TambahKendaraan() {
	var k Kendaraan
	fmt.Print("Plat Nomor: ")
	fmt.Scan(&k.PlatNo)

	if CariIndeksKendaraan(k.PlatNo) != -1 {
		fmt.Println("Plat nomor sudah terdaftar.")
		return
	}

	if lastIndexKendaraan >= 100 {
		fmt.Println("Data kendaraan penuh.")
		return
	}

	fmt.Print("Nama Pemilik: ")
	fmt.Scan(&k.NamaPemilik)
	fmt.Print("Tahun Produksi: ")
	fmt.Scan(&k.TahunProduksi)

	k.LastServis = Tanggal{0, 0, 0}

	DataKendaraan[lastIndexKendaraan] = k
	lastIndexKendaraan = lastIndexKendaraan + 1
	fmt.Println("Kendaraan berhasil ditambahkan.")
}

func HapusKendaraan() {
	var plat string
	fmt.Print("Plat Nomor yang akan dihapus: ")
	fmt.Scan(&plat)

	var idx int = CariIndeksKendaraan(plat)
	if idx == -1 {
		fmt.Println("Kendaraan tidak ditemukan.")
		return
	}

	// Geser elemen DataKendaraan ke kiri mulai dari idx
	var i int = idx
	for i < lastIndexKendaraan-1 {
		DataKendaraan[i] = DataKendaraan[i+1]
		i = i + 1
	}
	DataKendaraan[lastIndexKendaraan-1] = Kendaraan{}
	lastIndexKendaraan = lastIndexKendaraan - 1

	// Hapus semua riwayat servis terkait
	i = 0
	for i < lastIndexRiwayat {
		if DataRiwayatServis[i].PlatNo == plat {
			// Geser elemen DataRiwayatServis ke kiri mulai dari i
			var j int = i
			for j < lastIndexRiwayat-1 {
				DataRiwayatServis[j] = DataRiwayatServis[j+1]
				j = j + 1
			}
			DataRiwayatServis[lastIndexRiwayat-1] = RiwayatServis{}
			lastIndexRiwayat = lastIndexRiwayat - 1
		} else {
			i = i + 1
		}
	}

	fmt.Println("Kendaraan dan riwayat servisnya berhasil dihapus.")
}

func UbahKepemilikan() {
	var plat string
	var namaBaru string
	fmt.Print("Plat Nomor: ")
	fmt.Scan(&plat)

	var idx int = CariIndeksKendaraan(plat)
	if idx == -1 {
		fmt.Println("Kendaraan tidak ditemukan.")
		return
	}

	fmt.Printf("Pemilik lama: %s\n", DataKendaraan[idx].NamaPemilik)
	fmt.Print("Pemilik baru: ")
	fmt.Scan(&namaBaru)

	DataKendaraan[idx].NamaPemilik = namaBaru
	fmt.Println("Kepemilikan berhasil diubah.")
}

func MenuManajemenKendaraan() {
	var pilihan int
	var lanjut bool = true
	for lanjut {
		fmt.Println("\n====== Manajemen Kendaraan ======")
		fmt.Println("[1] Tambah Kendaraan")
		fmt.Println("[2] Hapus Kendaraan")
		fmt.Println("[3] Ubah Kepemilikan")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahKendaraan()
		case 2:
			HapusKendaraan()
		case 3:
			UbahKepemilikan()
		case 0:
			lanjut = false
		default:
			fmt.Println("Input tidak valid.")
		}
	}
}

// B. PENCATATAN RIWAYAT SERVIS

func CatatRiwayatServis() {
	var plat string
	fmt.Print("Plat Nomor: ")
	fmt.Scan(&plat)

	var idx int = CariIndeksKendaraan(plat)
	if idx == -1 {
		fmt.Println("Kendaraan tidak ditemukan.")
		return
	}

	TampilDaftarJenisKerusakan()
	var idKerusakan int
	fmt.Print("Pilih ID Kerusakan: ")
	fmt.Scan(&idKerusakan)

	var namaKerusakan string = CariNamaKerusakan(idKerusakan)
	if namaKerusakan == "Tidak diketahui" {
		fmt.Println("ID kerusakan tidak valid.")
		return
	}

	if lastIndexRiwayat >= 100 {
		fmt.Println("Data riwayat servis penuh.")
		return
	}

	fmt.Println("Tanggal Servis: ")
	var tgl Tanggal = InputTanggal()

	DataRiwayatServis[lastIndexRiwayat] = RiwayatServis{PlatNo: plat, WaktuServis: tgl, JenisKerusakanID: idKerusakan}
	lastIndexRiwayat = lastIndexRiwayat + 1

	// Update lastServis jika tanggal baru lebih akhir
	if TanggalKeInt(tgl) > TanggalKeInt(DataKendaraan[idx].LastServis) {
		DataKendaraan[idx].LastServis = tgl
	}

	fmt.Printf("Servis '%s' untuk plat %s berhasil dicatat.\n", namaKerusakan, plat)
}

// C. PENCARIAN DATA KENDARAAN

func SequentialSearchKendaraan(plat string) int {
	var i int = 0
	for i < lastIndexKendaraan {
		if DataKendaraan[i].PlatNo == plat {
			return i
		}
		i = i + 1
	}
	return -1
}

func BinarySearchKendaraan(platSorted [100]Kendaraan, plat string) int {
	var kiri int = 0
	var kanan int = lastIndexKendaraan - 1
	for kiri <= kanan {
		var tengah int = (kiri + kanan) / 2
		if platSorted[tengah].PlatNo == plat {
			return tengah
		} else if platSorted[tengah].PlatNo < plat {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func UrutkanPlatAlfabetis() [100]Kendaraan {
	var sorted [100]Kendaraan
	// Salin data aktif ke sorted
	var i int = 0
	for i < lastIndexKendaraan {
		sorted[i] = DataKendaraan[i]
		i = i + 1
	}
	// Insertion sort berdasarkan PlatNo
	i = 1
	for i < lastIndexKendaraan {
		var kunci Kendaraan = sorted[i]
		var j int = i - 1
		for j >= 0 && sorted[j].PlatNo > kunci.PlatNo {
			sorted[j+1] = sorted[j]
			j = j - 1
		}
		sorted[j+1] = kunci
		i = i + 1
	}
	return sorted
}

func CetakInfoKendaraan(k Kendaraan) {
	fmt.Printf("Plat No: %s\n", k.PlatNo)
	fmt.Printf("Pemilik: %s\n", k.NamaPemilik)
	fmt.Printf("Tahun Produksi: %d\n", k.TahunProduksi)
	fmt.Print("Terakhir Servis: ")
	if TanggalKeInt(k.LastServis) == 0 {
		fmt.Println("Belum pernah servis")
	} else {
		CetakTanggal(k.LastServis)
		fmt.Println()
	}
}

func MenuCariKendaraan() {
	if lastIndexKendaraan == 0 {
		fmt.Println("Belum ada data kendaraan.")
		return
	}

	var plat string
	fmt.Print("Masukkan Plat Nomor yang dicari: ")
	fmt.Scan(&plat)

	// Sequential Search
	fmt.Println("Sequential Search:")
	var idxSeq int = SequentialSearchKendaraan(plat)
	if idxSeq != -1 {
		fmt.Println("Kendaraan ditemukan:")
		CetakInfoKendaraan(DataKendaraan[idxSeq])
	} else {
		fmt.Println("Kendaraan tidak ditemukan.")
	}

	// Binary Search (pada salinan terurut)
	fmt.Println("\nBinary Search (pada data terurut berdasarkan plat)")
	var sorted [100]Kendaraan = UrutkanPlatAlfabetis()
	var idxBin int = BinarySearchKendaraan(sorted, plat)
	if idxBin != -1 {
		fmt.Println("Kendaraan ditemukan:")
		CetakInfoKendaraan(sorted[idxBin])
	} else {
		fmt.Println("Kendaraan tidak ditemukan.")
	}
}

// D. PENGURUTAN DATA

func SelectionSortTahun() {
	var n int = lastIndexKendaraan
	var i int = 0
	for i < n-1 {
		var idxMin int = i
		var j int = i + 1
		for j < n {
			if DataKendaraan[j].TahunProduksi < DataKendaraan[idxMin].TahunProduksi {
				idxMin = j
			}
			j = j + 1
		}
		DataKendaraan[i], DataKendaraan[idxMin] = DataKendaraan[idxMin], DataKendaraan[i]
		i = i + 1
	}
	fmt.Println("Data kendaraan diurutkan berdasarkan Tahun Produksi (Selection Sort).")
}

func InsertionSortLastServis() {
	var n int = lastIndexKendaraan
	var i int = 1
	for i < n {
		var kunci Kendaraan = DataKendaraan[i]
		var j int = i - 1
		for j >= 0 && TanggalKeInt(DataKendaraan[j].LastServis) > TanggalKeInt(kunci.LastServis) {
			DataKendaraan[j+1] = DataKendaraan[j]
			j = j - 1
		}
		DataKendaraan[j+1] = kunci
		i = i + 1
	}
	fmt.Println("Data kendaraan diurutkan berdasarkan Tanggal Servis Terakhir (Insertion Sort).")
}

func MenuPengurutan() {
	if lastIndexKendaraan == 0 {
		fmt.Println("Belum ada data kendaraan.")
		return
	}
	var pilihan int
	fmt.Println("\n====== Pengurutan Data ======")
	fmt.Println("[1] Urutkan berdasarkan Tahun Produksi (Selection Sort)")
	fmt.Println("[2] Urutkan berdasarkan Tanggal Servis Terakhir (Insertion Sort)")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		SelectionSortTahun()
		TampilDaftarKendaraan()
	case 2:
		InsertionSortLastServis()
		TampilDaftarKendaraan()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// E. STATISTIK

func StatistikPerBulan() {
	var bulan int
	var tahun int
	fmt.Print("Masukkan Bulan (1-12): ")
	fmt.Scan(&bulan)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&tahun)

	var jumlah int = 0
	var i int = 0
	for i < lastIndexRiwayat {
		if DataRiwayatServis[i].WaktuServis.MM == bulan &&
			DataRiwayatServis[i].WaktuServis.YY == tahun {
			jumlah = jumlah + 1
		}
		i = i + 1
	}
	fmt.Printf("Jumlah servis pada %02d/%d: %d kendaraan\n", bulan, tahun, jumlah)
}

func StatistikKerusakanTerbanyak() {
	if lastIndexRiwayat == 0 {
		fmt.Println("Belum ada riwayat servis.")
		return
	}

	var ids [50]int
	var frekuensi [50]int
	var jumlahUnik int = 0

	var i int = 0
	for i < lastIndexRiwayat {
		var id int = DataRiwayatServis[i].JenisKerusakanID
		var ditemukan bool = false
		var j int = 0
		for j < jumlahUnik && !ditemukan {
			if ids[j] == id {
				frekuensi[j] = frekuensi[j] + 1
				ditemukan = true
			}
			j = j + 1
		}
		if !ditemukan {
			ids[jumlahUnik] = id
			frekuensi[jumlahUnik] = 1
			jumlahUnik = jumlahUnik + 1
		}
		i = i + 1
	}

	var maxFreq int = 0
	var maxIdx int = 0
	i = 0
	for i < jumlahUnik {
		if frekuensi[i] > maxFreq {
			maxFreq = frekuensi[i]
			maxIdx = i
		}
		i = i + 1
	}

	fmt.Printf("Kerusakan terbanyak: %s (%d kali)\n",
		CariNamaKerusakan(ids[maxIdx]), maxFreq)

	fmt.Println("\nRekap seluruh jenis kerusakan:")
	i = 0
	for i < jumlahUnik {
		fmt.Printf("    %-25s: %d kali\n", CariNamaKerusakan(ids[i]), frekuensi[i])
		i = i + 1
	}
}

func MenuStatistik() {
	var pilihan int
	fmt.Println("\n====== Statistik ======")
	fmt.Println("[1] Jumlah servis per bulan")
	fmt.Println("[2] Jenis kerusakan paling sering")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		StatistikPerBulan()
	case 2:
		StatistikKerusakanTerbanyak()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// F. TAMPILAN DATA

func TampilDaftarKendaraan() {
	if lastIndexKendaraan == 0 {
		fmt.Println("Belum ada data kendaraan.")
		return
	}
	fmt.Println("\n=========== Daftar Kendaraan ============")
	fmt.Printf("  %-12s %-15s %-6s %s\n", "Plat No", "Pemilik", "Tahun", "Terakhir Servis")
	fmt.Println("  ----------------------------------------")
	var i int = 0
	for i < lastIndexKendaraan {
		var k Kendaraan = DataKendaraan[i]
		fmt.Printf("  %-12s %-15s %-6d ", k.PlatNo, k.NamaPemilik, k.TahunProduksi)
		if TanggalKeInt(k.LastServis) == 0 {
			fmt.Println("-")
		} else {
			CetakTanggal(k.LastServis)
			fmt.Println()
		}
		i = i + 1
	}
}

func TampilRiwayatServis() {
	if lastIndexRiwayat == 0 {
		fmt.Println("Belum ada riwayat servis.")
		return
	}
	fmt.Println("\n====== Riwayat Servis ======")
	fmt.Printf("  %-12s %-12s %s\n", "Plat No", "Tanggal", "Jenis Kerusakan")
	fmt.Println("  ---------------------------------------------")
	var i int = 0
	for i < lastIndexRiwayat {
		var r RiwayatServis = DataRiwayatServis[i]
		fmt.Printf("  %-12s ", r.PlatNo)
		CetakTanggal(r.WaktuServis)
		fmt.Printf("  %s\n", CariNamaKerusakan(r.JenisKerusakanID))
		i = i + 1
	}
}

func MenuTampilData() {
	var pilihan int
	fmt.Println("\n====== Tampilkan Data ======")
	fmt.Println("[1] Daftar Kendaraan")
	fmt.Println("[2] Riwayat Servis")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		TampilDaftarKendaraan()
	case 2:
		TampilRiwayatServis()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func TampilMenuUtama() {
	fmt.Println("\n==================================")
	fmt.Println("   AutoCare - Servis Kendaraan  ")
	fmt.Println("==================================")
	fmt.Println("[1] Manajemen Kendaraan")
	fmt.Println("[2] Catat Riwayat Servis")
	fmt.Println("[3] Cari Kendaraan")
	fmt.Println("[4] Urutkan Data")
	fmt.Println("[5] Statistik")
	fmt.Println("[6] Tampilkan Data")
	fmt.Println("[0] Keluar")
	fmt.Println("==================================")
	fmt.Print("Pilihan: ")
}

func InisialisasiDataDummy() {
	DataKendaraan[0] = Kendaraan{"B1234ABC", "Budi Santoso", 2018, Tanggal{15, 3, 2024}}
	DataKendaraan[1] = Kendaraan{"D5678XYZ", "Siti Rahayu", 2020, Tanggal{22, 7, 2024}}
	DataKendaraan[2] = Kendaraan{"F9012JKL", "Ahmad Fauzi", 2015, Tanggal{5, 11, 2023}}
	DataKendaraan[3] = Kendaraan{"H3456MNO", "Dewi Lestari", 2022, Tanggal{0, 0, 0}}
	DataKendaraan[4] = Kendaraan{"L7890PQR", "Riko Pratama", 2019, Tanggal{30, 1, 2025}}
	lastIndexKendaraan = 5

	DataRiwayatServis[0] = RiwayatServis{"B1234ABC", Tanggal{10, 1, 2024}, 1}
	DataRiwayatServis[1] = RiwayatServis{"B1234ABC", Tanggal{15, 3, 2024}, 6}
	DataRiwayatServis[2] = RiwayatServis{"D5678XYZ", Tanggal{5, 5, 2024}, 2}
	DataRiwayatServis[3] = RiwayatServis{"D5678XYZ", Tanggal{22, 7, 2024}, 1}
	DataRiwayatServis[4] = RiwayatServis{"D5678XYZ", Tanggal{22, 7, 2024}, 4}
	DataRiwayatServis[5] = RiwayatServis{"F9012JKL", Tanggal{5, 11, 2023}, 3}
	DataRiwayatServis[6] = RiwayatServis{"F9012JKL", Tanggal{20, 2, 2024}, 7}
	DataRiwayatServis[7] = RiwayatServis{"F9012JKL", Tanggal{1, 8, 2024}, 1}
	DataRiwayatServis[8] = RiwayatServis{"L7890PQR", Tanggal{10, 9, 2024}, 5}
	DataRiwayatServis[9] = RiwayatServis{"L7890PQR", Tanggal{30, 1, 2025}, 6}
	lastIndexRiwayat = 10

	fmt.Println("Data dummy berhasil dimuat.")
}

func main() {
	InisialisasiJenisKerusakan()
	InisialisasiDataDummy()

	var pilihan int
	var lanjut bool = true

	for lanjut {
		TampilMenuUtama()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			MenuManajemenKendaraan()
		case 2:
			CatatRiwayatServis()
		case 3:
			MenuCariKendaraan()
		case 4:
			MenuPengurutan()
		case 5:
			MenuStatistik()
		case 6:
			MenuTampilData()
		case 0:
			fmt.Println("Terima kasih telah menggunakan AutoCare. Sampai jumpa!")
			lanjut = false
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}
