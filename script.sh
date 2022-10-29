#!/bin/bash
cd "$(dirname "$0")"
cd temp

echo "downloading data..."
wget -q ftp://opendata:kagar1n@ftp.euipo.europa.eu:/Trademark/Full/2020/EUTMS*.zip
echo "data downloaded"

echo "unzipping archives..."
find . -depth -name '*.zip' -exec unzip -qn {} \; -delete
echo "files extracted"

touch file_names
rm file_names
touch file_names

touch output.csv
rm output.csv
touch output.csv

echo "creating list of files..."
for i in $(find . -type f -name "*.xml")
do
    echo $(realpath $i) >> file_names
done
echo "list created"

cd ..

echo "creating output.csv file..."
php xml2csv.php
echo "output.csv file created"

cd temp
echo "removing unnecesary data..."
find . ! -name output.csv -delete
echo "removed"