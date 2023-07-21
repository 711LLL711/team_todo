axios.defaults.headers.common['Authorization'] = `Bearer ${jwtToken}`;

let changeinfo_commit = document.getElementById('complete_change');
let changeinfo_headfile = document.getElementById('file');

changeinfo_headfile.addEventListener('change',function(e){
    const head = e.target.files[0];

    if(head){
        const temp = document.getElementById('changepage_useravatar');
        const reader = new FileReader();
        reader.onload = function(e){
            temp.src = e.target.result;
        };
        reader.readAsDataURL(head);
    }
    else{};
});

changeinfo_commit.addEventListener('click',async function(){
    const newuser_head = document.getElementById('changepage_useravatar').src;
    const newuser_name = document.getElementById('change_nickname').value;
    console.log(newuser_head);
    console.log(newuser_name);
    const new_userinfo = {
        newuser_head:newuser_head,
        newuser_name:newuser_name,
    };
    try{
        const response = await axios.post('/user/profile',new_userinfo);
        const responseData = response.json();
        if(responseData.success){
            window.location.reload();
        }
        else{
            alert('上传失败！');
        }
    }catch(error){
        alert('网络不佳，上传失败！');
    }
});
