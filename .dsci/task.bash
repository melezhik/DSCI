set -e
cd ~/projects/dsci
git pull
echo "go build"
go build main.go
echo "restart service"
sudo service dsci restart
