function UpdateLink(){
    let form = document.getElementById("link_update")
    let tip = document.getElementById("link_tip")
    let formData = JSON.stringify({
        id: parseInt(form.lid.value),
        title: form.ltitle.value,
        gravatar: form.lgravatar.value,
        description: form.ldescription.value,
        url: form.llink.value,
        email: form.lmail.value
    })
    let xhr = new XMLHttpRequest()
    xhr.open('POST', '/admin/link/edit')
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