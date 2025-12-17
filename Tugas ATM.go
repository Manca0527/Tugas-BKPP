package main

import "fmt"

type Akun struct {
	Akun  string
	Pin   int
	saldo float64
}

func main() {
	var dataAkun []Akun = []Akun{
		{"Ilyas", 1111, 200000},
		{"Ali", 2222, 250000},
	}

	fmt.Println("=== Selamat Datang di ATM ===")

	var Login *Akun
	var namaAkun string
	var pin int

	fmt.Println("\n=== Login ATM ===")
	fmt.Print("Masukan Nama Akun: ")
	fmt.Scanln(&namaAkun)

	fmt.Print("Masukan Pin: ")
	fmt.Scanln(&pin)

	var loginBerhasil bool = false

	var i int
	for i = 0; i < len(dataAkun); i++ {
		if dataAkun[i].Akun == namaAkun && dataAkun[i].Pin == pin {
			Login = &dataAkun[i]
			loginBerhasil = true
			break
		}
	}

	if loginBerhasil {
		fmt.Println("Login Berhasil")

		var isRunning bool = true

		for isRunning {
			var Pilihan int

			fmt.Println("\n=== Menu ATM ===")
			fmt.Println("1. Setor Tunai")
			fmt.Println("2. Tarik Tunai")
			fmt.Println("3. Transfer")
			fmt.Println("4. Cek Saldo")
			fmt.Println("5. Exit")
			fmt.Print("Pilih Menu (1-5): ")
			fmt.Scanln(&Pilihan)

			if Pilihan == 1 {
				var jumlah float64
				fmt.Print("Masukan Jumlah Setor Tunai: Rp")
				fmt.Scanln(&jumlah)

				var jumlahValid bool = jumlah > 0
				if !jumlahValid {
					fmt.Println("Jumlah Tidak Valid!!!")
				} else {
					Login.saldo = Login.saldo + jumlah
					fmt.Printf("Setoran berhasil! Saldo Baru: Rp%.2f\n", Login.saldo)
				}

			} else if Pilihan == 2 {
				var jumlah float64
				fmt.Print("Masukan Jumlah Tarik Tunai: Rp")
				fmt.Scanln(&jumlah)

				var jumlahValid bool = jumlah > 0
				var saldoCukup bool = jumlah <= Login.saldo

				if !jumlahValid {
					fmt.Println("Jumlah Tidak Valid!!!")
				} else if !saldoCukup {
					fmt.Println("Saldo Tidak Cukup!!!")
				} else {
					Login.saldo = Login.saldo - jumlah
					fmt.Printf("Penarikan Berhasil! Saldo Baru: Rp%.2f\n", Login.saldo)
				}

			} else if Pilihan == 3 {
				var NamaAkun string
				var jumlah float64

				fmt.Print("Masukan Nama Akun Tujuan: ")
				fmt.Scanln(&NamaAkun)

				fmt.Print("Masukan Jumlah Transfer: Rp")
				fmt.Scanln(&jumlah)

				var akunTujuan *Akun
				var akunDitemukan bool = false

				// GANTI: for i := range dataAkun { ... } -> loop indeks tanpa :=
				var j int
				for j = 0; j < len(dataAkun); j++ {
					if dataAkun[j].Akun == NamaAkun {
						akunTujuan = &dataAkun[j]
						akunDitemukan = true
						break
					}
				}

				var jumlahValid bool = jumlah > 0
				var saldoCukup bool = jumlah <= Login.saldo

				if !akunDitemukan {
					fmt.Println("Akun Tujuan Tidak Ditemukan!!!")
				} else if !jumlahValid {
					fmt.Println("Jumlah Tidak Valid!!!")
				} else if !saldoCukup {
					fmt.Println("Saldo Tidak Cukup!!!")
				} else {
					Login.saldo = Login.saldo - jumlah
					akunTujuan.saldo = akunTujuan.saldo + jumlah
					fmt.Printf("Transfer Berhasil! Saldo Baru: Rp%.2f\n", Login.saldo)
				}

			} else if Pilihan == 4 {
				fmt.Printf("Saldo Anda: Rp%.2f\n", Login.saldo)

			} else if Pilihan == 5 {
				isRunning = false
				fmt.Println("\nLogout berhasil!")
				fmt.Println("\n=================================")
				fmt.Println("   TERIMA KASIH TELAH GUNAKAN   ")
				fmt.Println("           ATM KAMI             ")
				fmt.Println("=================================")

			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		}

	} else {
		fmt.Println("Login Gagal!!!")
	}
}
