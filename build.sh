
pwd
echo "Building Gateway..."
cd gateway
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway main.go
echo "Building Connect..."
cd ../connect
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o connect main.go
echo "Building Core..."
cd ../core
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o core main.go
echo "Building Queue..."
cd ../queue
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o queue  main.go
