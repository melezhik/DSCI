set -e
cd ~/projects/dsci
git pull
echo "go build"
go build main.go
echo "restart dsci service"
sudo service dsci restart
