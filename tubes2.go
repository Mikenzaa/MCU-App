package main

import "fmt"

const NMAX = 10

type MCU struct{
	ID int
	namaPasien string
	tanggal string
	namaPaket string
	harga int
}
type tabMCU[NMAX] MCU
var indexPasien int = 0
var indexPaket int = 0
var nextId = 1

func main(){
	var mcu tabMCU
	var nama, namaEdit string
	var i,id int

for {
	fmt.Println("========= Selamat datang ========")
	fmt.Println("=        Medical Check Up       =")
	fmt.Println("=================================")
	fmt.Println("1. Modifikasi")
	fmt.Println("2. Tampilkan")
	fmt.Println("3. Cari")
	fmt.Println("4. Exit")
	fmt.Println("----------------")
	fmt.Println("Pilihan anda : ")
	fmt.Scan(&i)
	if i == 1 {
		fmt.Println("=================================")
		fmt.Println("=        Medical Check Up       =")
		fmt.Println("=================================")
		fmt.Println("1. Tambah")
		fmt.Println("2. Hapus")
		fmt.Println("3. Edit")
		fmt.Println("----------------")
		fmt.Println("Pilihan anda : ")
		fmt.Scan(&i)
			if i == 1{
				tambahData(&mcu)
			}else if i == 2 {
				printPasien(mcu)
				fmt.Println("----------------")
				hapusData(&mcu, &indexPasien, id)
				printPasien(mcu)
				fmt.Println("")
		}else if i == 3 {
			fmt.Println("=================================")
			fmt.Println("=        Medical Check Up       =")
			fmt.Println("=================================")
			fmt.Println("1. Nama")
			fmt.Println("2. Tanggal")
			fmt.Println("3. Paket")
			fmt.Println("----------------")
			fmt.Println("Pilihan anda : ")
			fmt.Scan(&i)
			if i == 1 {
				printPasien(mcu)
				editNamaPasien(&mcu, indexPasien, id, namaEdit)
				fmt.Println("")
				printPasien(mcu)
				fmt.Println("Data berhasil diubah")
				fmt.Println("")
			}else if i == 2 {
				printPasien(mcu)
				editTanggalPasien(&mcu, indexPasien, id, namaEdit)
				fmt.Println("")
				printPasien(mcu)
				fmt.Println("Data berhasil diubah")
				fmt.Println("")
			}else if i == 3 {
				printRekap(mcu)
				editPaketPasien(&mcu, indexPasien, id, namaEdit)
				fmt.Println("")
				printRekap(mcu)
				fmt.Println("Data berhasil diubah")
				fmt.Println("")
			}
		}
	}else if i == 2 {
		fmt.Println("=================================")
		fmt.Println("=        Medical Check Up       =")
		fmt.Println("=================================")
		fmt.Println("1. Paket")
		fmt.Println("2. Pasien")
		fmt.Println("3. Rekap")
		fmt.Println("----------------")
		fmt.Println("Pilihan anda : ")
		fmt.Scan(&i)
		if i == 1 { //menampilkan data pasien
			printPaket(mcu)
		}
		if i == 2 { //menampilkan data pasien
			printPasien(mcu)
		}
		if i == 3 {
			fmt.Println("=================================")
			fmt.Println("=        Medical Check Up       =")
			fmt.Println("=================================")
			fmt.Println("1. Menurut Tanggal")
			fmt.Println("2. Menurut Paket yang dipilih")
			fmt.Println("3. Tampilkan")
			fmt.Println("----------------")
			fmt.Println("Pilihan anda : ")
			fmt.Scan(&i)
			if i == 1 {
				fmt.Println("=================================")
				fmt.Println("=        Medical Check Up       =")
				fmt.Println("=================================")
				fmt.Println("1. Urutkan secara Descending(Besar ke kecil)")
				fmt.Println("2. Urutkan secara Ascending(Kecil ke besar)")
				fmt.Println("----------------")
				fmt.Println("Pilihan anda : ")
				fmt.Scan(&i)
					if i == 1 {
						selectionSortPasien(&mcu)
					}else if i == 2 {
						insertionSortPasien(&mcu)
					}
				}
			if i == 2 {
				sortPaket(&mcu)
			}
			if i == 3 {
				printRekap(mcu)
			}
		}
	}else if i == 3 {
		fmt.Println("=================================")
		fmt.Println("=        Medical Check Up       =")
		fmt.Println("=================================")
		fmt.Println("1. Pencarian berdasarkan nama")
		fmt.Println("2. Pencarian berdasarkan tanggal")
		fmt.Println("3. Pencarian berdasarkan paket")
		fmt.Println("----------------")
		fmt.Println("Pilihan anda : ")
		fmt.Scan(&i)
		if i == 1 {
			cariPasien(mcu, nama, indexPasien )
		}
		if i == 2 {
			cariTanggal(mcu, nama, indexPasien)
		}
		if i == 3 {
			cariPaket(mcu, nama, indexPasien )
		}
	}else if i == 4 {
			break
		}
	}
}
func tambahData(A *tabMCU){
	var harga,i int 

	if indexPasien <= NMAX{
		fmt.Println("Masukkan nama pasien yang ingin didaftarkan:")
		fmt.Scan(&A[indexPasien].namaPasien)
		fmt.Println("Masukkan tanggal pasien ingin melakukan MCU :")
		fmt.Scan(&A[indexPasien].tanggal)
		A[indexPasien].ID = nextId
		nextId++
		indexPasien++
		}
	fmt.Println("")
	fmt.Println("Pilih paket anda :")
	fmt.Println("1. Basic - Rp 500.000")
	fmt.Println("2. Plus - Rp 750.000 ")
	fmt.Println("3. Premium - Rp 1.000.000")
	fmt.Println("----------------")
	fmt.Println("Pilihan anda : ")
	fmt.Scan(&i)
		if i == 1 {
			harga =  500000
			A[indexPaket].namaPaket = "Basic"
			A[indexPaket].harga = harga
		}else if i == 2 {
			harga =  750000
			A[indexPaket].namaPaket = "Plus"
			A[indexPaket].harga = harga
		}else if i == 3 {
			harga =  1000000		
			A[indexPaket].namaPaket = "Premium"
			A[indexPaket].harga = harga
		}
		fmt.Println("Data berhasil ditambahkan")
		fmt.Println("")
		indexPaket++
		for i := 0 ; i < indexPaket ; i++{
			A[i].ID = i + 1
		}
}
//procedure untuk meng-output kan array pasien 
func printPasien(A tabMCU){
	fmt.Println("Daftar Data Pasien:")
	fmt.Printf("%2s %12s %12s \n" ,"ID", "Nama", "Tanggal" )
	for i:= 0 ; i < indexPasien ; i++{
		fmt.Printf("%2d %12s %12s \n",A[i].ID,A[i].namaPasien,A[i].tanggal)
	}
}
//procedure untuk meng-output kan array paket 
func printPaket(A tabMCU){
    fmt.Println("Daftar paket:")
	fmt.Printf("%2s %12s %12s %12s \n" ,"ID", "Nama", "Paket", "Harga" )
	for i:= 0 ; i < indexPaket ; i++{
		fmt.Printf("%2d %12s %12s %12d \n",A[i].ID,A[i].namaPasien,A[i].namaPaket,A[i].harga)
	}

}//procedure untuk meng-output kan array pasien dan paket 
func printRekap(A tabMCU){
	var totalHarga int
    fmt.Println("Daftar rekap:")
	fmt.Printf("%2s %12s %12s %12s %12s \n" ,"ID", "Nama", "Paket", "Tanggal", "Harga")
    for i := 0 ; i < indexPaket && i < indexPasien ; i++ {
        fmt.Printf("%2d %12s %12s %12s %12d \n", A[i].ID,A[i].namaPasien, A[i].namaPaket,A[i].tanggal,A[i].harga)
		totalHarga += A[i].harga
    }
	fmt.Println()
	fmt.Println("Total pemasukan layanan MCU :", totalHarga)
}
//procedure untuk edit nama pasien
func editNamaPasien(A *tabMCU, indexPasien int, idCari int, namaEdit string ){
	var found bool
	var i int
	found = false
	fmt.Println("Masukkan ID pasien yang ingin dicari:")
	fmt.Scan(&idCari)
	fmt.Println("Masukkan nama pasien yang ingin diubah:")
	fmt.Scan(&namaEdit)
	for i = 0; i < indexPasien ; i++{
		if A[i].ID == idCari {
			A[i].namaPasien = namaEdit 
			found = true
		}
	}
	if found == true {
		fmt.Println("Data berhasil diubah")
	} else if found == false {
		fmt.Println("Nama pasien yang anda masukkan tidak ditemukan")
		fmt.Println("Data tidak berhasil diubah")
	}
}
//procedure untuk edit tanggal pasien
func editTanggalPasien(A *tabMCU, indexPasien int, idCari int, tanggalEdit string ){
	var found bool
	var i int
	found = false
	fmt.Println("Masukkan ID pasien yang ingin dicari:")
	fmt.Scan(&idCari)
	fmt.Println("Masukkan Tanggal pasien yang ingin diubah:")
	fmt.Scan(&tanggalEdit)
	for i = 0; i < indexPasien ; i++{
		if A[i].ID == idCari {
			A[i].tanggal = tanggalEdit 
			found = true
		}
	}
	if found == true {
		fmt.Println("Data berhasil diubah")
	} else if found == false {
		fmt.Println("Nama pasien yang anda masukkan tidak ditemukan")
		fmt.Println("Data tidak berhasil diubah")
	}
}
//procedure untuk edit paket pasien
func editPaketPasien(A *tabMCU, indexPasien int, idCari int, paketEdit string ){
	var found bool
	var i,harga int
	found = false
	fmt.Println("Masukkan ID pasien yang ingin dicari:")
	fmt.Scan(&idCari)
	fmt.Println("")
	fmt.Println("Pilih paket untuk mengedit paket anda :")
	fmt.Println("1. Basic - Rp 500.000")
	fmt.Println("2. Plus - Rp 750.000 ")
	fmt.Println("3. Premium - Rp 1.000.000")
	fmt.Println("----------------")
	fmt.Println("Pilihan anda : ")
	fmt.Scan(&i)
		if i == 1 {
			harga =  500000
			paketEdit = "Basic"
		}else if i == 2 {
			harga =  750000
			paketEdit = "Plus"
		}else if i == 3 {
			harga =  1000000		
			paketEdit = "Premium"
		}
	for i = 0; i < indexPasien ; i++{
		if A[i].ID == idCari {
			A[i].namaPaket = paketEdit 
			A[i].harga = harga
			found = true
		}
	}
	if found == true {
		fmt.Println("Data berhasil diubah")
	} else if found == false {
		fmt.Println("Nama pasien yang anda masukkan tidak ditemukan")
		fmt.Println("Data tidak berhasil diubah")
		}

}
//function untuk mencari index pasien
func findEdit(A tabMCU, indexPasien int, idCari int)int{ //sequential search digunakan untuk mencari indeks pasien yang akan diedit
	var found int
	for i:= 0; i<indexPasien;i++{
		if A[i].ID == idCari{
			found = i
		} else{
			found = -1
		}
	}
	return found
}
//procedure untuk hapus nama pasien
func hapusData(A *tabMCU, indexPasien *int, ID int){
	var found int
	found = findHapusData(*A, *indexPasien, ID)
	if found != -1{
		*indexPasien = *indexPasien-1
		for found < *indexPasien{
			A[found].namaPasien = A[found+1].namaPasien
			A[found].ID = A[found+1].ID
			found++	
		} 		
	}else {
		fmt.Println("Data not found")
	}
	for i := 0 ; i < *indexPasien ; i++{
		A[i].ID = i + 1
	}
}
//function untuk mencari index pasien
func findHapusData(A tabMCU, indexPasien int, ID int)int{ //binary search digunakan untuk mencari indeks pasien yang akan dihapus
	var left, right, mid int
	right = indexPasien-1
	fmt.Println("Masukkan ID pasien yang ingin dihapus:")
	fmt.Scan(&ID)
	for left <= right{
		mid = (left+right) / 2
		if A[mid].ID < ID {
			left = mid + 1
		} else if A[mid].ID > ID{
			right = mid -1
		} else {
			return mid
		}
	}
	return -1
}
//procedure untuk mencari array namaPasien
func cariPasien(A tabMCU, nama string, indexPasien int){ 
	var i int
	fmt.Println("Masukkan nama yang ingin dicari")
	for i < indexPasien {
		fmt.Scan(&nama)
		for j := 0 ; j <= indexPasien-1 ; j++{
			if nama == A[j].namaPasien{
				fmt.Println("Nama ditemukan!")
				fmt.Printf("%2s %12s %12s %12s %12s \n" ,"ID", "Nama", "Paket", "Tanggal", "Harga")
				fmt.Printf("%2d %12s %12s %12s %12d \n", A[j].ID,A[j].namaPasien, A[j].namaPaket,A[j].tanggal,A[j].harga)
				fmt.Println("")
				break 
			}
			if j == indexPasien-1 && nama != A[j].namaPasien{
				fmt.Println("Nama tidak ditemukan!")
				break
			}
			if nama == "4"{
				break
			}
		}
		if nama == "4"{
			break
		}
		fmt.Println("Ketik 4 jika ingin berhenti mencari nama pasien")
	}
}
//procedure untuk mencari array tanggal
func cariTanggal(A tabMCU, nama string, indexPasien int){ 
	var i int
	fmt.Println("Masukkan tanggal yang ingin dicari:")
	for i < indexPasien {
		fmt.Scan(&nama)
		fmt.Printf("%2s %12s %12s %12s %12s \n" ,"ID", "Nama", "Paket", "Tanggal", "Harga")
		for j := 0 ; j <= indexPasien-1 ; j++{
			if nama == A[j].tanggal{
				fmt.Printf("%2d %12s %12s %12s %12d \n", A[j].ID,A[j].namaPasien, A[j].namaPaket,A[j].tanggal,A[j].harga)
				fmt.Println("Tanggal ditemukan!")
			}
			if j == indexPasien && nama != A[j].tanggal{
				fmt.Println("Tanggal tidak ditemukan!")
				break
			}
			if nama == "4"{
				break
			}
		}
		if nama == "4"{
			break
		}
		fmt.Println("Ketik 4 jika ingin berhenti mencari tanggal")
		fmt.Println("")
	}
}
//procedure untuk mencari array namaPaket
func cariPaket(A tabMCU, nama string, indexPasien int){ 
	var i int
	fmt.Println("Masukkan nama yang ingin dicari")
	for i < indexPaket {
		fmt.Println("Berikut list paket yang ingin ditemukan:")
		fmt.Println("1. Basic")
		fmt.Println("2. Plus")
		fmt.Println("3. Premium")
		fmt.Println("-------------")
		fmt.Println("4. Exit")
		fmt.Println("-------------")
		fmt.Println("Pilihan anda:")
		
		fmt.Scan(&nama)
		fmt.Printf("%2s %12s %12s %12s %12s \n" ,"ID", "Nama", "Paket", "Tanggal", "Harga")
		for j := 0 ; j <= indexPaket-1 ; j++{
			if nama == A[j].namaPaket {
				fmt.Printf("%2d %12s %12s %12s %12d \n", A[j].ID,A[j].namaPasien, A[j].namaPaket,A[j].tanggal,A[j].harga)
				fmt.Println("")
				fmt.Println("Paket ditemukan!")
			}else if j == indexPaket && nama != A[j].namaPaket {
				fmt.Println("")
				fmt.Println("Paket tidak ditemukan!")
				break
			}else if nama == "4" {
				break
			}
		}
		if nama == "4" {
			break
		}
		fmt.Println("")
	}
}
//procedure untuk sorting descending
func selectionSortPasien(A *tabMCU) {
    tempIndexPasien := indexPasien
    tempA := *A // buat variabel temporary
    for i := 1; i <= tempIndexPasien-1;  {
        Idx := i-1
        for j := i-1; j < tempIndexPasien; {
            if tempA[j].tanggal > tempA[Idx].tanggal {
                Idx = j
            }
			j++
        }
        temp := tempA[Idx]
		tempA[Idx] = tempA[i-1]
		tempA[i-1] = temp
		i++
    }
    printPasien(tempA) // cetak hasil sorting
}
//procedure untuk sorting ascending
func insertionSortPasien(A *tabMCU){
    tempIndexPasien := indexPasien
	tempA := *A 
    for i := 1 ; i <= tempIndexPasien-1;{
		j := i
        temp := tempA[j]
        for j > 0 && temp.tanggal < tempA[j-1].tanggal  {
            tempA[j] = tempA[j-1]
            j--
        }
        tempA[j] = temp //penyisipan array
		i++
    }
    printPasien(tempA) // cetak hasil sorting
}
//procedure untuk mengurutkan namaPaket
func sortPaket(A *tabMCU){
    tempIndexPaket := indexPaket
	tempA := *A 
    for i := 1 ; i <= tempIndexPaket-1;{
		j := i
        temp := tempA[j]
        for j > 0 && (temp.namaPaket == "Basic" && tempA[j-1].namaPaket == "Plus" || temp.namaPaket == "Basic" && tempA[j-1].namaPaket == "Premium" || temp.namaPaket == "Plus" && tempA[j-1].namaPaket == "Premium" || temp.namaPaket < tempA[j-1].namaPaket) {
            tempA[j] = tempA[j-1]
            j--
        }
        tempA[j] = temp //penyisipan array
		i++
    }
    printPaket(tempA) // cetak hasil sorting
} 