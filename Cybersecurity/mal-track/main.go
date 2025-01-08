package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func main() {
	fmt.Println("--------------------------------Mal-Track Remover--------------------------------")
	fileNames := []string{"maltrack.exe", "mal-track.exe"}
	keyNames := []string{"mal-track", "maltrack"}

	killProcess(fileNames)
	removeFromStartup(keyNames)
	searchAndRemove(fileNames)
	
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Println("Press 'Enter' to exit...")
	fmt.Scanln() 
}

func killProcess(processName []string) {
	killed := false

	for _, p := range processName {
		cmd := exec.Command("taskkill", "/IM", p, "/F")
		err := cmd.Run()
		if err == nil {
			killed = true
			fmt.Println(p + " has been killed.")
		}
	}

	if !killed {
		fmt.Println("No running malware process found")
	}
}

func removeFromStartup(keyNames []string) {
	removed := false
	rootKeys := []registry.Key{registry.CURRENT_USER, registry.LOCAL_MACHINE}
	paths := []string{`Software\Microsoft\Windows\CurrentVersion\Run`, `Software\Microsoft\Windows\CurrentVersion\RunOnce`}

	for _, keyName := range keyNames {
		for _, k := range rootKeys {
			for _, p := range paths {
				regKey, err := registry.OpenKey(k, p, registry.ALL_ACCESS)
				if err == nil {
					err = regKey.DeleteValue(keyName)
					if err == nil {
						removed = true
						fmt.Println(keyName + " removed from startup registry at " + p)
					}
				}
			}
		}
	}

	if !removed {
		fmt.Println("Mal-track not found in startup registry.")
	}
}

func searchAndRemove(fileNames []string) {
	found := false
	ipList := []string{}

	for _, f := range fileNames {
		filepath.WalkDir("C:\\Users", func(path string, d os.DirEntry, err error) error {
			if err == nil {
				if strings.EqualFold(d.Name(), f) {
					found = true
					ipList = findIP(path, ipList)
					err = os.Remove(path)
					if err != nil {
						fmt.Println("Could not remove " + f + ": " + err.Error())
					} else {
						fmt.Printf("%s removed from %s\n", f, path)
					}
				}
			}
			return nil
		})
	}

	uniqueIPList := func(ipList []string) []string {
		exists := make(map[string]bool)
		result := []string{}
		for _, ip := range ipList {
			if !exists[ip] {
				result = append(result, ip)
				exists[ip] = true
			}
		}
		return result
	}(ipList)

	if len(uniqueIPList) > 0 {
		fmt.Println("IP addresses that might be associated with the attacker:")
		for _, ip := range uniqueIPList {
			fmt.Println("\t" + ip)
		}
	} else {
		fmt.Println("Attacker IP address not found.")
	}

	if !found {
		fmt.Println("No malware executable found.")
	}
}

func findIP(path string, ipList []string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ipList
	}

	ipRegex := `((?:[0-9]{1,3}\.){3}[0-9]{1,3}):[0-9]{1,5}`
	re := regexp.MustCompile(ipRegex)

	matches := re.FindAllSubmatch(data, -1)
	if matches != nil {
		for _, match := range matches {
			ipList = append(ipList, string(match[1]))
		}
	}
	return ipList
}
