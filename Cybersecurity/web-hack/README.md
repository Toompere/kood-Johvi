# [web-hack](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/web-hack)

## About

Web-hack is a project focused on learning how to secure web applications by exploiting vulnerabilities in [DVWA(Damn Vulnerable Web Application)](https://github.com/digininja/DVWA). This document outlines how I exploited three different vulnerabilities in DVWA and provides recommendations to safeguard your applications from these threats. Additionally, I’ve developed a web shell that allows users to add files, delete files, and execute commands.

## Setup

The environment was set up by installing Kali Linux in VirtualBox, followed by deploying DVWA within it. I used Apache2 as the web server and MySQL to manage the database.

## Vulnerabilities

### Brute force
A brute force attack is a method where an attacker tries all possible password combinations until the correct one is found.
- #### Low security
    In the lowest security setting there is no protection against brute force attacks. I used [wfuzz](https://github.com/xmendez/wfuzz) for attacking using top1000 and top10000 password lists from [SecLists](https://github.com/danielmiessler/SecLists). I started with the username `admin` and wfuzz quickly found out that the password was indeed `password`. 
    
    With successful login a "profile image" was displayed in DVWA and by inspecting its location I found that all other users images were inside the same folder. Using those usernames and brute force attacks i found out all of the passwords:
    ```
    1337:charley
    admin:password
    gordonb:abc123
    pablo:letmein
    smithy:password
    ```
- #### Medium security
    This version added a 2 second sleep timer was added in case of invalid credentials. This made the brute force attacks much slower but the same method could still be used.
- #### High security
    Here a [Anti-CSRF token](https://portswigger.net/web-security/csrf/bypassing-token-validation) was used. Every time the page reloaded a new visually hidden token was generated. For this level I used [BurpSuite](https://portswigger.net/burp). With the pitchfork attack I was able the use recursive grab to extract the token and successfully attack the application.
- #### Impossible security
    A 15 minute lock out period was added after 3 unsuccessful login attempts. this makes brute force attacks extremely slow. So unless you have set a very common and simple password it will be almost impossible to get the password using brute force.

### SQL Injection
A SQL injection attack consists of insertion or "injection" of a SQL query via the input data from the client to the application. A successful SQL injection exploit can read sensitive data from the database, modify database data (insert/update/delete), execute administration operations on the database (such as shutdown the DBMS), recover the content of a given file present on the DBMS file system (load_file) and in some cases issue commands to the operating system.y
- #### Low security
    A basic form is available where users are expected to enter a user ID, and it will display the corresponding first and last name. However, the input lacks proper security validation, making it vulnerable to SQL injection attacks. By entering a payload like `'UNION SELECT user, password FROM users#` into the input field, an attacker can retrieve a list of all usernames and their MD5-hashed passwords. Since MD5 is not a secure hashing algorithm, these passwords can easily be cracked.
- #### Medium security
    In this case, instead of a text input, there is a dropdown list of user IDs from which you can select and view the corresponding first and last names. However, this can still be bypassed. By using the browser's inspector tool, you can modify the selected ID in the dropdown to something like `1 UNION SELECT user, password FROM users`, which will again expose the list of usernames and their MD5-hashed passwords.
- #### High security
    This is very similar to the low level, however this time the attacker is inputting the value in a different manner. The input values are being transferred to the vulnerable query via session variables using another page, rather than a direct GET request. However, with the same SQL injection payload as before `'UNION SELECT user, password FROM users#`, the attacker can still retrieve the list of usernames and MD5 password hashes.
- #### Impossible security
    The queries are now parameterized queries (rather than being dynamic). This means the query has been defined by the developer, and has distinguish which sections are code, and the rest is data.
### Weak Session ID
A Session ID is a unique code that tracks your activity on a website, like keeping you logged in or remembering items in your shopping cart. Hackers can steal this code through methods like session hijacking, allowing them to take over your session and act as if they were you. This can give them access to your personal information or let them perform actions like making purchases or changing account settings without your permission.
- #### Low security
    The Session ID is generated by simply incrementing a number starting from 0. This approach offers no security, as attackers can easily predict and identify valid Session IDs, making users vulnerable.
- #### Medium security
    In this level the ID is generated using the current timestamp. This method is also highly vulnerable to attacks because attackers can predict the approximate time the session was created, making it easier to guess valid Session IDs.
- #### High security
    Here the ID is generated similarly to the low security method by incrementing a number, but this time it is encoded using MD5. While this provides slightly more security than the first approach, it is still vulnerable, as MD5 is a weak hashing algorithm and attackers can easily reverse or brute-force the encoded values.
- #### Impossible security
    At the impossible level, the ID is generated by encoding a combination of a random number, a timestamp, and the word 'impossible' using SHA-1. While this method provides enhanced security compared to earlier approaches, SHA-1 is considered outdated. Therefore, it is advisable to use SHA-3 or UUIDs for improved security.

## Web Shell
I created a simple web shell in php which can be used to execute shell commands, upload and delete files. If the shell file is uploaded to DVWA by file upload module. it can be accessed from ```dvwa/hackable/uploads/shell.php```. There is a html form where you can enter commands and add/remove files.

## Audit

 [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/web-hack/audit)


## Author
[MargusT](https://01.kood.tech/git/MargusT)

## Warning
⚠️ These methods and tools are for educational purposes only, so that you have a better understanding of how to protect against similar vulnerabilities. You must ensure that you do not attempt any exploit-type activity without the explicit permission of the owner of the machine, system or application. Failure to obtain permission risks breaking the law.