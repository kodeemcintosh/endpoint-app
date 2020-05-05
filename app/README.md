
The executable for this app is meant for MacOS amd64 architecture

For testing purposes, I have included a file called directories.input that contains all the commands required for this tech challenge

In order to execute commands from a file:

./endpoint-app/app/app --file={name_of_command_file}

i.e. /endpoint-app/app/app --file=directories.input


In order to excute individual commands:

./endpoint-app/app/app CREATE automobiles
./endpoint-app/app/app CREATE automobiles/truck


// Building the app

running the following command while in a golang environment will generate a executable file for MacOS

env GOOS=darwin GOARCH=amd64 go build github.com/kvmac/endpoint-app/app