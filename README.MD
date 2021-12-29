## Note

- Golang mempunyai package bawaan sendiri untuk unit test bernama `testing`
- Nama file untuk di test harus berakhiran test contoh `hello_world_test.go`
- Nama function untuk test harus ditulis berawalan Test contoh `TestHelloWorld`
- Function test harus mempunyai parameter `t *testing.T` contoh `func TestHelloWorld(t *testing.T)`
- Function test tidak boleh mempunyai return value
- Untuk menggagalkan unit test bisa menggunakan t.Fail(), t.FailNow(), t.Error(), t.Fatal()
- Assertion di golang tidak disediakan secara default, tapi kita harus menggunakan **Testify** `go get github.com/stretchr/testify`
- assert jika terjadi error akan memanggil Fail(), require ketika terjadi error akan memanggil FailNow()
- Skip() untuk membatalkan test
- Before After test bisa menggunakan parameter testing.M untuk mengatur eksekusi unit test menggunakan main, contoh kita harus membuat function bernama TestMain dengan parameter testing.M
- Untuk melakukan subtest kita bisa menggunakan function Run() dari testing.T
- perintah untuk menjalankan subtest `go test -v -run=TestSubTest/Ravi`