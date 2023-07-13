const url = "/edit/page";

function EditPage(path) {
  const html = document.getElementById(path).innerText;

  const formData = {
    html: html,
    path: path,
  };

  console.log(path)
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
        window.location.href = getCookie("LastPath")
      } else if (response.status === 403) {
        console.error("Erişim reddedildi: Yönetici izni gerekiyor.");
      } else if (response.status === 404) {
        console.error("Dosya bulunamıyor: Kullanıcı bulunamadı");
        document.getElementById("error").style["display"] = "block";
        document.getElementById("error").innerHTML = `
                Dosya bulunamadı
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
