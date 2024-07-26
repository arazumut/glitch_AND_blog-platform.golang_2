  <a href="https://www.python.org" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/python/python-original.svg" alt="python" width="40" height="40"/> 
    </a> <br>
Bu kod, bir Flask uygulamasını tanımlar ve çalıştırır. Flask, Python dilinde yazılmış bir mikro web framework'tür. Kodun ne yaptığını ve nasıl dağıtılacağını aşağıda özetledim:

Kodun Yaptıkları:
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
