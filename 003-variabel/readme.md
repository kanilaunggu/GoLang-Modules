# Variabel Golang

Variabel adalah nama yang mewakili data yang disimpan dalam memori komputer. Inilah beberapa hal penting yang perlu diketahui tentang variabel Golang:

Golang adalah bahasa yang diketik secara statis, yang berarti bahwa ketika variabel dideklarasikan, mereka dapat secara eksplisit atau implisit menetapkan jenis bahkan sebelum program Anda berjalan.

Golang mengharuskan setiap variabel yang Anda nyatakan di dalam main () digunakan di suatu tempat di program Anda.
Anda bisa menetapkan nilai baru ke variabel yang sudah ada, tetapi harus berupa nilai dari tipe yang sama.

Variabel yang dideklarasikan dalam kurung kurawal {}, kurung kurawal pembukaan {memperkenalkan lingkup baru yang berakhir dengan kurung kurawal}.

### Beberapa ketentuan Variabel
Berikut adalah aturan berikut untuk memberi nama variabel Golang:
1. Nama harus dimulai dengan huruf, dan dapat memiliki sejumlah huruf dan angka tambahan.
2. Nama variabel tidak dapat dimulai dengan angka.
3. Nama variabel tidak boleh berisi spasi.
4. Jika nama variabel dimulai dengan huruf kecil, itu hanya dapat diakses dalam paket saat ini yang dianggap sebagai variabel yang tidak diekspor.
5. Jika nama variabel dimulai dengan huruf kapital, dapat diakses dari paket di luar paket saat ini yang dianggap sebagai variabel yang diekspor.
6. Jika nama terdiri dari beberapa kata, setiap kata setelah yang pertama harus menggunakan huruf besar seperti ini: empName, EmpAddress, dll.
7. Nama variabel adalah case-sensitive (mobil, Mobil dan MOBIL adalah tiga variabel yang berbeda).
