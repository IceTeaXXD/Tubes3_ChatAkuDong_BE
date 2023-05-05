# Tubes3_ChatAkuDong_BE
> Tugas Besar 3 IF2211 - Strategi Algoritma

> Live demo [_here_](https://www.youtube.com/watch?v=OcAE0r-t8uI&feature=youtu.be). <!-- If you have the project hosted somewhere, include the link here. -->

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Screenshots](#screenshots)
* [Setup](#setup)
* [Acknowledgements](#acknowledgements)


## General Information
Dalam tugas besar ini, kami diminta untuk membangun sebuah aplikasi ChatGPT sederhana dengan mengaplikasikan pendekatan QA yang paling sederhana tersebut. Pencarian pertanyaan yang paling mirip dengan pertanyaan yang diberikan pengguna dilakukan dengan algoritma pencocokan string Knuth-Morris-Pratt (KMP) dan Boyer-Moore (BM). Regex digunakan untuk menentukan format dari pertanyaan (akan dijelaskan lebih lanjut pada bagian fitur aplikasi). Jika tidak ada satupun pertanyaan pada database yang exact match dengan pertanyaan pengguna melalui algoritma KMP ataupun BM, maka gunakan pertanyaan termirip dengan kesamaan setidaknya 90% Apabila tidak ada pertanyaan yang kemiripannya di atas 90%, maka chatbot akan memberikan maksimum 3 pilihan pertanyaan yang paling mirip untuk dipilih oleh pengguna.


## Technologies Used
- GoLang
- Gin
- Gorm
- ElephantSQL (Hosting for PostgreSQL)
- CompileDaemon


## Features
List the ready features here:
- Simple Arithmetic Calculator
- Day and Date Calculator
- Simple Chatbot (with KMP, BM, and Regex)

## Setup
1. To run locally, you need to install GoLang first. You can download it [here](https://golang.org/dl/).
2. Clone this repository.
```
git clone https://github.com/IceTeaXXD/Tubes3_ChatAkuDong_BE.git
```
3. Run the program
```
cd src
go run main.go
```
Alternatively, to run the program with hot reload, you can use CompileDaemon. Check [_here_](https://github.com/githubnemo/CompileDaemon)
```
cd src
CompileDaemon -command="./cad"
```

## Screenshots
![](doc/login.png)
<br>
![](doc/chat.png)

## Acknowledgements
- Tuhan Yang Maha Esa
- Dosen Pengampu Mata Kuliah IF2211 Strategi Algoritma
- Asisten Pengampu Mata Kuliah IF2211 Strategi Algoritma
- Our loved ones who support us to finish this project 💖