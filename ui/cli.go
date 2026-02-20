package ui

import (
	"bufio"
	"fmt"
	"koda-b6-weekly1-go/modules"
	"os"
	"strconv"
	"strings"
)

func StartApp() {
	store, err := modules.LoadData("data/menu.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	menuModule := modules.NewMenuModule(store.Categories, store.Menus)
	cartModule := modules.NewCartModule()
	historyModule := modules.NewHistoryModule()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nSelamat Datang Di McD")
		fmt.Println("1. Lihat semua makanan")
		fmt.Println("2. Lihat keranjang")
		fmt.Println("3. Checkout")
		fmt.Println("4. Lihat history pembelian")
		fmt.Println("0. Keluar")

		fmt.Print("\nPilih menu (0-4): ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Input harus angka!")
			continue
		}

		switch choice {
		case 1:
			showAllMenus(menuModule, cartModule, reader)

		case 2:
			showCart(cartModule)

		case 3:
			checkout(cartModule, historyModule)

		case 4:
			showHistory(historyModule)

		case 0:
			fmt.Println("Terima kasih sudah datang ke McD!")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func showAllMenus(menuModule *modules.MenuModule, cartModule *modules.CartModule, reader *bufio.Reader) {
	for {
		fmt.Println("\nDAFTAR MENU McD")

		for i := 0; i < len(menuModule.Categories); i++ {
			category := menuModule.Categories[i]
			categoryName := menuModule.GetCategoryName(category.ID)

			fmt.Println(categoryName)

			menus := menuModule.GetMenusByCategory(category.ID)
			for j := 0; j < len(menus); j++ {
				fmt.Printf("%d. %s - Rp%d\n", menus[j].ID, menus[j].Name, menus[j].Price)
			}

			fmt.Println()
		}

		fmt.Print("Masukkan ID menu yang ingin dipesan (0 untuk kembali): ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)

		menuID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Input harus angka!")
			continue
		}

		if menuID == 0 {
			return
		}

		menu, found := menuModule.GetMenuByID(menuID)
		if !found {
			fmt.Println("Menu tidak ditemukan!")
			continue
		}

		err = cartModule.AddToCart(menu)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Berhasil ditambahkan ke keranjang!")
		}
	}
}

func showCart(cartModule *modules.CartModule) {
	fmt.Println("\nKERANJANG BELANJA")

	if cartModule.IsEmpty() {
		fmt.Println("Keranjang masih kosong.")
		return
	}

	for i := 0; i < len(cartModule.Items); i++ {
		item := cartModule.Items[i]
		fmt.Printf("%d. %s | Qty: %d | Rp%d\n", item.Menu.ID, item.Menu.Name, item.Qty, item.SubTotal)
	}

	fmt.Println("\n-------------------------")
	fmt.Println("TOTAL BELANJA: Rp", cartModule.GetTotal())
	fmt.Println("-------------------------")
}

func checkout(cartModule *modules.CartModule, historyModule *modules.HistoryModule) {
	fmt.Println("\nINVOICE PEMBELIAN")

	if cartModule.IsEmpty() {
		fmt.Println("Keranjang masih kosong, tidak bisa checkout.")
		return
	}

	for i := 0; i < len(cartModule.Items); i++ {
		item := cartModule.Items[i]
		fmt.Printf("- %s | Qty: %d | Rp%d\n", item.Menu.Name, item.Qty, item.SubTotal)
	}

	total := cartModule.GetTotal()

	fmt.Println("\n-------------------------")
	fmt.Println("TOTAL BAYAR: Rp", total)
	fmt.Println("-------------------------")

	err := historyModule.AddTransaction(cartModule.Items, total)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cartModule.ClearCart()
	fmt.Println("Checkout berhasil! Transaksi masuk ke history.")
}

func showHistory(historyModule *modules.HistoryModule) {
	fmt.Println("\nHISTORY PEMBELIAN")

	if historyModule.IsEmpty() {
		fmt.Println("Belum ada transaksi checkout.")
		return
	}

	history := historyModule.GetAll()

	for i := 0; i < len(history); i++ {
		trx := history[i]

		fmt.Printf("\nTransaksi #%d (%s)\n", trx.ID, trx.CreatedAt.Format("2006-01-02 15:04:05"))
		for j := 0; j < len(trx.Items); j++ {
			item := trx.Items[j]
			fmt.Printf("- %s | Qty: %d | Rp%d\n", item.Menu.Name, item.Qty, item.SubTotal)
		}

		fmt.Println("TOTAL:", trx.Total)
		fmt.Println("-------------------------")
	}
}