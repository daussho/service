# Parking service

Web service yang akan saya buat adalah web service untuk melihat kapasitas tempat parkir di ITB (Sipil, SR dan Saraga) secara real-time. Alat karcis akan menambah jumlah kendaraan saat ini saat mencetak karcis baru dan jumlah kendaraan akan otomatis berkurang saat ada kendaraan keluar. Dengan web service ini, diharapkan orang-orang yang ingin parkir di ITB dapat lebih mudah mencari tempat parkir karena tidak perlu berkeliling untuk mencari tempat parkir yang masih kosong.

# Functional Requirement

  - Mengirimkan response dalam JSON berisi nama tempat parkir, kapasitas, sisa tersedia, dan tarif
  - Update data tiket dilakukan via dashboard admin tanpa password

# Non functional Requirement
  - Using Golang 1.9
  - Database using mysql
  - OS ubuntu 16.04
  - Minimum RAM 250MB
  - Storage 10GB

### Installation

- Install Golang 1.9
- Clone project to $GOPATH
```sh
$ cd $GOPATH/src
$ git clone https://github.com/daussho/service.git
```
- Import database.sql to MySQL
- Instal Beego
```sh
$ go get github.com/astaxie/beego
$ go get -v github.com/beego/bee
```
- Install mysql driver for beego
```sh
$ go get -u github.com/go-sql-driver/mysql
```
- Run project
```sh
$ bee run service
```
# ACCESS
- http://ip:port/ (access home)
- http://ip:port/data/{id} (GET JSON data, with id=1: Parkiran Sipil, id=2: Parkiran SR, id=3: Parkiran Saraga)
- http://ip:port/input (Update tiket data)