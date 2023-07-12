# go-cms
### go ile yapmış olduğum basit bir content manegment system
### Yapılacaklar listesi

- [X] Blog yönetim sistemi (ekleme, silme ve düzenleme)
- [X] Galeri Sistemi (ekleme, silme ve düzenleme)
- [X] Kullanıcı yönetimi
- [X] Blog Beğeni sistemi
- [ ] Toplu seçme
- [X] E-ticaret / ürün yönetimi ( ödeme sistemleri eklenecek ve toplu ürün yükleme )
- [ ] Forum yönetii
- [ ] Doğrulama sistemleri
- [ ] ReCaptcha
- [ ] Toplu silme
- [ ] CSV veya JSON ile toplu Blog, ürün vs. Ekleme
- [ ] Kolay kurulum
- [X] Kolay düzenleme ve estetik tasarım 
- [ ] WYSIWYG
- [X] SEO
- [ ] Chart

### Yapılan düzeltmeler (v0.3 BETA)
- Ufak Method düzeltmeleri <br>
``Delete fonksiyonları Delete Methodlarına çevrildi``
- Status düzeltmeleri <br>
``Artık redirect yerine Forbidden, Not found , vs. (400,401,404,500,...) statuleri gönderiyor``
- Routerde ki düzeltmeler <br>
``Main.go dosyası artık daha clean``
- Veritabanında ki düzeltmeler <br>
``Database klasöründe ki gereksiz connectler kaldırıldı, kod daha okunur hale getirildi``
- Kod daha clean hale getirildi <br>
``Routerlar Handlerlar ile ayrıldı``
- Buglar fixlendi
``
Fiberin Statik cache'ini yedim.
sildiğiniz Ürünlerde fotoğraf yoksa sunucu kapanıyordu, onuda düzelttim
``


### Siteden görseller
![Ekran görüntüsü 2023-07-05 032223](https://github.com/Hasan-Kilici/go-cms/assets/105741983/fb83c0b9-4e92-4a6b-9909-101600709bea)
![Ekran görüntüsü 2023-07-05 032434](https://github.com/Hasan-Kilici/go-cms/assets/105741983/cdf34277-171a-4fa2-b48c-ef21fb504c28)
![Ekran görüntüsü 2023-07-05 032522](https://github.com/Hasan-Kilici/go-cms/assets/105741983/cd2121fd-92be-4b18-811a-0d206cb7c0a9)
![Ekran görüntüsü 2023-07-05 032609](https://github.com/Hasan-Kilici/go-cms/assets/105741983/f7038642-3ae2-4770-9a98-68224a4da751)
![Ekran görüntüsü 2023-07-05 032747](https://github.com/Hasan-Kilici/go-cms/assets/105741983/d1347b50-76e0-43de-ba8d-c041f2fdb1bf)

