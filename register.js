const url = '/register';

function ValidateEmail(mail) {
    return(/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(mail))
}

function ValidatePassword(password) {
    let error = null;
    if (password.includes(" ")) {
        error = "Şifre boşluk içeremez";
    }
  
    if (password.length < 8) {
        error = "Şifre en az 8 karakter olmak zorundadır.";
    }

    if (password.length > 32){
        error = "Şifre en fazla 32 karakter olmalıdır.";
    }

    if (!/\d/.test(password)) {
        error = "Şifrede sayılar yer almalıdır.";
    }

    if (!/[a-zA-Z]/.test(password)) {
        error = "Şifrede metin bulunmak zorundadır.";
    }
  
    if (!/[!@#$%^&*(),.?":{}\|_<>]/.test(password)) {
        error = "Şifrede özel karakter bulunmak zorundadır.";
    }

    if(error){
        document.getElementById("error").style["display"] = "block";
        document.getElementById("error").innerHTML = error;
    }

    return !error;
}

function ValidateUsername(username){
    if (username.includes(" ")) {
        return false;
    }

    return true
}

function Register(){
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('pass').value;

    const formData = {
        username: username,
        email: email,
        pass: password
    };

    if(!username){
        document.getElementById("error").style["display"] = "block";
        document.getElementById("error").innerHTML = `
            Kullanıcı adı boş bırakılamaz.
        `;
    }

    if(!ValidateUsername(username)){
        document.getElementById("error").style["display"] = "block";
        document.getElementById("error").innerHTML = `
            Kullanıcı adı boşluk içermemelidir.
        `;
    }

    if(!ValidateEmail(email)){
        document.getElementById("error").style["display"] = "block";
        document.getElementById("error").innerHTML = `
            Bu Email geçersiz.
        `;
    }


    if(username && ValidateUsername(username) && ValidateEmail(email) && ValidatePassword(password)){
        fetch(url, {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json'
        },
            body: JSON.stringify(formData)
        })
        .then(response => {
            if (response.ok) {
                console.log('Kullanıcı başarıyla kaydedildi');
                window.location.href = getCookie("LastPath")
            } else if (response.status === 401) {
                console.error('Yetkilendirme hatası: Kullanıcı girişi gerekiyor.');
                document.getElementById("error").style["display"] = "block";
                document.getElementById("error").innerHTML = `
                    Bu email kullanılıyor
                `;
            } else if (response.status === 403) {
                console.error('Erişim reddedildi: Yönetici izni gerekiyor.');
            } else {
                console.error('Bir hata oluştu.');
            }
        })
        .then(data => {
            console.log('İstek başarıyla gönderildi:', data);
        })
        .catch(error => {
            console.error('İstek gönderilirken hata oluştu:', error);
        });
    }
}