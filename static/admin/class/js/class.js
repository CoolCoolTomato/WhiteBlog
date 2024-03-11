function deleteClass(){
    let radios = document.getElementsByName("classId")
    let tip = document.getElementById("class_tip")
    let selectedId;
    radios.forEach(radio => {
        if(radio.checked) {
            selectedId = radio.value;
        }
    });
    if (!selectedId){
        tip.innerText = "未选择分类!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let jsonData = JSON.stringify({id: parseInt(selectedId)})
    let xhr= new XMLHttpRequest()
    xhr.open('DELETE', '/admin/class/delete')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(jsonData)
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
