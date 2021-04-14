# groupie-tracker-visualizations

The website uses only standard Golang packages and should be able to run on any computer with the recent version of Golang

How to run:
```
go run .
```
or 
```
go build -o <name-of-the-binary>
<name-of-the-binary>
```
Using Docker:
```
docker build -t groupie-tracker .
docker run -p 9000:9000 groupie-tracker
```
or by launching the script
```
chmod 777 launch.sh
./launch.sh
```