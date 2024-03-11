function switch2Image() {
    let s = document.getElementById("switch")
    s.style.transform = "translateX(0)"
}
function switch2File(){
    let s = document.getElementById("switch")
    s.style.transform = "translateX(-100%)"
}
function delImage(){
    let radios = document.getElementsByName("imageId")
    let tip = document.getElementById("image_tip")
    let selectedId;
    radios.forEach(radio => {
        if(radio.checked) {
            selectedId = radio.value;
        }
    });
    if (!selectedId){
        tip.innerText = "未选择图片!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let jsonData = JSON.stringify({
        id: parseInt(selectedId),
    })
    let xhr= new XMLHttpRequest()
    xhr.open('DELETE', '/admin/article/deleteimg')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(jsonData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE){
            window.location.reload()
        }
    }
}

function uploadFile(){
    let upload_input = document.getElementById("upload_input")
    let tip = document.getElementById("file_tip")
    let file = upload_input.files[0];
    if (!file){
        tip.innerText = "未选择文件!!!"
        return
    }
    let formData = new FormData()
    formData.append('file', file)
    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/file/upload')
    xhr.send(formData)
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE){
            if (xhr.status === 200) {
                window.location.reload()
            } else {
                alert("上传失败")
            }
        }
    }
}
function delFile(){
    let radios = document.getElementsByName("fileId")
    let tip = document.getElementById("file_tip")
    let selectedId;
    radios.forEach(radio => {
        if(radio.checked) {
            selectedId = radio.value;
        }
    });
    if (!selectedId){
        tip.innerText = "未选择文件!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let jsonData = JSON.stringify({
        id: parseInt(selectedId),
    })
    let xhr= new XMLHttpRequest()
    xhr.open('DELETE', '/admin/file/delete')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(jsonData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE){
            window.location.reload()
        }
    }
}
