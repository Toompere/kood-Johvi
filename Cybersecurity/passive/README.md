# [passive](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/passive)

## About

This is a simple OSINT script which you can use to search for a persons address and phone number, information about an IP address and social media information using username.

## OSINT

OSINT (Open Source Intelligence) is the practice of collecting and analyzing publicly available information to gain insights about a target, be it a person, organization, or topic. Unlike espionage or hacking, OSINT strictly involves legal and ethical data gathering from sources accessible to the public. These sources can include websites, social media profiles, forums, news outlets, public records, and more.

## Dependencies

- [Requests](https://pypi.org/project/requests/)
- [BeautifulSoup4](https://pypi.org/project/beautifulsoup4/)
- [ipapi.co](https://ipapi.co/)
- [Sherlock Project](https://sherlockproject.xyz/)

## Usage

- Setup:   
     ```bash
     git clone https://01.kood.tech/git/MargusT/passive.git
     cd passive
     pip install -r requirements.txt
- Search for a persons address and phone number(Currently works with Estonian addresses and numbers): 

    ```bash
    python3 passive.py -fn "full name" 
- Search for IP addresses ISP and location:
    ```bash
    python3 passive.py -ip "ip address" 
- Search for social media presence using username:
    ```bash
    python3 passive.py -u "username" 

## Audit
<img src="passive.gif" width="1400"> 

 [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/passive/audit)


## Autor
[MargusT](https://01.kood.tech/git/MargusT)

## Warning
⚠️ These methods and tools are for educational purposes only, so that you have a better understanding of how to protect against similar vulnerabilities. You must ensure that you do not attempt any exploit-type activity without the explicit permission of the owner of the machine, system or application. Failure to obtain permission risks breaking the law.