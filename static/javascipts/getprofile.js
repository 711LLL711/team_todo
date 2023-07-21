axios.defaults.headers.common['Authorization'] = `Bearer ${jwtToken}`;

async function ondomcontentloaded(){
    let useravatar =  document.getElementById('avatar');
    let usernickname = document.getElementById('usersID');
    let changepage_avatar = document.getElementById('changepage_useravatar')
    let changepage_name = document.getElementById('changepage_username');
    try {
        const userprofile = await axios.get(`/users/?${localid}/profile`);
        if(!userprofile.error){
            changepage_avatar.src = userprofile.avatar;
            changepage_name.innerHTML = userprofile.nickname;
            useravatar.src = userprofile.avatar;
            usernickname.innerHTML = userprofile.nickname;
        }else{

            useravatar.src = '../images/user-regular.svg';
            usernickname.innerHTML = '';
            changepage_avatar.src = '../images/user-regular.svg';
            changepage_name.innerHTML = '';
        }
            
    }catch(error){
        useravatar.src = '../images/user-regular.svg';
        usernickname.innerHTML = '';
        changepage_avatar.src = '../images/user-regular.svg';
        changepage_name.innerHTML = '默认idxxx';
    }
}

document.addEventListener('DOMContentLoaded', ondomcontentloaded);

