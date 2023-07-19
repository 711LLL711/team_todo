
document.getElementById('register').addEventListener('click',async function(){
    const new_username = document.getElementById('newuser_name').value;
    const new_useremail = document.getElementById('newuser_email').value;
    const new_userpassword = document.getElementById('newuser_password').value;
    const new_userpassword_again = document.getElementById('newuser_password_again').value;
    if(new_userpassword!=new_userpassword_again){
        return error;
    }
    else {
        const token = 'your_generated_token';

        const newuser_data = {
            new_username:new_username,
            new_useremail:new_useremail,
            new_userpassword:new_userpassword,
        };

        try{
            const response = await fetch('/user/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(newuser_data)
            });

        const responseData = await response.json();

        if (responseData.success) {
            alert('Registration successful!');
            // 在这里可以进行跳转或其他后续处理
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

        const token = 'your_generated_token';

        const olduser_data = {
            old_username:old_username,
            old_userpassword:old_userpassword
        };
        try{
            const response = await fetch('/user/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(olduser_data)
            });

        const responseData = await response.json();

        if (responseData.success) {
            alert('Registration successful!');
            // 在这里可以进行跳转或其他后续处理
            window.location.href ='url';
        } else {
            alert('Registration failed: ' + responseData.message);
        }
        } catch (error) {
            alert('Error occurred during registration: ' + error.message);
        }

});
