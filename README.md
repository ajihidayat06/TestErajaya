# TestErajaya
berisi endpoint product yang bertujuan untuk test seleksi employee PT Erajaya

Pada mini project ini berisi endpoint yang digunakan untuk menambah dan menampilkan list data product.
desain architecture yang saya gunakan pada mini project ini yaitu clean architecture.
alasan penggunaan architecture ini yaitu: 
-> architectur nya memisahkan antara bisnis layer dan teknologi layer. ketika terjadi perubahan diantara layer tersebut, maka kemungkinan mempengaruhi layer lainnya sangat kecil (lebih mudah mantainance)
-> mudah untuk di test, karena penggunaan interface pada kodenya 
-> sudah banyak digunakan, maka lebih banyak contoh best practice yang dapat ditiru dan dimodifikasi
-> ingin belajar dan mencoba menggunakan clean architecture (alasan pribadi, karena belum pernah menggunakannya)  

pada mini project ini yang saya gunakan:
a. gorm untuk konek DB dan pengoperasian query
b. golang fiber untuk framework

==========================================================================================
1. Add data product
   unruk add data product sudah saya sesuaikan field yang dibutuhkan sesuai dengan instruksi yang ada. berikut untuk curl postman add product 

curl --location 'http://localhost:3400/api/product' \
--header 'Authorization: admin:1707035205313128100' \
--header 'Content-Type: application/json' \
--data '{
    "name": "z flip samsung (new)",
    "price": 9000000,
    "description": "samsung z flip new",
    "quantity": 20
}'

===========================================================================================
2. List data product 
   List data product saya buat sesuai dengan instruksi. saya tambahkan juga untuk query page limit nya dengan nama paramnya "page" , "limit" 
   cara penggunaan param order sebagai berikut : 
     - misal order by name (Z-A), maka penulisan pada paramnya "order=name desc'
     - begitu juga untuk field-field lainnya, "order=price desc"
    berikut curl untuk list data product

    curl --location 'http://localhost:3400/api/product?page=1&limit=2&order=name%20desc' \
--header 'Authorization: admin:1707035205313128100'

=============================================================================================
3. endpoint login
  pada bagian ini saya menggunakan redis untuk menyimpan data token dari user yang login.. user yang saya buat hanya admin (hardcode di kodingan).
  token yang didapat harus di set ke header authorization agar endpoint2 lainnya bisa diakses.
  username : admin,
  password: admin123

  curl --location 'http://localhost:3400/api/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "admin",
    "password": "admin123"
}'

================================================================================================
4. endpoint logout
  untuk menghapus token user dari redis

  curl --location 'http://localhost:3400/api/logout' \
--header 'Authorization: admin:1707036096154383000'

================================================================================================
5. docker disini hanya saya gunakan untuk menjalankan container postgre dan redis (karena saya masih bejalar docker jadi baru bisa pull dan run saja)
![image](https://github.com/ajihidayat06/TestErajaya/assets/67070990/25120138-1230-4d65-9401-f3d34e25b776)
![image](https://github.com/ajihidayat06/TestErajaya/assets/67070990/b1b73d14-b765-492a-a96b-c947cb67dc01)

