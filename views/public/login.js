const url = "/login";

function Login() {
  const email = document.getElementById("email").value;
  const password = document.getElementById("pass").value;

  const formData = {
    email: email,
    pass: password,
  };

  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  })
    .then((response) => {
      if (response.ok) {
        console.log("Kullanıcı başarıyla kaydedildi");
        window.location.href = getCookie("LastPath");
      } else if (response.status === 401) {
        console.error("Yetkilendirme hatası: Kullanıcı girişi gerekiyor.");
        document.getElementById("error").style["display"] = "block";
        document.getElementById("error").innerHTML = `
                Email ve ya Şifre yanlış
            `;
      } else if (response.status === 403) {
        console.error("Erişim reddedildi: Yönetici izni gerekiyor.");
      } else if (response.status === 404) {
        console.error("Dosya bulunamıyor: Kullanıcı bulunamadı");
        document.getElementById("error").style["display"] = "block";
        document.getElementById("error").innerHTML = `
                Bu email ile kaydolmuş kullanıcı bulunamadı
            `;
      } else {
        console.error("Bir hata oluştu.");
      }
    })
    .then((data) => {
      console.log("İstek başarıyla gönderildi:", data);
    })
    .catch((error) => {
      console.error("İstek gönderilirken hata oluştu:", error);
    });
}
