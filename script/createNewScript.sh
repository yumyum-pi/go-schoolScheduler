#!/bin/bash
# This scripts create new script with chmod +x

NAME="" #name of the folder
folder=$(pwd)

get_name() {
    #get the folder name
    echo "> Enter the script name"
    read NAME

    #input formatting
    #make the input lower case
    NAME=$( tr '[:upper:]' '[:lower:]' <<< $NAME) 
    #remove the white space
    NAME="${NAME// /_}" 
}

get_name

FILENAME="${NAME}.sh"

#make the file directory
FILE_DIR="$folder/script/$FILENAME"


FILE_DATA="#!/bin/bash \necho \"new Script\" "
echo -e $FILE_DATA > $FILE_DIR #write the data to the file

#make the script executable
chmod +x $FILE_DIR