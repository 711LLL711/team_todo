
document.getElementById('register').addEventListener('click',async function(){
    const new_username = document.getElementById('newuser_name').value;
    const new_useremail = document.getElementById('newuser_email').value;
    const new_userverifycode = document.getElementById('newuser_verifycode').value;
    const new_userpassword = document.getElementById('newuser_password').value;
    const new_userpassword_again = document.getElementById('newuser_password_again').value;
    empty_by_register();
    if(new_userpassword!==''&&new_userpassword_again!==''&&new_username!==''&&new_useremail!==''&&new_userverifycode!==''){
        const keyLength = 32; // 256位密钥长度
        let randomKey = '';
        const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
        for (let i = 0; i < keyLength; i++) {
            randomKey += characters.charAt(Math.floor(Math.random() * characters.length));
        }

        console.log(randomKey);

        const newuser_data = {
            new_username:new_username,
            new_useremail:new_useremail,
            new_userverifycode:new_userverifycode,
            new_userpassword:new_userpassword,
        };
    
        try{
            const response = await axios('/user/register', newuser_data);

            const responseData = await response.json();

            if (responseData.success) {
                let input_verifycode = document.getElementById('newuser_verifycode');
                showError(input_verifycode.parentNode,'验证码不能为空！');
            } else {
                alert('Registration failed: ' + responseData.message);
            }
            } catch (error) {
                alert('Error occurred during registration: ' + error.message);
        }
    };
});

document.getElementById('login').addEventListener('click',async function(){
    const old_username = document.getElementById('olduser_name').value;
    const old_userpassword = document.getElementById('olduser_password').value;
    empty_by_login();
    if(old_username.trim()&&old_userpassword.trim()){
        
        const olduser_data = {
            old_username:old_username,
            old_userpassword:old_userpassword
        };
        try{
            const response = await axios('/user/login', olduser_data);

        const responseData = await response.json();

        if (responseData.success) {
            localStorage.setItem('jwt',responseData.token);
            localStorage.setItem('localid',responseData.id);
            window.location.href = 'home.html';
        } else {
            alert('登录失败: ' + responseData.message);
        }
        } catch (error) {
            alert('网络异常: ' + error.message);
        }
    }
});

document.getElementById('verifycode_get').addEventListener('click',async function(){
    const new_useremail = document.getElementById('newuser_email').value;
    let new_userverifycode = document.getElementById('newuser_verifycode');
    if(new_useremail!==''){
        const email = {
            email:new_useremail,
        };
        try {
            const response = await axios('/user/verify-code',email);

        const responseData = await response.json();
    
        if(responseData.success){
            showError(new_userverifycode.parentNode,"验证码已发送");
        }
        else{
            showError(new_userverifycode.parentNode,"请xx秒后再试!");
        }
        }catch (error) {
            alert('Error occurred');
        }
    }
});
function empty_by_register(){
    let input_username = document.getElementById('newuser_name');
    let input_email = document.getElementById('newuser_email');
    let input_verifycode = document.getElementById('newuser_verifycode');
    let input_password_again = document.getElementById('newuser_password_again');
    let input_password = document.getElementById('newuser_password');
    if(input_username.value.trim()===''){
        showError(input_username,'用户名不能为空！');
    }
    if(input_email.value.trim()===''){
        showError(input_email,'邮箱不能为空！');
    }
    if(input_verifycode.value.trim()===''){
        showError(input_verifycode.parentNode,'验证码不能为空！');
    }
    if(input_password.value.trim()===''){
        showError(input_password,'密码不能为空！');
    }
    if(input_password.value !== input_password_again.value){
        showError(input_password_again,'两次密码不匹配！');
    }
    input_username.addEventListener('focus',function(){
        resetInputStyle(input_username);
    });
    input_email.addEventListener('focus',function(){
        resetInputStyle(input_email);
    });
    input_password.addEventListener('focus',function(){
        resetInputStyle(input_password);
    });
    input_verifycode.addEventListener('focus',function(){
        resetInputStyle(input_verifycode.parentNode);
    });
    input_password_again.addEventListener('focus',function(){
        resetInputStyle(input_password_again);
    });
}

function showError(inputElement, message) {
    // 获取对应的错误提示元素
    const errorMessage = inputElement.nextElementSibling;
    errorMessage.style.display = 'block';
}
function resetInputStyle(inputElement) {
    // 隐藏错误提示
    const errorMessage = inputElement.nextElementSibling;
    errorMessage.style.display = 'none';
}

function empty_by_login(){
    let input_oldusername = document.getElementById('olduser_name');
    let input_olduserpwd = document.getElementById('olduser_password');
    if(input_oldusername.value.trim()===''){
        showError(input_oldusername,'用户名不能为空！');
    }
    if(input_olduserpwd.value.trim()===''){
        showError(input_olduserpwd,'密码不能为空！');
    }
    input_oldusername.addEventListener('focus',function(){
        resetInputStyle(input_oldusername);
    });
    input_olduserpwd.addEventListener('focus',function(){
        resetInputStyle(input_olduserpwd);
    });
}
