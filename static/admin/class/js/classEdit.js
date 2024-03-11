function editClass(){
    let form = document.getElementById("class_edit")
    let tip = document.getElementById("class_tip")
    let formData = JSON.stringify({
        id: parseInt(form.cid.value),
        name: form.cname.value,
        belong: parseInt(form.belong.value),
        created_date: form.created_date.value+":00Z",
        updated_date: form.updated_date.value+":00Z"
    })

    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/class/edit?id=' + form.cid.value)
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
