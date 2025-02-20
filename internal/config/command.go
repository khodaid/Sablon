package config

import "flag"

var Migrate bool // Flag untuk migrasi
var Seed bool    // Flag untuk seeder

// Fungsi untuk inisialisasi flag
func init() {
	// Definisikan flag untuk migrasi database
	flag.BoolVar(&Migrate, "migrate", false, "Run the auto-migration")
	// flag.Parse() // Parse flag dari command line

	// Definisikan flag untuk seeder database
	flag.BoolVar(&Seed, "seed", false, "Run the auto-seeder")
	flag.Parse() // Parse flag dari command line
}
