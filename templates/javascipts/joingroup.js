
var join_firstgroup = document.getElementById('join_firstgroup');
var join_group = document.getElementById('join_group');
var commit_verifycode = document.getElementById('commit-verifycode');
var back_tohomepage = document.getElementById('back-tohomepage');


join_firstgroup.addEventListener('click',function(){
    show_inputverifycode();
});
join_group.addEventListener('click',function(){
    show_inputverifycode();
});

back_tohomepage.addEventListener('click',function(){
    hid_inputverifycode();
});

commit_verifycode.addEventListener('click',async function(){
    const number1 = document.getElementById('number1').value;
    const number2 = document.getElementById('number2').value;
    const number3 = document.getElementById('number3').value;
    const number4 = document.getElementById('number4').value;
    const verify_number = ""+number1+number2+number3 +number4+""
    try{
        const response = await axios.get(`/?${verify_number}`);
        const data = response.data;
        if (data.success) {
            alert('join successful!');
        } else {
            showError();
        }
        } catch (error) {
            alert('Error occurred during registration: ' + error.message);
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
    errorMessage.style.display = 'block';
}