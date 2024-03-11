function checkLink(self){
    self.children[0].checked = self.children[0].checked !== true;
}


function showAddForm(){
    let form = document.getElementById("link_add_form")
    form.style.display = "unset"
}
function hideAddForm(){
    let form = document.getElementById("link_add_form")
    form.style.display = "none"
}
function addLink(){
    let form = document.getElementById("link_add")
    let tip = document.getElementById("link_tip")
    let formData = JSON.stringify({
        title: form.ltitle.value,
        gravatar: form.lgravatar.value,
        description: form.ldescription.value,
        url: form.llink.value,
        email: form.lmail.value
    })
    let xhr = new XMLHttpRequest()
    xhr.open('POST', '/admin/link/add')
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

function showUpdateForm(){
    let radios = document.getElementsByName("linkId")
    let tip = document.getElementById("link_tip")
    let checkedRadio
    radios.forEach(radio => {
        if(radio.checked) {
            checkedRadio = radio;
        }
    });
    if (!checkedRadio){
        tip.innerText = "未选择友链!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let form = document.getElementById("link_update_form")
    let lid = form.querySelector('input[name="lid"]')
    lid.value = checkedRadio.value
    let ltitle = form.querySelector('input[name="ltitle"]')
    ltitle.value = checkedRadio.parentNode.querySelector('h2').innerText
    let lgravatar = form.querySelector('input[name="lgravatar"]')
    lgravatar.value = checkedRadio.parentNode.querySelector('img').src
    let ldescription = form.querySelector('input[name="ldescription"]')
    ldescription.value = checkedRadio.parentNode.querySelector('p').innerText
    let llink = form.querySelector('input[name="llink"]')
    llink.value = checkedRadio.parentNode.querySelector('a').href
    let lmail = form.querySelector('input[name="lmail"]')
    lmail.value = checkedRadio.parentNode.querySelector('span').innerText
    form.style.display = "unset"
}
function hideUpdateForm(){
    let form = document.getElementById("link_update_form")
    form.style.display = "none"
}
function updateLink(){
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
function deleteLink(){
    let radios = document.getElementsByName("linkId")
    let tip = document.getElementById("link_tip")
    let selectedId;
    radios.forEach(radio => {
        if(radio.checked) {
            selectedId = radio.value;
        }
    });
    if (!selectedId){
        tip.innerText = "未选择友链!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let jsonData = JSON.stringify({id: parseInt(selectedId)})

    let xhr= new XMLHttpRequest()
    xhr.open('DELETE', '/admin/link/delete')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(jsonData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE){
            if (xhr.status === 200){
                window.location.reload()                }
            else {
                tip.innerText = "发送失败"
                setTimeout(() => {
                    tip.innerText = ''
                }, 3000)
            }
        }
    }
}