

For testing purposes, I have included a file called directories.input that contains all the commands required for this tech challenge


In order to execute commands from a file:

./app --file={name_of_command_file}

i.e. ./app --file=directories.input


In order to excute individual commands:

./app CREATE automobiles
./app CREATE automobiles/truck


// Building the app

running the following command while in a golang environment will generate a executable file for MacOS

env GOOS=darwin GOARCH=amd64 go build github.com/kvmac/endpoint-app/app