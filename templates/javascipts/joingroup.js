const jwtToken = localStorage.getItem('jwt');
const localid = localStorage.getItem('id');
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
    try {
        const response = await axios.get('/groups');
        const responseData = response.data;
        if(responseData.success){
            alert('create successfully!');
            //群组加一
        }
        else{
            alert('群组数量过多/距离上次创建时间太短!');
            //简单报错
        }

    }catch(error){
        alert('error network!');
    }
});
create_group.addEventListener('click',async function(){
    try {
        const response = await axios.get('/groups');
        const responseData = response.data;
        if(responseData.success){
            alert('create successfully!');
            //群组加一
        }
        else{
            alert('群组数量过多/距离上次创建时间太短!');
            //简单报错
        }

    }catch(error){
        alert('error network!');
    }
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

}
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
