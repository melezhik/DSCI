set -e
cd ~/projects/dsci
git pull
go build main.go
sudo service dsci restart
