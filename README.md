Instructions

Download Go

wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin >> ~/.profile
source ~/.profile
go version