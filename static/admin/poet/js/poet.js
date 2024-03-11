function showAddForm(){
    let form = document.getElementById("poet_add_form")
    form.style.display = "unset"
}
function hideAddForm(){
    let form = document.getElementById("poet_add_form")
    form.style.display = "none"
}
function addPoet(){
    let form = document.getElementById("poet_add")
    let tip = document.getElementById("poet_tip")
    let formData = JSON.stringify({
        content: form.content.value,
        source: form.source.value,
    })
    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/poet/add')
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
    let radios = document.getElementsByName("poetId")
    let tip = document.getElementById("poet_tip")
    let checkedRadio
    radios.forEach(radio => {
        if(radio.checked) {
            checkedRadio = radio;
        }
    });
    if (!checkedRadio){
        tip.innerText = "未选择内容!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let form = document.getElementById("poet_update_form")
    let idInput = form.querySelector('input[name="pid"]')
    idInput.value = checkedRadio.value
    let contentInput = form.querySelector('input[name="content"]')
    contentInput.value = checkedRadio.parentNode.querySelector('span').innerText
    let sourceInput = form.querySelector('input[name="source"]')
    sourceInput.value = checkedRadio.parentNode.parentNode.querySelectorAll('td')[1].innerText
    form.style.display = "unset"
}
function hideUpdateForm(){
    let form = document.getElementById("poet_update_form")
    form.style.display = "none"
}
function updatePoet(){
    let form = document.getElementById("poet_update")
    let tip = document.getElementById("poet_tip")
    let formData = JSON.stringify({
        id: parseInt(form.pid.value),
        content: form.content.value,
        source: form.source.value,
    })
    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/poet/edit')
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

function deletePoet(){
    let radios = document.getElementsByName("poetId")
    let tip = document.getElementById("poet_tip")
    let selectedId;
    radios.forEach(radio => {
        if(radio.checked) {
            selectedId = radio.value;
        }
    });
    if (!selectedId){
        tip.innerText = "未选择内容!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let jsonData = JSON.stringify({id: parseInt(selectedId)})
    let xhr= new XMLHttpRequest()
    xhr.open('DELETE', '/admin/poet/delete')
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