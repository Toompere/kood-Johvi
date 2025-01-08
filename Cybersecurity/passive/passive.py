import argparse
from bs4 import BeautifulSoup
from collections import Counter
import os
import re
import requests
import subprocess

def search_user(username):
    sites = ['GitHub', 'LinkedIn', 'Reddit', 'Medium', 'YouTube']
    command = ['sherlock', username]
    for site in sites:
        command.extend(['--site', site])

    output = subprocess.run(command, capture_output=True, text=True)
    output_lines = output.stdout.splitlines()
    results = {site: 'no' for site in sites} 

    for line in output_lines:
        for site in sites:
            if site.lower() in line.lower():
                results[site] = 'yes'

    result = (
        f'GitHub: {results["GitHub"]}\n'
        f'LinkedIn: {results["LinkedIn"]}\n'
        f'Reddit: {results["Reddit"]}\n'
        f'Medium: {results["Medium"]}\n'
        f'YouTube: {results["YouTube"]}'
    )
    if os.path.exists(f"{username}.txt"):
        os.remove(f"{username}.txt")
    return result

def search_ip(ip):
    url = f'https://www.ipapi.co/{ip}/json'

    response = requests.get(url) 

    if response.status_code == 200:
        data = response.json()

        err = data.get('error')
        if err:
            print('Error: ' + data.get('reason'))
            exit(1)
        
        isp = data.get('org')
        lat = data.get('latitude')
        lon = data.get('longitude')

        return f'ISP: {isp}\nCity Lat/Lon:	({lat}) / ({lon})'

    return 'Not found'

def search_full_name(full_name):
    name = full_name.split(' ')
    if len(name) != 2:
        print('You need to enter first and last name')
        exit(1)

    first_name= 'First name: ' + name[0] + '\n'
    last_name= 'Last name: ' + name[1] + '\n'
    address = 'Address: ' + search_data(name[0], name[1], 'address') + '\n'
    number = 'Number: ' + search_data(name[0], name[1], 'phone')
    result = first_name + last_name + address + number
    
    return result

def search_data(first_name, last_name, attr):
    url = f'https://www.google.com/search?q={first_name}+{last_name}+{attr}'
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"
    }
    response = requests.get(url, headers=headers) 
    if response.status_code == 200:
        soup = BeautifulSoup(response.content, 'html.parser')
        text_content = soup.get_text()
        
        if attr == 'address':
            pattern = re.compile(
            r'\w[^,:]+,[^,:]+,\W\d{5}',
            re.IGNORECASE
            )

        if attr == 'phone':
            pattern = re.compile(
            r'\+\d{3}\s?\d{2,4}\s?\d{4,6}'
            )

        re_data = pattern.findall(text_content)
        data_counter = Counter(re_data)
        most_common = data_counter.most_common(1)  
        if most_common:
            return most_common[0][0]
        else:
            return 'Not found'
    else:
        return 'Not found'

def save_result(result):   
    filename = 'result.txt'
    i = 2
    while os.path.exists(filename):
        filename = f'result{i}.txt'
        i += 1
    
    with open(filename, 'w') as file:
        file.write(result)

    print(result)
    print(f'Saved in {filename}')

def main():
    parser = argparse.ArgumentParser(description="Welcome to passive v1.0.0")

    parser.add_argument(
        '-fn', 
        type=str, 
        help='Search with full-name'
    )
    parser.add_argument(
        '-ip', 
        type=str, 
        help='Search with ip address'
    )
    parser.add_argument(
        '-u', 
        type=str, 
        help='Search with username'
    )

    args = parser.parse_args()

    if args.fn:
        result = search_full_name(args.fn)
    if args.ip:
        result = search_ip(args.ip)
    if args.u:
        result = search_user(args.u)

    save_result(result)

if __name__ == "__main__":
    main()