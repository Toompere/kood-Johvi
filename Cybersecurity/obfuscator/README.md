# [obfuscator](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/obfuscator)

## Overview

**Obfuscator** is a project designed to explore polymorphic encryption and demonstrate how to alter a program's signature with each execution.

### Polymorphic Encryption

Polymorphic encryption is an advanced cryptographic technique that dynamically adapts its encryption patterns to create varied, constantly changing ciphertext, making detection and reverse-engineering more challenging. Unlike traditional encryption, which produces consistent outputs for identical plaintext and keys, polymorphic encryption modifies its algorithm with each encryption attempt. This adaptive approach—often achieved by varying encryption logic or using unique keys per instance—helps evade signature-based detection methods. While commonly associated with malware obfuscation, it also holds potential in cybersecurity contexts for countering static analysis methods.

### Signature Variation per Execution

Modifying the program's signature with each execution involves changing specific elements so that its recognizable "signature" varies each time it runs. This can be achieved by injecting random code, reordering instructions, or modifying non-functional sections of the program, such as unused padding bytes. These techniques, typical in polymorphic or metamorphic malware, allow the program to evade detection by making each version appear unique, thereby defeating static signature-based scanning tools and complicating code analysis.

## Key Features

- **Dynamic Signature Alteration**: On each execution, the program reads the content of its executable, appends an empty byte, and creates a modified file to replace the original, changing its signature.
  
- **Reverse Shell Connection**: After altering its signature, the program establishes a reverse shell connection to the attacker.

## Prerequisites

- **Linux**: Tested on Kali Linux as the target system and Ubuntu as the attacker system.
- **netcat**: Needs to be installed on the attacker’s machine.

### Installation Instructions

1. **Set Up a Linux Virtual Machine**:
   - Install any Linux OS version on a virtual machine to act as the victim system.

2. **Install Netcat**:
   - Install netcat on the attacker's system.

3. **Configure and Compile the Program**:
   - In `obfuscator.go`, define the attacker’s IP address and desired port on lines 47 and 48.
   - Open a terminal, navigate to the project directory, and compile the program with:
     ```bash
     go build obfuscator.go
     ```
   - Transfer the compiled program to the victim system.

4. **Start netcat on the Attacker System**:
   - Begin listening for the reverse shell connection by running:
   ```bash
   nc -lp <port>
   ```
5. **Execute the Program on the Victim System**:
   - Run the program on the victim machine. You can now enter shell commands on the attacker’s  system, where you initiated netcat.

## Video

Video showing the program in action can be found [here](https://www.youtube.com/watch?v=fdfFKgVmNWw).

## Audit
 [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/obfuscator/audit)


## Autor
[MargusT](https://01.kood.tech/git/MargusT)

## Warning
⚠️ These methods and tools are for educational purposes only, so that you have a better understanding of how to protect against similar vulnerabilities. You must ensure that you do not attempt any exploit-type activity without the explicit permission of the owner of the machine, system or application. Failure to obtain permission risks breaking the law.