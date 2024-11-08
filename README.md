  <a href="https://www.python.org" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/python/python-original.svg" alt="python" width="40" height="40"/> 
    </a> <br>
    <h1>Glitch</h1>
    
Bu kod, bir Flask uygulamasını tanımlar ve çalıştırır. Flask, Python dilinde yazılmış bir mikro web framework'tür. Kodun ne yaptığını ve nasıl dağıtılacağını aşağıda özetledim:

Kodun Yaptıkları:
//Produced By. K.Umut Araz
Flask Uygulaması Oluşturma:
Flask sınıfını kullanarak bir web uygulaması oluşturulur.
Rota Tanımlama:
Ana rota ("/") tanımlanır ve bu rotaya bir HTTP GET isteği gönderildiğinde, JSON formatında bir yanıt döndürülür. Yanıt, {"msg": "hello deploy from flask glitch"} şeklindedir.
Uygulamayı Çalıştırma:


Eğer bu dosya doğrudan çalıştırılırsa, Flask sunucusu başlatılır (app.run()).
Dağıtım İçin Yapılacaklar:
Kodun Hazırlanması:

main.py adlı bir dosya oluşturulacak ve yukarıdaki kod içine yerleştirilecektir.
Gunicorn Kullanımı:

Gunicorn, Python WSGI HTTP Sunucusu'dur. Flask uygulamasını daha üretken bir ortamda çalıştırmak için kullanılır.
gunicorn main:app -w 1 --log-file - komutunu çalıştırarak uygulama başlatılır. Bu komut, main.py dosyasındaki app isimli Flask uygulamasını çalıştırır. -w 1 ile bir worker (işçi) işlemi belirtilir, --log-file - ise günlüklerin standart çıktıya (stdout) yazılmasını sağlar.
GitHub ve Glitch Üzerinden Dağıtım:

Kodunuzu GitHub'a yükleyin.
Glitch platformuna giriş yapın ve yeni bir proje oluşturun.
Glitch projenizi GitHub deposuna bağlayın ve kodunuzu içe aktarın.
Glitch, start betiği olarak gunicorn main:app -w 1 --log-file - komutunu kullanarak uygulamanızı başlatır.
Bu adımlar sonucunda, Flask uygulamanız Glitch üzerinde dağıtılmış ve erişilebilir olacaktır.



 <a href="https://golang.org/" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> 
    </a>
    <h1>Chat Platform 2</h1>
    Proje Amacı
  
  
Bu proje, Go programlama dili kullanılarak geliştirilen basit bir blog platformu uygulamasıdır. Proje, kullanıcıların kayıt olup giriş yapabileceği, blog yazıları oluşturabileceği, bu yazılara yorum ekleyebileceği ve veritabanında saklayabileceği temel bir web uygulaması sağlar.

Kullanılan Teknolojiler
Go (Golang): Uygulamanın geliştirilmesinde kullanılan programlama dili.
Gin: Go ile hızlı ve kolay bir web sunucusu geliştirmeyi sağlayan bir web framework’ü.,
GORM: Go için güçlü bir ORM (Object-Relational Mapping) kütüphanesi. Veritabanı işlemlerini yönetmek için kullanılır.
SQLite: Hafif ve sunucusuz bir veritabanı yönetim sistemi. Bu projede veritabanı olarak kullanıldı.
Proje Yapısı ve Dosyalar
main.go: Uygulamanın giriş noktasıdır. Bu dosyada, web sunucusu başlatılır, veritabanı bağlantısı kurulur ve HTTP rotaları yapılandırılır.

models.go: Bu dosyada, uygulamanın veri modelleri tanımlanmıştır. Projede üç ana model kullanılır:

User: Kullanıcı bilgilerini saklayan modeldir (kullanıcı adı ve şifre).
Post: Blog yazılarını temsil eden modeldir (başlık, içerik ve kullanıcıya ait ID).
Comment: Blog yazılarına eklenen yorumları saklayan modeldir (yorum içeriği ve ilgili yazıya ait ID).,
database.go: Veritabanı yapılandırması ve bağlantısının kurulduğu dosyadır. SQLite veritabanına bağlanır ve veri modelleri için tablolar oluşturur.

routes.go: HTTP rotalarının (endpoints) tanımlandığı dosyadır. Bu rotalar, gelen HTTP isteklerini belirli işlevlere yönlendirir.

/register: Yeni kullanıcı kaydı oluşturur.
/login: Kullanıcı oturumu açma işlevini sağlar (bu işlev henüz uygulamada eksik).
/posts: Yeni bir blog yazısı oluşturur.
/comments: Bir blog yazısına yeni yorum ekler.
handlers.go: HTTP rotaları ile ilişkili işlevlerin uygulandığı dosyadır. Gelen istekleri işler ve veritabanı ile etkileşime geçer. Kullanıcı kaydı, blog yazısı oluşturma ve yorum ekleme işlevlerini içerir.

Uygulamanın İşlevleri
Kullanıcı Kayıt ve Oturum Açma:

Kullanıcılar /register rotası üzerinden kayıt olabilirler. Bu işlem sırasında, kullanıcı bilgileri veritabanına kaydedilir.
Kullanıcı oturumu açma işlevi (henüz eksik olan) /login rotası üzerinden yapılacaktır.
Blog Yazısı Oluşturma:

Kayıtlı bir kullanıcı, /posts rotası üzerinden yeni bir blog yazısı oluşturabilir. Kullanıcının ID'si, yazıya otomatik olarak atanır.
Yorum Ekleme:

Kullanıcılar, /comments rotası üzerinden mevcut blog yazılarına yorum ekleyebilirler. Yorumlar, ilgili yazının ID'si ile ilişkilendirilir ve veritabanına kaydedilir.
Veritabanı Yapısı
User Tablosu:

ID: Kullanıcıyı benzersiz olarak tanımlayan birincil anahtar.
Username: Kullanıcının adı (benzersiz olmalı).
Password: Kullanıcının şifresi (hashlenmiş olabilir).
Post Tablosu:

ID: Blog yazısını benzersiz olarak tanımlayan birincil anahtar.
Title: Yazının başlığı.
Content: Yazının içeriği.
UserID: Yazıyı oluşturan kullanıcıya ait ID.
Comment Tablosu:

ID: Yorumu benzersiz olarak tanımlayan birincil anahtar.
Content: Yorum içeriği.
PostID: Yorumun yapıldığı yazıya ait ID.
Sonuç ve Geliştirme Önerileri
Bu proje, Go dilini ve onun popüler framework'lerini (Gin ve GORM) öğrenmek için iyi bir başlangıç projesidir. Uygulama, temel blog işlevlerini kapsar ve bu işlevlerin Go dilinde nasıl uygulanacağını gösterir.

Geliştirme Önerileri:

Kullanıcı Oturumu ve Kimlik Doğrulama: /login rotası üzerinden kullanıcı oturumu açma ve JWT (JSON Web Token) ile kimlik doğrulama eklenebilir.
Yazı ve Yorum Düzenleme: Kullanıcıların yazı ve yorumlarını düzenleyebileceği işlevler eklenebilir.
Yazı Listesi Görüntüleme: Ana sayfada tüm yazıların listelendiği bir endpoint eklenebilir.
Güvenlik İyileştirmeleri: Şifrelerin hashlenmesi ve kimlik doğrulama gibi güvenlik önlemleri eklenebilir.
Frontend Arayüzü: Basit bir HTML/CSS arayüzü eklenerek, kullanıcıların bu işlevleri bir web arayüzü üzerinden gerçekleştirmesi sağlanabilir.

