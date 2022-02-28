# Sending-data-from-API-to-device-with-mqtt---SQLITE
#API
API den http kütüphanesi ile coin verilerini json formatında okuyoruz daha sonra json formatından değişkene atıyoruz.

#MQTT
Atadığımız değişkeni MQTT yardımı ile kayıtlı cihaz Ip adresine yolluyoruz uzaktan veri gönderimi sağlamış oluyoruz. MQTT içerisinde ise abone oluyoruz ve daha sonra abone olunan yoldan veri yayınlama(Gönderme) işlemini sağlıyoruz.

#GORM
veritabanı işlemlerini gorm ile yapıyoruz her veri gönderimi yapıldığında saat ve tarih cinsinden kaydını veritabanına gerçekleştiriyor.
