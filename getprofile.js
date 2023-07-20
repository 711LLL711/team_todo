
function ondomcontentloaded(){
    let useravatar =  document.getElementById('avatar');
    let userid = document.getElementById('usersID');
    axios.get("/users/${userID}/profile")
    .then(response =>{
        useravatar.src = response.avatar;
        userid.innerHTML = response.id; 
    })
    .catch(_error =>{
        useravatar.src = '../img/user-regular.svg';
        userid.innerHTML = '';
    })
}

document.addEventListener('DOMContentLoaded',ondomcontentloaded);

