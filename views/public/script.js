function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
  }

function deleteUser(userToken) {
    let redirect = getCookie("LastPath")
    fetch(`/delete/blog/${userToken}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => {
        if (response.ok) {
          console.log('Kullanıcı başarıyla silindi.');
        } else if (response.status === 401) {
          console.error('Yetkilendirme hatası: Kullanıcı girişi gerekiyor.');
        } else if (response.status === 403) {
          console.error('Erişim reddedildi: Yönetici izni gerekiyor.');
        } else {
          console.error('Bir hata oluştu.');
        }
        setTimeout(()=>{
            window.location.href = redirect
        },300)
      }).catch(error => {
        console.error('İstek hatası:', error);
      });
  }

  function deleteProduct(token) {
    let redirect = getCookie("LastPath")
    return fetch(`/delete/product/${token}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => {
        if (response.ok) {
            console.log('Kullanıcı başarıyla silindi.');
          } else if (response.status === 401) {
            console.error('Yetkilendirme hatası: Kullanıcı girişi gerekiyor.');
          } else if (response.status === 403) {
            console.error('Erişim reddedildi: Yönetici izni gerekiyor.');
          } else {
            console.error('Bir hata oluştu.');
          }
          setTimeout(()=>{
              window.location.href = redirect
          },300)
      }).catch(error => {
        console.error('İstek hatası:', error);
      });
  }
  
  function deleteBlog(token) {
    let redirect = getCookie("LastPath")
    return fetch(`/delete/blog/${token}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => {
        if (response.ok) {
            console.log('Kullanıcı başarıyla silindi.');
          } else if (response.status === 401) {
            console.error('Yetkilendirme hatası: Kullanıcı girişi gerekiyor.');
          } else if (response.status === 403) {
            console.error('Erişim reddedildi: Yönetici izni gerekiyor.');
          } else {
            console.error('Bir hata oluştu.');
          }
          setTimeout(()=>{
              window.location.href = redirect
          },300)
      }).catch(error => {
        console.error('İstek hatası:', error);
      });
  }
  
  function deleteGalery(token) {
    let redirect = getCookie("LastPath")
    return fetch(`/delete/galery/${token}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => {
        if (response.ok) {
            console.log('Kullanıcı başarıyla silindi.');
          } else if (response.status === 401) {
            console.error('Yetkilendirme hatası: Kullanıcı girişi gerekiyor.');
          } else if (response.status === 403) {
            console.error('Erişim reddedildi: Yönetici izni gerekiyor.');
          } else {
            console.error('Bir hata oluştu.');
          }
          setTimeout(()=>{
              window.location.href = redirect
          },300)
      }).catch(error => {
        console.error('İstek hatası:', error);
      });
  }