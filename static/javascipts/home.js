const jwtToken = localStorage.getItem('jwt');
const localid = localStorage.getItem('id');
axios.defaults.headers.common['Authorization'] = `Bearer ${jwtToken}`;


let task_botton = document.getElementById('tasks-botton');
let group_botton = document.getElementById('groups-botton');
let security_botton = document.getElementById('security-botton');
let about_me_botton = document.getElementById('about_me-botton');
group_botton.addEventListener('click',function(){
    let group = document.getElementById('groups');
    let task = document.getElementById('tasks');
    let group_content=document.getElementById('groups-content');
    let right_box = document.getElementsByClassName('right-content')[0];
    right_box.style.width='45%';
    right_box.style.transition = "all 1.3s ease-in-out"
    group.style.height='98%';
    group.style.width='70%';
    group.style.position='absolute'
    group.style.transition = "all 1.3s ease-in-out";
    task.style.visibility='hidden';
    if(group.style.visibility=='hidden'){
        group.style.visibility='visible';
        group.style.zIndex=1;
        task.style.zIndex=0;
    }
    group_content.style.top='12%';
    group_content.style.borderRadius='0';
    group_content.style.height='86%';
    showtype3elems();
});  
function showtype3elems(){
    const type3elements = document.querySelectorAll('.type3');
        for(const element of type3elements){
            element.style.transition= 'all 1.3s ease-in-out';
            element.style.visibility='visible';
        }
};

task_botton.addEventListener('click',function(){
    let task = document.getElementById('tasks');
    let group = document.getElementById('groups');
    let task_content=document.getElementById('tasks-content');
    let right_box = document.getElementsByClassName('right-content')[0];
    right_box.style.width='40%';
    right_box.style.transition = "all 1.3s ease-in-out"
    task.style.height='98%';
    task.style.width='70%';
    task.style.position='absolute'
    task.style.transition = "all 1.3s ease-in-out";
    if(task.style.visibility=='hidden'){
        task.style.visibility='visible';
        group.style.zIndex=0;
        task.style.zIndex=1;
    }
    group.style.visibility='hidden';
    task_content.style.top='12%';
    task_content.style.borderRadius='0';
    task_content.style.height='86%';
    })  

//群组内容的查询,任务内容查询也同理。
const all_biggergroups = document.querySelectorAll('.bigger_simplegroup');
all_biggergroups.forEach(li =>{
    li.addEventListener('click',function(){
        try {
            const response = axios.get(`/groups/${localid}/info`)
        }catch(error){
            alert('查询失败，请检查网络！');
        }
        let detail=document.getElementById('details_content');
        const groupid = li.querySelector('h3').innerHTML;
        var details_of_group = new Array();
        const newdetail = document.createElement("div");
        newdetail.innerHTML="<span><img src=\"../img/群组头像.jpg\" alt=\"\" id=\"chosen_one_head\"></span><label for=\"\">"+groupid+"</label>"
        newdetail.classList.add("single_detail","chosen_one");
        detail.innerHTML='';
        console.log(newdetail);
        detail.appendChild(newdetail);
    })
})
//homepage_botton
function refreshPage() {
    location.reload();
    }      
document.getElementById('homepage_botton').addEventListener('click',function(){
    refreshPage();
});

//contact_us
//随便写的
let contact_us = document.getElementById('contact_us_button');
contact_us.addEventListener('click',function(){
    window.location.href = 'https://github.com/711LLL711/team_todo/blob/main/templates/javascipts/sign_&_log.js';
});

//个人主页
about_me_botton.addEventListener('click',function(){
    let mypage = document.getElementById('my-page2');
    mypage.style.visibility = 'visible';

})

    //日期盒子
var currentDate = new Date();
var monthNumber = currentDate.getMonth() + 1;
function getMonthName(monthNumber) {
    const monthNames = [
        "Jan", "Feb", "Mar", "Apri", "May", "June",
        "July", "Aug", "Sep", "Oct", "Nov", "Dec"
    ];
    return monthNames[monthNumber - 1] || 'Unknown';
}
let month_box = document.getElementById('month');
let date_box = document.getElementById('date');
month_box.textContent = getMonthName(monthNumber);
date_box.textContent = currentDate.getDate();

