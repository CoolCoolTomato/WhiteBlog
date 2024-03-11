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
/**/
function addArticle(){
    let title = document.getElementById("title_input").value
    let content = document.getElementById("content").value
    let created_date = document.getElementById("created_date").value+":00Z"
    let updated_date = document.getElementById("updated_date").value+":00Z"
    let radios = document.getElementsByName("class")
    let selectedId;
    let image_id = document.getElementsByName('image_id')
    let images = Array.from(image_id).map(input => parseInt(input.value))
    radios.forEach(radio => {
        if(radio.checked) {
            selectedId = radio.value;
        }
    });
    if (!selectedId){
        alert("未选择分类")
        return
    }
    let formData = JSON.stringify({
        article: {
            title: title,
            content: content,
            classid: parseInt(selectedId),
            created_date: created_date,
            updated_date: updated_date,
        },
        image: images
    })
    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/article/add')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(formData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE){
            if (xhr.status === 200){
                window.location.reload()
            }
            else {
                window.alert('上传失败')
            }
        }
    }
}

function ListClass(self){
    if (self.className === "class_li hide"){
        showClassList(self)
    }
    else {
        hideClassList(self)
    }
}
function showClassList(self){
    let ul = self.children[1]
    self.className = "class_li show";
    ul.style.height = ul.children.length * 1.8 + "rem"
}
function hideClassList(self){
    let ul = self.children[1]
    self.className = "class_li hide";
    ul.style.height = "0";
}

function showOption(){
    let switch_card = document.getElementById("switch_card")
    switch_card.style.left = "0"
}
function showUpload(){
    let switch_card = document.getElementById("switch_card")
    switch_card.style.left = "-100%"
}

function upload(){
    let upload_input = document.getElementById("upload_input")
    let file = upload_input.files[0];
    if (!file){
        return
    }
    let formData = new FormData()
    formData.append('file', file)
    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/article/uploadimg')
    xhr.send(formData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                // 上传成功
                let res = JSON.parse(xhr.responseText)
                let file = document.createElement('div')
                file.className = 'file';

                let a = document.createElement('a')
                a.textContent = res['FileName']
                a.onclick = function (){
                    copy_link(res['FilePath'])
                }
                file.appendChild(a)

                let button = document.createElement('button')
                button.textContent = '删除'
                button.onclick = function (){
                    del_file(this)
                }
                file.appendChild(button)

                let input = document.createElement('input')
                input.name = 'image_id'
                input.value = res['ID']
                input.type = 'hidden'
                file.appendChild(input)

                let file_list = document.getElementById("file_list")
                file_list.appendChild(file)

            } else {
                // 上传失败
                alert("上传失败")
            }
        }
    }
}
function del_file(self){
    let id = parseInt(self.parentNode.querySelector('input').value)
    let jsonData = JSON.stringify({
        id: id,
    })
    let xhr= new XMLHttpRequest()
    xhr.open('DELETE', '/admin/article/deleteimg')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(jsonData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE){
            if (xhr.status === 200){
                self.parentNode.remove()
            }
            else {
                self.parentNode.remove()
                alert('删除失败')
            }
        }
    }
}
function copy_link(text){
    navigator.clipboard.writeText(text).then(() => {});
}