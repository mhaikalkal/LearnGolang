package mapping

import "fmt"

func MapFunc() {
	// Cara 1
	var organization map[string]int
	organization = map[string]int{}

	organization["PT Narcistic"] = 1
	organization["PT Sampoerna Tbk."] = 245
	organization["Alphabet Corp."] = 45618

	fmt.Println(organization)
	fmt.Println(organization["PT Narcistic"])
	fmt.Println(organization["PT Sampoerna Tbk."])
	fmt.Println(organization["Alphabet Corp."])


	/*
	// Cara Horizontal
	var koin1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var koin2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	var koin4 = make(map[string]int)
	var koin5 = *new(map[string]int) * = pointer, nanti dibahas lebih lanjutnya

	*/

	// Fusion map + slice
	var games = []map[string]string{
		{
			"id": "1", 
			"name": "Minecraft", 
			"genre": "Adventure",
		},
		{
			"id": "2", 
			"name": "The Elder Scrolls V: Skyrim", 
			"genre": "Action",
		},
	}

	for _, game := range games {
		fmt.Printf("Game ke-%s dengan genre %s yang berjudul %s.\n", game["id"], game["genre"], game["name"])

	}


	var hardwares = []map[string]string{
		{"id": "1", "vga": "RTX 3080", "keyboard": "Logitech MX Series", "mouse": "logitech G103 Prodigy"},
		{"id": "2", "monitor": "Samsung", "ssd": "WD Black 1 TB"},
		{"id": "3", "headphone": "Rexus Vonix", "monitor": "LG", "ssd": "WD Green"},

	}

	for _, hardware := range hardwares {
		fmt.Printf(hardware["ssd"])
	}

}