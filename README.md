# Todo Uygulaması

Bu proje, kullanıcıların yapılacak işleri listeleyebilecekleri ve yönetebilecekleri modern bir web uygulamasıdır. Go dili ile yazılmış bir backend RestAPI ve React ile geliştirilmiş bir frontend arayüzünden oluşmaktadır.

## Proje Hakkında

Bu Todo uygulaması, kullanıcıların yapılacaklar listelerini oluşturmasına, düzenlemesine ve takip etmesine olanak tanır. Temiz bir mimari (clean architecture) prensiplerini takip eden backend ve modern bir kullanıcı arayüzü sunan frontend bileşenlerinden oluşur.

### Temel Özellikler

- **Kullanıcı Kimlik Doğrulama Sistemi**: JWT (JSON Web Token) tabanlı güvenli kimlik doğrulama sistemi
- **İki Tür Kullanıcı**: Normal kullanıcılar ve admin kullanıcıları
- **Todo Listeleri Yönetimi**: Listeleri oluşturma, görüntüleme, güncelleme ve silme işlemleri
- **Todo Öğeleri Yönetimi**: Liste içindeki görevleri oluşturma, güncelleme ve silme imkanı
- **Yumuşak Silme (Soft Delete)**: Silinen verilerin tamamen silinmeden işaretlenmesi
- **Tamamlanma Yüzdesi**: Her liste için tamamlanma oranı hesaplanması
- **Rol Tabanlı Erişim**: Admin kullanıcıları tüm listeleri görebilirken, normal kullanıcılar sadece kendi listelerini görüntüleyebilir

## Teknik Özellikler

### Backend (Go)

- **Programlama Dili**: Go (Golang)
- **Web Framework**: Gorilla Mux (HTTP Router)
- **Kimlik Doğrulama**: JWT (github.com/dgrijalva/jwt-go)
- **CORS Desteği**: github.com/rs/cors kütüphanesi kullanılarak Cross-Origin Resource Sharing yönetimi
- **Veri Depolama**: Bellek içi depolama (gerçek bir veritabanı yerine mock servis kullanılmıştır)

### Frontend (React)

- **Framework**: React 19.1.0
- **Routing**: React Router DOM 7.5.3
- **UI Bileşenleri**: Material-UI 7.0.2
- **HTTP İstekleri**: Axios 1.9.0
- **State Management**: React Context API (AuthContext)

## Proje Yapısı

```
/
├── todo-app/                # Backend uygulaması
│   ├── cmd/                 # Uygulamanın giriş noktası
│   │   └── main.go         # Ana uygulama başlangıcı
│   ├── internal/            # İç bileşenler
│   │   ├── handlers/        # HTTP istek işleyicileri
│   │   ├── middleware/      # Ara yazılımlar (Auth vb.)
│   │   ├── models/          # Veri modelleri
│   │   └── services/        # İş mantığı servisleri
│   ├── pkg/                 # Dış paketler
│   │   ├── auth/            # Kimlik doğrulama
│   │   └── utils/           # Yardımcı fonksiyonlar
│   ├── go.mod               # Go bağımlılık yönetimi
│   ├── go.sum               # Go bağımlılık sağlaması
│   └── README.md            # Backend dokümantasyonu
│
└── todo-frontend/           # Frontend uygulaması
    ├── public/              # Statik dosyalar
    ├── src/                 # Kaynak kodları
    │   ├── components/      # React bileşenleri
    │   │   ├── Login.js     # Giriş ekranı
    │   │   └── TodoLists.js # Todo listeleri ekranı
    │   ├── context/         # React context
    │   │   └── AuthContext.js # Kimlik doğrulama contexti
    │   ├── App.js           # Ana uygulama bileşeni
    │   └── index.js         # Frontend giriş noktası
    ├── package.json         # NPM bağımlılıkları
    └── README.md            # Frontend dokümantasyonu
```

## API Endpoints

### Kimlik Doğrulama
- `POST /login` - Giriş yapma ve JWT token alma

### Todo Listeleri
- `GET /api/todos` - Tüm todo listelerini getirme (admin tümünü, kullanıcı kendininkini)
- `POST /api/todos` - Yeni bir todo listesi oluşturma
- `GET /api/todos/{id}` - Belirli bir todo listesini getirme
- `PUT /api/todos/{id}` - Todo listesini güncelleme
- `DELETE /api/todos/{id}` - Todo listesini silme (soft delete)

### Todo Öğeleri
- `GET /api/todos/{listId}/items` - Bir todo listesindeki tüm öğeleri getirme
- `POST /api/todos/{listId}/items` - Yeni bir todo öğesi oluşturma
- `PUT /api/todos/{listId}/items/{itemId}` - Todo öğesini güncelleme
- `DELETE /api/todos/{listId}/items/{itemId}` - Todo öğesini silme (soft delete)

## Kurulum

### Gereksinimler
- Go 1.18 veya üzeri
- Node.js 16 veya üzeri
- npm veya yarn

### Backend Kurulumu
1. Proje klasörüne gidin:
   ```bash
   cd todo-app
   ```

2. Bağımlılıkları yükleyin:
   ```bash
   go mod download
   ```

3. Uygulamayı başlatın:
   ```bash
   go run cmd/main.go
   ```
   
4. Uygulama şu adreste çalışacaktır: `http://localhost:8080`

### Frontend Kurulumu
1. Frontend klasörüne gidin:
   ```bash
   cd todo-frontend
   ```

2. Bağımlılıkları yükleyin:
   ```bash
   npm install
   ```
   veya
   ```bash
   yarn install
   ```

3. Geliştirme sunucusunu başlatın:
   ```bash
   npm start
   ```
   veya
   ```bash
   yarn start
   ```

4. Uygulama tarayıcınızda şu adreste açılacaktır: `http://localhost:3000`

## Kullanım

### Giriş Yapma
Uygulama ile birlikte gelen ön tanımlı kullanıcılar:

1. Normal Kullanıcı:
   - Kullanıcı Adı: `user1`
   - Şifre: `password1`

2. Admin Kullanıcısı:
   - Kullanıcı Adı: `admin`
   - Şifre: `admin123`

### Kimlik Doğrulama
Login sayfasından giriş yaptıktan sonra bir JWT token alacaksınız. Bu token, tarayıcının localStorage'ında saklanır ve sonraki API isteklerinde kullanılır.

### Todo Listeleri
- Ana sayfada todo listeleri görüntülenir
- Her liste için tamamlanma yüzdesi hesaplanır
- "Yeni Liste" butonu ile yeni todo listesi oluşturulabilir

### Todo Öğeleri
- Bir listeye tıklayarak içindeki öğeleri görüntüleyebilirsiniz
- "Yeni Öğe" butonu ile listeye yeni görev ekleyebilirsiniz
- Her öğe için tamamlanma durumunu değiştirebilirsiniz
- Öğeleri düzenleyebilir veya silebilirsiniz

## Kodlama Standartları ve Mimari

### Backend

- **Temiz Mimari**: Uygulama, bağımsız katmanlar halinde yapılandırılmıştır
- **Middleware Kullanımı**: Kimlik doğrulama ve güvenlik kontrolleri için middleware'ler
- **Servis Tabanlı**: İş mantığı servis katmanında yürütülür

### Frontend

- **Bileşen Tabanlı**: Tekrar kullanılabilir React bileşenleri
- **Context API**: Uygulama genelinde kimlik doğrulama durumunun yönetilmesi
- **Material-UI**: Modern ve duyarlı bir kullanıcı arayüzü için UI framework'ü
- **Route Koruması**: Kimlik doğrulaması olmayan kullanıcılar için korumalı rotalar

## Güvenlik Özellikleri

- **JWT Tabanlı Kimlik Doğrulama**: Güvenli token tabanlı yetkilendirme
- **Role Dayalı Erişim Kontrolü**: Farklı kullanıcı türleri için farklı erişim hakları
- **CORS Yapılandırması**: Cross-Origin istekleri için güvenli yapılandırma
- **Şifre Koruması**: Şifreler modellerde gösterilmez (`json:"-"`)

## Geliştirme ve Üretim

### Geliştirme Ortamı
- Backend ve frontend'i ayrı ayrı geliştirme sunucularında çalıştırılabilir
- Frontend uygulaması, geliştirme sırasında `http://localhost:3000` adresinde çalışır
- Backend API, `http://localhost:8080` adresinde hizmet verir

### Üretim Ortamına Geçiş
- Frontend için üretim derlemesi:
  ```bash
  cd todo-frontend
  npm run build
  ```

- Üretim derleme çıktısı, statik bir web sunucusu üzerinden sunulabilir
- Backend, doğrudan sunucu üzerinde veya Docker konteyneri içinde çalıştırılabilir

## Not

- Uygulama şu anda gerçek bir veritabanı yerine bellek içi depolama kullanmaktadır
- Tüm silme işlemleri soft delete olarak gerçekleştirilir (veriler tamamen silinmez, silindi olarak işaretlenir)
- Tüm zaman damgaları UTC formatındadır