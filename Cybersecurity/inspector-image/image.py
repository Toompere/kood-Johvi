import argparse
from PIL import Image
import re

def location(filename):
    img = Image.open(filename)
    img_exif = img.getexif()

    if not img_exif:
        print('No location info')
    else:
        try:
            # 34853 is key for gpsinfo dictionary where key 2 is latitude and key 4 is longitude
            print(f'Lat/Lon: ({img_exif.get_ifd(34853)[2][0]}) / ({img_exif.get_ifd(34853)[4][0]})')
        except KeyError:
            print('No location info')

def pgp(filename):
    with open(filename, "rb") as image:
        data = image.read()
        offset = data.index(bytes.fromhex('ffd9'))
        image.seek(offset + len('ffd9') // 2)
        output = image.read().decode('utf-8')
        match = re.search(r'-+B.+END.+-', output, re.DOTALL)
        if match:
            print(match.group(0))
        else:
            print('No PGP Key found')

def main():
    parser = argparse.ArgumentParser(description="Inspector Image v1.0.0")

    parser.add_argument('-map', type=str, help='Show the locatin where the image was taken')
    parser.add_argument('-steg', type=str, help='Show hidden pgp key')

    args = parser.parse_args()

    if args.map:
        location(args.map)
    if args.steg:
        pgp(args.steg)

    
if __name__ == "__main__":
    main()