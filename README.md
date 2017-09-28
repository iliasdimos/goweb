## Simple web server written in GO

Used for demostration purposes. 

Prints out the Host name in color background.


Environmental variables 
```
    PORT (Default: 8080)
    COLOR (Default: white)
```
Build with:
```
docker build -t simple-go-server .
```

Run 
```
docker run -d -p 8080:8080 -e COLOR=yellow simple-go-server
```