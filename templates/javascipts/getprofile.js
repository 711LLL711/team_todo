axios.defaults.headers.common['Authorization'] = `Bearer ${jwtToken}`;

async function ondomcontentloaded(){
    let useravatar =  document.getElementById('avatar');
    let usernickname = document.getElementById('usersID');
    let changepage_avatar = document.getElementById('changepage_useravatar')
    let changepage_name = document.getElementById('changepage_username');
    try {
        const userprofile = await axios.get(`/users/${usernickname}/profile`);
        const userprofile_response = userprofile.data;
        if(userprofile_response){
            changepage_avatar.src = userprofile_response.avatar;
            changepage_name.innerHTML = userprofile_response.nickname;
            useravatar.src = userprofile_response.avatar;
            usernickname.innerHTML = userprofile_response.nickname;
        }else{

            useravatar.src = '../img/user-regular.svg';
            usernickname.innerHTML = '';
            changepage_avatar.src = '../img/user-regular.svg';
            changepage_name.innerHTML = '';
        }
            
    }catch(error){
        useravatar.src = '../img/user-regular.svg';
        usernickname.innerHTML = '';
        changepage_avatar.src = '../img/user-regular.svg';
        changepage_name.innerHTML = '默认idxxx';
    }
}

document.addEventListener('DOMContentLoaded', ondomcontentloaded);

