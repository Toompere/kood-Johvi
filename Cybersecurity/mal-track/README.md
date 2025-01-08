# [mal-track](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/mal-track)

## Overview

The **Mal-Track Remover** is a tool written in Go designed to detect, kill, and remove a specific malware—`mal-track.exe`/`maltrack.exe`—from a Windows system. This project aims to help understand the basic behavior of a computer virus and simple methods to eradicate it. The remover performs four key actions:

1. **Terminate the Malware Process**: It detects and kills the running malware process.
2. **Remove Startup Entries**: It searches for and deletes any malware entries in Windows startup registry keys to prevent the virus from starting again after reboot.
3. **Extract Potential Attacker IP**: It scans the malware executable for embedded IP addresses that could belong to the attacker.
4. **Delete the Executable**: It removes the malware executable from the file system.

## Key Features

- **Process Termination**: If `mal-track.exe` or `maltrack.exe` is running, the remover will locate and terminate the process.
  
- **Startup Cleanup**: It searches four registry locations for keys associated with the malware and deletes them to prevent automatic execution on startup:
  - `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
  - `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce`
  - `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run`
  - `HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunOnce`

- **IP Address Extraction**: The remover scans the malware executable for IP addresses using regular expressions. These IP addresses might be used by the malware to communicate with its command-and-control server, although this method has limitations in detecting more advanced obfuscation techniques.

- **Executable Deletion**: Finally, the tool deletes the `maltrack.exe` file from the system to ensure it is no longer present.

## Prerequisites

- A Windows virtual machine for testing the tool.
- Golang installed on either the virtual machine or the host system for compiling the remover executable.

### Installation Instructions

1. **Set up a Windows Virtual Machine**:
   - Download and install any Windows OS version on a virtual machine (VM).
   - Ensure the VM is isolated for safety.
   
2. **Install Golang**:
   - Download and install [Go](https://golang.org/dl/) on your Windows VM or host system.

3. **Compile the Remover**:
   - Open a terminal or command prompt and navigate to the project directory.
   - Run the following command to compile the remover:
     ```bash
     GOOS=windows GOARCH=amd64 go build -o "Mal-Track Remover.exe" main.go
     ```

4. **Run the Malware**:
   - You may need to disable Windows Defender for the malware to run properly.

5. **Run the Remover**:
    - You may need to execute the remover with administator privileges.


## Video

Video showing the remover in action can be found [here](https://www.youtube.com/watch?v=5h8pG1BAu_w).

## Audit
 [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/mal-track/audit)


## Autor
[MargusT](https://01.kood.tech/git/MargusT)

## Warning
⚠️ These methods and tools are for educational purposes only, so that you have a better understanding of how to protect against similar vulnerabilities. You must ensure that you do not attempt any exploit-type activity without the explicit permission of the owner of the machine, system or application. Failure to obtain permission risks breaking the law.