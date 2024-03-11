/*默认时间*/
window.addEventListener('load', () => {
    let datetimepicker = document.querySelectorAll('input[type="datetime-local"]');
    let now = new Date();
    let month = now.getMonth() + 1;
    if(month < 10) {
        month = '0' + month;
    }
    let date = now.getDate();
    if(date < 10) {
        date = '0' + date;
    }
    let hour = now.getHours()
    if (hour < 10){
        hour = '0' + hour
    }
    let min = now.getMinutes()
    if (min < 10){
        min = '0' + min
    }
    const datetime = now.getFullYear() + '-' + month + '-' + date + "T" + hour + ":" + min;
    datetimepicker.forEach(picker => {
        picker.value = datetime;
    })
})
function addClass(){
    let form = document.getElementById("class_add")
    let tip = document.getElementById("class_tip")
    let formData = JSON.stringify({
        name: form.cname.value,
        belong: parseInt(form.belong.value),
        created_date: form.created_date.value+":00Z",
        updated_date: form.updated_date.value+":00Z"
    })
    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/class/add')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(formData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE){
            if (xhr.status === 200){
                window.location.reload()
            }
            else {
                tip.innerText = "发送失败"
                setTimeout(() => {
                    tip.innerText = ''
                }, 3000)
            }
        }
    }
}
