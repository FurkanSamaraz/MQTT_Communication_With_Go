# Yapılması gerekenler
Mqtt sanal cihaz kullanımı için hivemq kullanıcaz ve bir ip adresini gerekli olduğu için domain dönüştürücüyü kullanacağız.
--http://www.hivemq.com/demos/websocket-client/
--https://domaintoipconverter.com/index.php

# Sending-data-from-API-to-device-with-mqtt---SQLITE
<img align="left" alt="Go" width="70px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" style="padding-right:30px;"/>

<img align="left" alt="SQLite" width="70px" src="https://upload.wikimedia.org/wikipedia/commons/3/38/SQLite370.svg" style="padding-right:30px;"/>

<img align="left" alt="MQTT" width="70px" src="https://upload.wikimedia.org/wikipedia/commons/e/e0/Mqtt-hor.svg" style="padding-right:30px;"/>. 
 
# API
API den http kütüphanesi ile coin verilerini json formatında okuyoruz daha sonra json formatından değişkene atıyoruz.

# MQTT
Atadığımız değişkeni MQTT yardımı ile kayıtlı cihaz Ip adresine yolluyoruz uzaktan veri gönderimi sağlamış oluyoruz. MQTT içerisinde ise abone oluyoruz ve daha sonra abone olunan yoldan veri yayınlama(Gönderme) işlemini sağlıyoruz.

# GORM
veritabanı işlemlerini gorm ile yapıyoruz her veri gönderimi yapıldığında saat ve tarih cinsinden kaydını veritabanına gerçekleştiriyor.
