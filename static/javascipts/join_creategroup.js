axios.defaults.headers.common['Authorization'] = `Bearer ${jwtToken}`;

var join_firstgroup = document.getElementById('join_firstgroup');
var join_group = document.getElementById('join_group');
var commit_verifycode = document.getElementById('commit-verifycode');
var back_tohomepage = document.getElementById('back-tohomepage');
var create_firstgroup = document.getElementById('create_firstgroup');
var create_group = document.getElementById('create_group');

join_firstgroup.addEventListener('click',function(){
    show_inputverifycode();
});
join_group.addEventListener('click',function(){
    show_inputverifycode();
});

back_tohomepage.addEventListener('click',function(){
    hid_inputverifycode();
});
//邀请码提交监听
commit_verifycode.addEventListener('click',async function(){
    const number1 = document.getElementById('number1').value;
    const number2 = document.getElementById('number2').value;
    const number3 = document.getElementById('number3').value;
    const number4 = document.getElementById('number4').value;
    const verify_number = ""+number1+number2+number3 +number4+"";

    try{
        const response = await axios.get(`/groups/join?${verify_number}`);
        const data = response.data;
        if (data.success) {
            alert('join successful!');
            //群组加一，最好有动画效果
        } else {
            error_while_wrong();
        }
        } catch (error) {
            alert('加入群组失败,请检查网络!');
        }
});
//创建群组请求监听
create_firstgroup.addEventListener('click',async function(){
    let create_grouppage = document.getElementById('create-page');
    create_grouppage.style.visibility = 'visible';
});
create_group.addEventListener('click',async function(){
    let create_grouppage = document.getElementById('create-page');
    create_grouppage.style.visibility = 'visible';
});
//创建群组的请求提交

var complete_groupchange = document.getElementById('complete_creategroup');
var group_avatar = document.getElementById('groupfile');
group_avatar.addEventListener('change',function(e){
    const head = e.target.files[0];
    console.log(head);
    console.log(111);
    if(head){
        const group_avatarimg = document.getElementById('changepage_groupavatar');
        const reader = new FileReader();
        reader.onload = function(e){
            group_avatarimg.src = e.target.result;
        };
        reader.readAsDataURL(head);
    }
    else{ alert('wrong!');};
});

complete_groupchange.addEventListener('click',async function(){
    const newgroup_img = document.getElementById('changepage_groupavatar').src;
    const newgroup_name = document.getElementById('change_groupnickname').value;
    console.log(newgroup_img);
    console.log(newgroup_name);
    const newgroup_info = {
        newgroup_head:newgroup_img,
        newgroup_name:newgroup_name,
    };
    try{
        const response = await axios.post('/groups',newgroup_info);
        const responseData = response.json();
        if(responseData.success){
            showinvitecode();
        }
        else{
            alert('上传失败！');
        }
    }catch(error){
        alert('网络不佳，上传失败！');
    }
});

complete_groupchange.addEventListener('click',function(){
    let group_img = document.getElementById('changepage_groupavatar').src;
    let group_nickname = document.getElementById('change_groupnickname').value;
});


function hid_inputverifycode(){
    let blockdiv = document.getElementById('blockdiv');
    blockdiv.style.visibility = 'hidden';
}

function show_inputverifycode(){
    let blockdiv = document.getElementById('blockdiv');
    blockdiv.style.visibility = 'visible';
}

function showError(inputElement,message){
    const errorMessage = inputElement;
    errorMessage.innerHTML = '邀请码有误!';
}
function show_editgroup(){
    let create_grouppage = document.getElementById('')
};
function error_while_wrong(){
    var errorline = document.getElementById('invitecode-wrong');
    showError(errorline,'邀请码有误!');
    var input_numbers = document.getElementsByClassName('input');
    for(const input_number of input_numbers){
        input_number.addEventListener('focus',function(){
            var errorline = document.getElementById('invitecode-wrong');
            errorline.innerHTML = '';
        })
    };
};
async function showinvitecode(){
    var invitecode_number = document.getElementById('invitecode_number');
    try{
        const response = await axios.get(`groups/?${localid}/code`);
        const responseData = response.json();
        if(responseData){
            var invitecode = document.getElementById('invitecode');
            invitecode.style.visibility='visible';
            invitecode_number.innerHTML = responseData.code;
        }
        else {
            alert('获取失败，请检查网络!');
        }
    }catch(error){
        alert('获取失败，请检查网络!');
    }
};