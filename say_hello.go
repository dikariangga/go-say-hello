package go_say_hello

/**
Jika ada perubahan major, contoh menyebabkan error ketika upgrade
version. Sebaiknya rename nama module. Tambahkan "\v2" dan seterusnya
agar tidak menimbulkan error ketika pemaikaian.
*/

func SayHello(name string) string {
	return "Hello " + name
}
