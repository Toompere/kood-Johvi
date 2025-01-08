# [evasion](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/evasion)

## Overview

**Evasion** is a project designed to deepen understanding of how antivirus software operates and explore methods of bypassing it for educational purposes.

### How Antiviruses Work

Antivirus software detects, prevents, and removes malware from computers using various methods:

- **Signature-Based Detection**: Compares files to a database of known malware signatures, quickly identifying familiar threats.
- **Heuristic Analysis**: Examines code behavior to detect previously unknown threats, though it may sometimes trigger false positives.
- **Behavioral Analysis**: Monitors system activity in real time, halting potential threats based on suspicious patterns before they cause harm.
- **Machine Learning**: Identifies malware patterns, detecting zero-day and sophisticated attacks by continuously learning from new data.
- **Additional Features**: Quarantine and sandboxing isolate suspicious files, minimizing risk by allowing safe examination.

### How to Bypass Antiviruses

Attackers use various techniques to bypass antivirus software:

- **Obfuscation**: Disguises malicious code to avoid detection.
- **Polymorphic and Metamorphic Malware**: Alters its appearance to evade signature-based detection.
- **Code Injection**: Embeds malware in trusted applications to blend with legitimate processes.
- **Fileless Malware**: Operates directly in memory, avoiding detection by traditional disk scans.
- **System Tools**: Uses legitimate tools to mimic normal activity.
- **Crypters**: Wrap malware in layers that hide its true nature.
- **Social Engineering**: Tricks users into allowing malware access, bypassing antivirus defenses altogether.

## Key Features

- **Program Encryption**: Evasion accepts a program (e.g., malware) as an argument and encrypts it for further processing.
  
- **Integer Incrementation**: During runtime, Evasion first increments an integer from 0 to 100,001 to simulate activity and delay detection.

- **Sleep Timer**: After incrementation, a 101-second sleep timer initiates to further obscure the process.

- **Decryption & Execution**: When the sleep timer completes, the program is decrypted and executed.

## Prerequisites

- **Windows Virtual Machine**: Set up a Windows VM for testing purposes.
- **Golang**: Install Golang on the system to compile and run Evasion.

### Installation Instructions

1. **Set Up a Windows Virtual Machine**:
   - Download and install any Windows OS version on a virtual machine.

2. **Install Golang**:
   - Download and install [Go](https://golang.org/dl/) on your Windows VM.

3. **Compile the Program**:
   - Open a terminal or command prompt, navigate to the project directory, and run the following command:
     ```bash
     GOOS=windows GOARCH=amd64 go build -o evasion.exe evasion.go
     ```

4. **Encrypt the Program**:
   - Run **evasion.exe** with the target program as an argument to encrypt it.

5. **Execute the Encrypted Program**:
    - Run the encrypted program. For demonstration it is printing out the incrementation and sleep timer in the command prompt.


## Video

Video showing the program in action can be found [here](https://www.youtube.com/watch?v=SLW0pKNRyyE).

## Audit
 [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/evasion/audit)


## Autor
[MargusT](https://01.kood.tech/git/MargusT)

## Warning
⚠️ These methods and tools are for educational purposes only, so that you have a better understanding of how to protect against similar vulnerabilities. You must ensure that you do not attempt any exploit-type activity without the explicit permission of the owner of the machine, system or application. Failure to obtain permission risks breaking the law.