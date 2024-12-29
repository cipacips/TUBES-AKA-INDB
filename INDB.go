package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type show struct {
	title, genre, status string
	seasons              int
	rating               float32
	isMovie              bool
}

var tvShowList = []show{
	{"Stranger Things", "Sci-Fi", "Belum Ditonton", 4, 8.7, false},
	{"The Crown", "Drama", "Belum Ditonton", 5, 8.6, false},
	{"Money Heist", "Action", "Belum Ditonton", 5, 8.3, false},
	{"The Witcher", "Fantasy", "Belum Ditonton", 2, 8.2, false},
	{"Bridgerton", "Romance", "Belum Ditonton", 2, 7.3, false},
	{"Ozark", "Thriller", "Belum Ditonton", 4, 8.4, false},
	{"The Queen's Gambit", "Drama", "Belum Ditonton", 1, 8.6, false},
	{"Jessica Jones", "Action, Crime, Drama", "Belum Ditonton", 3, 7.9, false},
	{"Luke Cage", "Action, Crime, Drama", "Belum Ditonton", 2, 7.3, false},
	{"Iron Fist", "Action, Adventure, Crime", "Belum Ditonton", 2, 6.5, false},
	{"The Defenders", "Action, Adventure, Crime", "Belum Ditonton", 1, 7.3, false},
	{"The Umbrella Academy", "Action, Adventure, Comedy", "Belum Ditonton", 3, 8.0, false},
	{"13 Reasons Why", "Drama, Mystery, Thriller", "Belum Ditonton", 4, 7.6, false},
	{"Orange Is the New Black", "Comedy, Crime, Drama", "Belum Ditonton", 7, 8.1, false},
	{"Sex Education", "Comedy, Drama", "Belum Ditonton", 3, 8.3, false},
	{"The Haunting of Hill House", "Drama, Horror, Mystery", "Belum Ditonton", 1, 8.6, false},
	{"Master of None", "Comedy, Drama", "Belum Ditonton", 3, 8.3, false},
	{"Russian Doll", "Adventure, Comedy, Drama", "Belum Ditonton", 2, 7.9, false},
	{"Altered Carbon", "Action, Drama, Sci-Fi", "Belum Ditonton", 2, 8.0, false},
	{"Lost in Space", "Adventure, Drama, Family", "Belum Ditonton", 3, 7.3, false},
	{"The OA", "Drama, Fantasy, Mystery", "Belum Ditonton", 2, 7.8, false},
	{"Sense8", "Drama, Mystery, Sci-Fi", "Belum Ditonton", 2, 8.3, false},
	{"The Society", "Drama, Mystery, Sci-Fi", "Belum Ditonton", 1, 7.1, false},
	{"Chilling Adventures of Sabrina", "Drama, Fantasy, Horror", "Belum Ditonton", 4, 7.5, false},
	{"Locke & Key", "Drama, Fantasy, Horror", "Belum Ditonton", 3, 7.4, false},
	{"Shadowhunters", "Action, Drama, Fantasy", "Belum Ditonton", 3, 6.6, false},
	{"The Innocents", "Drama, Fantasy, Mystery", "Belum Ditonton", 1, 6.3, false},
	{"Daybreak", "Action, Adventure, Comedy", "Belum Ditonton", 1, 6.7, false},
	{"Raising Dion", "Drama, Sci-Fi", "Belum Ditonton", 2, 7.2, false},
	{"The Politician", "Comedy, Drama", "Belum Ditonton", 2, 7.5, false},
	{"Santa Clarita Diet", "Comedy, Horror", "Belum Ditonton", 3, 7.8, false},
	{"Disenchantment", "Animation, Action, Adventure", "Belum Ditonton", 3, 7.2, false},
	{"The Dragon Prince", "Animation, Action, Adventure", "Belum Ditonton", 3, 8.4, false},
	{"Carmen Sandiego", "Animation, Action, Adventure", "Belum Ditonton", 4, 7.9, false},
	{"She-Ra and the Princesses of Power", "Animation, Action, Adventure", "Belum Ditonton", 5, 7.9, false},
	{"Voltron: Legendary Defender", "Animation, Action, Adventure", "Belum Ditonton", 8, 8.1, false},
	{"Castlevania", "Animation, Action, Adventure", "Belum Ditonton", 4, 8.2, false},
	{"Aggretsuko", "Animation, Comedy, Drama", "Belum Ditonton", 4, 8.0, false},
	{"Hilda", "Animation, Adventure, Comedy", "Belum Ditonton", 2, 8.6, false},
	{"Kipo and the Age of Wonderbeasts", "Animation, Action, Adventure", "Belum Ditonton", 3, 8.3, false},
	{"The Hollow", "Animation, Action, Adventure", "Belum Ditonton", 2, 7.2, false},
	{"Trollhunters: Tales of Arcadia", "Animation, Action, Adventure", "Belum Ditonton", 3, 8.4, false},
	{"3Below: Tales of Arcadia", "Animation, Action, Adventure", "Belum Ditonton", 2, 7.7, false},
	{"Wizards: Tales of Arcadia", "Animation, Action, Adventure", "Belum Ditonton", 1, 8.2, false},
	{"The Dark Crystal: Age of Resistance", "Adventure, Drama, Family", "Belum Ditonton", 1, 8.4, false},
	{"Green Eggs and Ham", "Animation, Adventure, Comedy", "Belum Ditonton", 1, 8.2, false},
	{"Arcane", "Animation, Action, Adventure", "Belum Ditonton", 1, 9.0, false},
	{"Squid Game", "Drama, Thriller", "Belum Ditonton", 1, 8.5, false},
	{"The Sandman", "Fantasy, Drama", "Belum Ditonton", 1, 7.9, false},
	{"Sweet Tooth", "Drama, Fantasy", "Belum Ditonton", 2, 8.0, false},
	{"Alice in Borderland", "Action, Sci-Fi", "Belum Ditonton", 2, 7.6, false},
	{"Dark", "Drama, Sci-Fi", "Belum Ditonton", 3, 8.8, false},
	{"Lupin", "Action, Mystery", "Belum Ditonton", 3, 7.5, false},
	{"The End of the F***ing World", "Drama, Comedy", "Belum Ditonton", 2, 8.1, false},
	{"Mindhunter", "Crime, Thriller", "Belum Ditonton", 2, 8.6, false},
	{"Narcos", "Drama, Crime", "Belum Ditonton", 3, 8.8, false},
	{"The Haunting of Bly Manor", "Drama, Horror", "Belum Ditonton", 1, 7.3, false},
	{"The Punisher", "Action, Crime", "Belum Ditonton", 2, 8.1, false},
	{"Daredevil", "Action, Crime", "Belum Ditonton", 3, 8.6, false},
	{"The Irregulars", "Drama, Fantasy", "Belum Ditonton", 1, 5.9, false},
	{"Shadow and Bone", "Drama, Fantasy", "Belum Ditonton", 2, 7.6, false},
	{"Fate: The Winx Saga", "Fantasy, Drama", "Belum Ditonton", 2, 6.9, false},
	{"Warrior Nun", "Action, Fantasy", "Belum Ditonton", 2, 7.0, false},
	{"The Order", "Drama, Horror", "Belum Ditonton", 2, 6.8, false},
	{"Pacific Rim: The Black", "Animation, Action", "Belum Ditonton", 2, 7.2, false},
	{"Blood of Zeus", "Animation, Action", "Belum Ditonton", 1, 7.5, false},
	{"Love, Death & Robots", "Animation, Sci-Fi", "Belum Ditonton", 3, 8.4, false},
	{"Trese", "Animation, Fantasy", "Belum Ditonton", 1, 7.1, false},
	{"Yasuke", "Animation, Action", "Belum Ditonton", 1, 6.2, false},
	{"Resident Evil: Infinite Darkness", "Animation, Horror", "Belum Ditonton", 1, 5.8, false},
	{"DOTA: Dragon's Blood", "Animation, Fantasy", "Belum Ditonton", 3, 7.7, false},
	{"BoJack Horseman", "Animation, Comedy", "Belum Ditonton", 6, 8.7, false},
	{"Big Mouth", "Animation, Comedy", "Belum Ditonton", 7, 7.9, false},
	{"F Is for Family", "Animation, Comedy", "Belum Ditonton", 5, 8.0, false},
	{"Black Mirror", "Drama, Sci-Fi", "Belum Ditonton", 5, 8.8, false},
	{"The Seven Deadly Sins", "Animation, Action", "Belum Ditonton", 5, 7.8, false},
	{"Baki", "Animation, Action", "Belum Ditonton", 3, 7.4, false},
	{"Kengan Ashura", "Animation, Action", "Belum Ditonton", 2, 7.7, false},
	{"Record of Ragnarok", "Animation, Action", "Belum Ditonton", 2, 6.7, false},
	{"Beastars", "Animation, Drama", "Belum Ditonton", 2, 8.0, false},
	{"Devilman Crybaby", "Animation, Horror", "Belum Ditonton", 1, 7.9, false},
	{"Great Pretender", "Animation, Adventure", "Belum Ditonton", 2, 8.3, false},
	{"Disenchantment Part 4", "Animation, Comedy", "Belum Ditonton", 4, 7.2, false},
	{"She-Ra (2024 Reboot)", "Animation, Action", "Belum Ditonton", 1, 8.0, false},
	{"Tales of the Teenage Mutant Ninja Turtles", "Animation, Action", "Belum Ditonton", 3, 7.4, false},
	{"Dragon Age: Absolution", "Animation, Fantasy", "Belum Ditonton", 1, 7.6, false},
	{"Kung Fu Panda: The Dragon Knight", "Animation, Comedy", "Belum Ditonton", 2, 7.1, false},
	{"Inside Job", "Animation, Comedy", "Belum Ditonton", 2, 7.6, false},
	{"Q-Force", "Animation, Comedy", "Belum Ditonton", 1, 6.3, false},
	{"Cannon Busters", "Animation, Action", "Belum Ditonton", 1, 6.7, false},
	{"High Score Girl", "Animation, Comedy", "Belum Ditonton", 2, 7.9, false},
	{"Neo Yokio", "Animation, Fantasy", "Belum Ditonton", 1, 6.0, false},
	{"Seis Manos", "Animation, Action", "Belum Ditonton", 1, 7.2, false},
	{"The Cuphead Show!", "Animation, Comedy", "Belum Ditonton", 3, 7.6, false},
	{"Carmen Sandiego (2024)", "Animation, Action", "Belum Ditonton", 1, 8.0, false},
	{"Cyberpunk: Edgerunners", "Animation, Sci-Fi", "Belum Ditonton", 1, 8.6, false},
}

var movieList = []show{
	{"The King", "Biography, Drama, History", "Belum Ditonton", 1, 7.3, true},
	{"The Irishman", "Biography, Crime, Drama", "Belum Ditonton", 1, 7.8, true},
	{"Marriage Story", "Drama, Romance", "Belum Ditonton", 1, 7.9, true},
	{"6 Underground", "Action, Thriller", "Belum Ditonton", 1, 6.1, true},
	{"The Two Popes", "Biography, Comedy, Drama", "Belum Ditonton", 1, 7.6, true},
	{"Extraction", "Action, Thriller", "Belum Ditonton", 1, 6.7, true},
	{"The Old Guard", "Action, Adventure, Fantasy", "Belum Ditonton", 1, 6.6, true},
	{"Project Power", "Action, Crime, Sci-Fi", "Belum Ditonton", 1, 6.0, true},
	{"Enola Holmes", "Adventure, Crime, Drama", "Belum Ditonton", 1, 6.6, true},
	{"The Trial of the Chicago 7", "Drama, History, Thriller", "Belum Ditonton", 1, 7.8, true},
	{"Mank", "Biography, Drama", "Belum Ditonton", 1, 6.8, true},
	{"The Midnight Sky", "Drama, Fantasy, Sci-Fi", "Belum Ditonton", 1, 5.6, true},
	{"Outside the Wire", "Action, Adventure, Fantasy", "Belum Ditonton", 1, 5.4, true},
	{"Malcolm & Marie", "Drama, Romance", "Belum Ditonton", 1, 6.7, true},
	{"Army of the Dead", "Action, Crime, Horror", "Belum Ditonton", 1, 5.7, true},
	{"Fatherhood", "Comedy, Drama", "Belum Ditonton", 1, 6.6, true},
	{"Sweet Girl", "Action, Drama, Thriller", "Belum Ditonton", 1, 5.5, true},
	{"Red Notice", "Action, Comedy, Thriller", "Belum Ditonton", 1, 6.3, true},
	{"Don't Look Up", "Comedy, Drama, Sci-Fi", "Belum Ditonton", 1, 7.2, true},
	{"The Adam Project", "Action, Adventure, Comedy", "Belum Ditonton", 1, 6.7, true},
	{"The Gray Man", "Action, Thriller", "Belum Ditonton", 1, 6.5, true},
	{"Day Shift", "Action, Comedy, Fantasy", "Belum Ditonton", 1, 6.1, true},
	{"Slumberland", "Adventure, Comedy, Family", "Belum Ditonton", 1, 6.7, true},
	{"Glass Onion: A Knives Out Mystery", "Comedy, Crime, Drama", "Belum Ditonton", 1, 7.1, true},
	{"The Pale Blue Eye", "Crime, Horror, Mystery", "Belum Ditonton", 1, 6.6, true},
	{"You People", "Comedy, Romance", "Belum Ditonton", 1, 5.5, true},
	{"Your Place or Mine", "Comedy, Romance", "Belum Ditonton", 1, 5.6, true},
	{"Luther: The Fallen Sun", "Crime, Drama, Mystery", "Belum Ditonton", 1, 6.4, true},
	{"Murder Mystery 2", "Action, Comedy, Crime", "Belum Ditonton", 1, 5.7, true},
	{"Extraction 2", "Action, Thriller", "Belum Ditonton", 1, 7.0, true},
	{"Beasts of No Nation", "Drama, War", "Belum Ditonton", 1, 7.7, true},
	{"The Ridiculous 6", "Comedy, Western", "Belum Ditonton", 1, 4.8, true},
	{"Crouching Tiger, Hidden Dragon: Sword of Destiny", "Action, Adventure", "Belum Ditonton", 1, 6.1, true},
	{"Pee-wee's Big Holiday", "Adventure, Comedy", "Belum Ditonton", 1, 6.1, true},
	{"Special Correspondents", "Comedy", "Belum Ditonton", 1, 5.8, true},
	{"The Do-Over", "Action, Comedy", "Belum Ditonton", 1, 5.7, true},
	{"The Fundamentals of Caring", "Comedy, Drama", "Belum Ditonton", 1, 7.3, true},
	{"Tallulah", "Drama", "Belum Ditonton", 1, 6.7, true},
	{"XOXO", "Drama, Music", "Belum Ditonton", 1, 5.3, true},
	{"ARQ", "Sci-Fi, Thriller", "Belum Ditonton", 1, 6.4, true},
	{"The Siege of Jadotville", "Action, Drama", "Belum Ditonton", 1, 7.2, true},
	{"Mascots", "Comedy", "Belum Ditonton", 1, 5.7, true},
	{"I Am the Pretty Thing That Lives in the House", "Horror, Mystery", "Belum Ditonton", 1, 4.6, true},
	{"7 años", "Drama", "Belum Ditonton", 1, 6.8, true},
	{"True Memoirs of an International Assassin", "Action, Comedy", "Belum Ditonton", 1, 5.9, true},
	{"Spectral", "Action, Sci-Fi", "Belum Ditonton", 1, 6.3, true},
	{"Barry", "Biography, Drama", "Belum Ditonton", 1, 5.9, true},
	{"Clinical", "Horror, Thriller", "Belum Ditonton", 1, 5.1, true},
	{"iBoy", "Action, Crime", "Belum Ditonton", 1, 6.0, true},
	{"Imperial Dreams", "Drama", "Belum Ditonton", 1, 6.7, true},
	{"Girlfriend's Day", "Comedy, Drama", "Belum Ditonton", 1, 5.2, true},
	{"I Don't Feel at Home in This World Anymore", "Comedy, Crime", "Belum Ditonton", 1, 6.9, true},
	{"Burning Sands", "Drama", "Belum Ditonton", 1, 6.1, true},
	{"Deidra & Laney Rob a Train", "Comedy, Crime", "Belum Ditonton", 1, 6.1, true},
	{"The Discovery", "Drama, Romance", "Belum Ditonton", 1, 6.3, true},
	{"Win It All", "Comedy, Drama", "Belum Ditonton", 1, 6.2, true},
	{"Sandy Wexler", "Comedy", "Belum Ditonton", 1, 5.1, true},
	{"Small Crimes", "Crime, Drama", "Belum Ditonton", 1, 5.8, true},
	{"Handsome: A Netflix Mystery Movie", "Comedy, Mystery", "Belum Ditonton", 1, 5.3, true},
	{"War Machine", "Comedy, Drama", "Belum Ditonton", 1, 6.0, true},
	{"Shimmer Lake", "Crime, Drama", "Belum Ditonton", 1, 6.2, true},
	{"Okja", "Action, Adventure", "Belum Ditonton", 1, 7.3, true},
	{"To the Bone", "Drama", "Belum Ditonton", 1, 6.8, true},
	{"The Incredible Jessica James", "Comedy, Romance", "Belum Ditonton", 1, 6.5, true},
	{"Naked", "Comedy, Romance", "Belum Ditonton", 1, 5.4, true},
	{"Death Note", "Crime, Drama", "Belum Ditonton", 1, 4.5, true},
	{"Little Evil", "Comedy, Horror", "Belum Ditonton", 1, 5.7, true},
	{"#REALITYHIGH", "Comedy, Romance", "Belum Ditonton", 1, 5.2, true},
	{"First They Killed My Father", "Biography, Drama", "Belum Ditonton", 1, 7.1, true},
	{"Gerald's Game", "Drama, Horror", "Belum Ditonton", 1, 6.5, true},
	{"The Babysitter", "Comedy, Horror", "Belum Ditonton", 1, 6.3, true},
	{"1922", "Crime, Drama", "Belum Ditonton", 1, 6.2, true},
	{"Wheelman", "Action, Crime", "Belum Ditonton", 1, 6.4, true},
	{"The Meyerowitz Stories (New and Selected)", "Comedy, Drama", "Belum Ditonton", 1, 6.9, true},
	{"Mudbound", "Drama, War", "Belum Ditonton", 1, 7.4, true},
	{"Bright", "Action, Crime", "Belum Ditonton", 1, 6.3, true},
	{"The Cloverfield Paradox", "Drama, Horror", "Belum Ditonton", 1, 5.5, true},
	{"Mute", "Mystery, Sci-Fi", "Belum Ditonton", 1, 5.4, true},
	{"All Quiet on the Western Front", "Drama, War", "Belum Ditonton", 1, 7.8, true},
	{"Athena", "Action, Drama", "Belum Ditonton", 1, 6.8, true},
	{"Atlantics", "Drama, Fantasy", "Belum Ditonton", 1, 6.7, true},
	{"The Night Comes for Us", "Action, Thriller", "Belum Ditonton", 1, 6.9, true},
	{"Apostle", "Horror, Mystery", "Belum Ditonton", 1, 6.3, true},
	{"The Sea Beast", "Animation, Adventure", "Belum Ditonton", 1, 7.1, true},
	{"Stealing Raden Saleh", "Action, Crime", "Belum Ditonton", 1, 7.3, true},
	{"Extreme Job", "Action, Comedy", "Belum Ditonton", 1, 7.1, true},
	{"The Fall of the House of Usher", "Drama, Horror", "Belum Ditonton", 1, 7.5, true},
	{"Beckham: Limited Series", "Documentary, Sport", "Belum Ditonton", 1, 7.8, true},
	{"Lupin", "Action, Crime", "Belum Ditonton", 1, 7.5, true},
	{"Tunnel", "Crime, Drama", "Belum Ditonton", 1, 7.3, true},
	{"How to Make Millions Before Grandma Dies", "Comedy, Drama", "Belum Ditonton", 1, 7.0, true},
	{"The Lost City", "Action, Adventure", "Belum Ditonton", 1, 6.1, true},
	{"The Man from U.N.C.L.E.", "Action, Adventure", "Belum Ditonton", 1, 7.3, true},
	{"Emily the Criminal", "Crime, Drama", "Belum Ditonton", 1, 6.7, true},
	{"Blackhat", "Action, Crime", "Belum Ditonton", 1, 5.4, true},
	{"Collateral", "Crime, Drama", "Belum Ditonton", 1, 7.5, true},
	{"The Birds", "Horror, Thriller", "Belum Ditonton", 1, 7.7, true},
	{"Carol", "Drama, Romance", "Belum Ditonton", 1, 7.2, true},
	{"Lift", "Action, Comedy", "Belum Ditonton", 1, 5.5, true},
	{"Badland Hunters", "Action, Drama", "Belum Ditonton", 1, 5.9, true},
	{"Red Notice 2", "Action, Comedy", "Belum Ditonton", 1, 6.3, true},
	{"Extraction 3", "Action, Thriller", "Belum Ditonton", 1, 7.5, true},
	{"Don't Look Up 2", "Comedy, Drama", "Belum Ditonton", 1, 7.2, true},
	{"The Killer", "Action, Crime", "Belum Ditonton", 1, 6.9, true},
	{"Pain Hustlers", "Drama, Thriller", "Belum Ditonton", 1, 6.5, true},
	{"Bright 2", "Action, Fantasy", "Belum Ditonton", 1, 6.7, true},
	{"Nowhere", "Thriller, Drama", "Belum Ditonton", 1, 6.8, true},
	{"The Gray Man 2", "Action, Thriller", "Belum Ditonton", 1, 6.8, true},
	{"Glass Onion 2", "Mystery, Crime", "Belum Ditonton", 1, 7.4, true},
	{"Spiderhead", "Sci-Fi, Thriller", "Belum Ditonton", 1, 5.8, true},
	{"The Adam Project 2", "Action, Sci-Fi", "Belum Ditonton", 1, 7.0, true},
	{"Day Shift 2", "Action, Comedy", "Belum Ditonton", 1, 6.0, true},
	{"Slumberland 2", "Fantasy, Adventure", "Belum Ditonton", 1, 6.5, true},
	{"Troll 2", "Adventure, Action", "Belum Ditonton", 1, 6.2, true},
	{"Sweet Girl 2", "Drama, Action", "Belum Ditonton", 1, 5.8, true},
	{"Your Place or Mine 2", "Comedy, Romance", "Belum Ditonton", 1, 5.6, true},
	{"The Pale Blue Eye 2", "Drama, Mystery", "Belum Ditonton", 1, 6.4, true},
	{"First They Killed My Father 2", "Drama, History", "Belum Ditonton", 1, 7.1, true},
	{"The Cloverfield Paradox 2", "Sci-Fi, Mystery", "Belum Ditonton", 1, 5.8, true},
	{"The Siege of Jadotville 2", "War, Drama", "Belum Ditonton", 1, 7.3, true},
	{"Shimmer Lake 2", "Mystery, Drama", "Belum Ditonton", 1, 6.3, true},
	{"Mudbound 2", "Drama, War", "Belum Ditonton", 1, 7.4, true},
	{"Spectral 2", "Sci-Fi, Action", "Belum Ditonton", 1, 6.4, true},
	{"1922 2", "Horror, Thriller", "Belum Ditonton", 1, 6.2, true},
}

var users = map[string]string{}
var profiles = map[string]string{}
var loggedInUser string
var activeProfile string

func main() {
	fmt.Println("Selamat Datang di Aplikasi Netflix Sederhana!")
	loginOrRegister()
	createOrSelectProfile()
	mainMenu()
}

func loginOrRegister() {
	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih (1/2/3): ")
		var choice string
		fmt.Scan(&choice)

		if choice == "1" {
			var username, password string
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)

			if pass, exists := users[username]; exists && pass == password {
				loggedInUser = username
				fmt.Println("Login berhasil!")
				break
			} else {
				fmt.Println("Username atau password salah.")
			}
		} else if choice == "2" {
			var username, password string
			fmt.Print("Buat Username: ")
			fmt.Scan(&username)
			fmt.Print("Buat Password: ")
			fmt.Scan(&password)

			if _, exists := users[username]; exists {
				fmt.Println("Username sudah digunakan.")
			} else {
				users[username] = password
				fmt.Println("Registrasi berhasil!")
			}
		} else if choice == "3" {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			os.Exit(0)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func createOrSelectProfile() {
	for {
		if _, exists := profiles[loggedInUser]; !exists {
			fmt.Println("Belum ada profil, buat profil baru.")
			fmt.Print("Nama Profil: ")
			var profileName string
			fmt.Scan(&profileName)
			profiles[loggedInUser] = profileName
			activeProfile = profileName
			fmt.Println("Profil berhasil dibuat!")
			clearscreen() // Membersihkan terminal setelah profil dibuat
			break
		} else {
			fmt.Println("Profil tersedia untuk pengguna ini: ")
			fmt.Println("1.", profiles[loggedInUser])
			fmt.Print("Pilih profil (1): ")
			var choice string
			fmt.Scan(&choice)
			if choice == "1" {
				activeProfile = profiles[loggedInUser]
				fmt.Println("Profil berhasil dipilih!")
				clearscreen() // Membersihkan terminal setelah profil dipilih
				break
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		}
	}
}

func mainMenu() {
	for {
		fmt.Println("\nMenu Utama:")
		fmt.Println("1. Today's Pick")
		fmt.Println("2. Top Netflix TV Show")
		fmt.Println("3. Top Netflix Originals Movie")
		fmt.Println("4. Browse by Genre")
		fmt.Println("5. Search by Title")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih (1/2/3/4/5/0): ")
		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			todaysPick()
		case "2":
			topNetflixTVShowIterative()
		case "3":
			topNetflixOriginalsMovieRecursive(movieList, len(movieList))

		case "4":
			browseByGenre()
		case "5":
			searchByTitle()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func todaysPick() {
	rand.Seed(time.Now().UnixNano())
	allShows := append(tvShowList, movieList...)
	rand.Shuffle(len(allShows), func(i, j int) { allShows[i], allShows[j] = allShows[j], allShows[i] })

	// Pastikan slice digunakan
	fmt.Println("\nToday's Pick :")
	for _, show := range allShows[:5] {
		fmt.Printf("Judul: %s, Genre: %s, Rating: %.1f\n", show.title, show.genre, show.rating)
	}
}

func topNetflixTVShowIterative() {
	fmt.Println("\nTop Netflix TV Shows:")

	sortedShows := make([]show, len(tvShowList))
	copy(sortedShows, tvShowList)

	// Mulai pengukuran waktu
	fmt.Println("Memulai pengukuran waktu...")
	start := time.Now()

	// Sorting iteratif
	for i := 0; i < len(sortedShows)-1; i++ {
		for j := 0; j < len(sortedShows)-i-1; j++ {
			if sortedShows[j].rating < sortedShows[j+1].rating {
				sortedShows[j], sortedShows[j+1] = sortedShows[j+1], sortedShows[j]
			}
		}
	}

	// Akhiri pengukuran waktu
	elapsed := time.Since(start)
	fmt.Println("Pengukuran waktu selesai.")
	fmt.Printf("Waktu untuk sorting: %s\n", elapsed)

	// Debug output sederhana
	for i := 0; i < len(sortedShows) && i < 10; i++ {
		fmt.Printf("%d. Judul: %s, Rating: %.1f\n", i+1, sortedShows[i].title, sortedShows[i].rating)
	}

	// Tampilkan hasil menggunakan paginatedPrint
	paginatedPrint(sortedShows)
}

func topNetflixOriginalsMovieRecursive(movies []show, n int) {

	fmt.Println("Mengukur waktu sorting...")
	start := time.Now()

	// Proses sorting rekursif
	if n == 1 {
		fmt.Println("\nTop Netflix Originals Movies:")
		// Tampilkan hasil menggunakan paginatedPrint
		paginatedPrint(movies)
		return
	}

	// Cari elemen terbesar di sisa array
	maxIndex := 0
	for i := 1; i < n; i++ {
		if movies[i].rating > movies[maxIndex].rating {
			maxIndex = i
		}
	}

	// Swap elemen terbesar dengan elemen terakhir
	movies[maxIndex], movies[n-1] = movies[n-1], movies[maxIndex]

	// Rekursi untuk elemen berikutnya
	elapsed := time.Since(start)
	fmt.Printf("Waktu untuk sorting: %s\n", elapsed)
	topNetflixOriginalsMovieRecursive(movies, n-1)
}

func browseByGenre() {
	fmt.Println("\nBrowse by Genre: Pilih hingga 5 genre (pisahkan dengan koma, contoh: Horror, Action, Drama)")
	var inputGenres string
	fmt.Scan(&inputGenres)

	selectedGenres := strings.Split(inputGenres, ",")
	if len(selectedGenres) > 5 {
		fmt.Println("Maksimal hanya bisa memilih 5 genre.")
		return
	}

	for i := range selectedGenres {
		selectedGenres[i] = strings.TrimSpace(strings.ToLower(selectedGenres[i]))
	}

	allShows := append(tvShowList, movieList...)
	var filteredShows []show

	for _, s := range allShows {
		showGenres := strings.Split(strings.ToLower(s.genre), ",")
		for i := range showGenres {
			showGenres[i] = strings.TrimSpace(showGenres[i])
		}

		for _, selectedGenre := range selectedGenres {
			for _, showGenre := range showGenres {
				if selectedGenre == showGenre {
					filteredShows = append(filteredShows, s)
					break
				}
			}
		}
	}

	// Sorting hasil berdasarkan rating (descending)
	for i := 0; i < len(filteredShows)-1; i++ {
		for j := 0; j < len(filteredShows)-i-1; j++ {
			if filteredShows[j].rating < filteredShows[j+1].rating {
				filteredShows[j], filteredShows[j+1] = filteredShows[j+1], filteredShows[j]
			}
		}
	}

	fmt.Println("\nHasil Pencarian berdasarkan Genre:")
	paginatedPrint(filteredShows)
}

func searchByTitle() {
	fmt.Println("\nSearch by Title: Masukkan judul film atau series")
	var keyword string
	fmt.Scan(&keyword)
	keyword = strings.ToLower(strings.TrimSpace(keyword))

	allShows := append(tvShowList, movieList...)
	var foundShows []show

	for _, s := range allShows {
		if strings.Contains(strings.ToLower(s.title), keyword) {
			foundShows = append(foundShows, s)
		}
	}

	if len(foundShows) == 0 {
		fmt.Println("Maaf, tidak ada film atau series yang cocok.")
	} else {
		paginatedPrint(foundShows)
	}
}

func paginatedPrint(shows []show) {
	// Opsi default untuk pageSize
	pageSize := 5

	// Pilihan untuk mengubah pageSize
	for {
		fmt.Println("\nPilih jumlah item per halaman:")
		fmt.Println("1. 5")
		fmt.Println("2. 10")
		fmt.Println("3. 15")
		fmt.Println("4. 20")
		fmt.Println("5. 25")
		fmt.Println("6. 50")
		fmt.Print("Pilih (1/2/3/4/5/6): ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			pageSize = 5
		case 2:
			pageSize = 10
		case 3:
			pageSize = 15
		case 4:
			pageSize = 20
		case 5:
			pageSize = 25
		case 6:
			pageSize = 50
		default:
			fmt.Println("Pilihan tidak valid, gunakan default 5.")
			pageSize = 5
		}

		// Konfirmasi jumlah item per halaman
		fmt.Printf("Jumlah item per halaman diatur ke %d.\n", pageSize)
		break
	}

	totalPages := (len(shows) + pageSize - 1) / pageSize
	currentPage := 0

	for {
		// Hitung rentang data untuk halaman ini
		start := currentPage * pageSize
		end := start + pageSize
		if end > len(shows) {
			end = len(shows)
		}

		for {
			// Mulai pengukuran waktu untuk setiap halaman
			startTime := time.Now()

			// Hitung rentang data untuk halaman ini
			start := currentPage * pageSize
			end := start + pageSize
			if end > len(shows) {
				end = len(shows)
			}

			// Header halaman
			fmt.Printf("\n===== Hasil (Halaman %d/%d) =====\n", currentPage+1, totalPages)

			// Tampilkan item di halaman ini
			for i := start; i < end; i++ {
				fmt.Printf("%d. Judul: %s\n   Genre: %s\n   Rating: %.1f\n", i+1, shows[i].title, shows[i].genre, shows[i].rating)
			}

			// Hitung waktu eksekusi halaman
			elapsedTime := time.Since(startTime)
			fmt.Printf("\nWaktu untuk menampilkan halaman ini: %s\n", elapsedTime)

			// Footer halaman
			fmt.Println("===============================")

			// Opsi navigasi
			fmt.Print("\nNavigasi (n: Next, p: Previous, c: Change page size, m: Main Menu): ")
			var navChoice string
			fmt.Scan(&navChoice)

			switch navChoice {
			case "n":
				if currentPage < totalPages-1 {
					currentPage++
				} else {
					fmt.Println("Anda sudah di halaman terakhir.")
				}
			case "p":
				if currentPage > 0 {
					currentPage--
				} else {
					fmt.Println("Anda sudah di halaman pertama.")
				}
			case "c":
				fmt.Println("\nKembali ke pengaturan jumlah item per halaman.")
				paginatedPrint(shows) // Rekursi untuk memulai ulang dengan pageSize baru
				return
			case "m":
				return
			default:
				fmt.Println("Pilihan tidak valid. Coba lagi.")
			}
		}
	}

}

func clearscreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Gunakan "cls" untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
