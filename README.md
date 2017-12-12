# Parking service

Service ini digunakan untuk meminta data kapasitas, jumlah tersedia dan tarif dari sebuah tempat parkir dengan mengirim request GET dengan response dalam JSON.

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
$ cd $GOPATH
$ git clone https://github.com/daussho/service.git
```
- Instal Beego
```sh
$ go get github.com/astaxie/beego
```
- Install mysql for beego
```sh
$ go get -u github.com/go-sql-driver/mysql
```
- Run project
```sh
$ bee run service
```
