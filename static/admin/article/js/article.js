function deleteArticle(){
    let radios = document.getElementsByName("articleId")
    let tip = document.getElementById("article_tip")
    let selectedId;
    radios.forEach(radio => {
        if(radio.checked) {
            selectedId = radio.value;
        }
    });
    if (!selectedId){
        tip.innerText = "未选择文章!!!"
        setTimeout(() => {
            tip.innerText = '';
        }, 3000);
        return
    }
    let jsonData = JSON.stringify({id: parseInt(selectedId)})

    let xhr= new XMLHttpRequest()
    xhr.open('DELETE', '/admin/article/delete')
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
function findArticle() {
    let selectClass = document.getElementById('selectClass')
    let index = selectClass.selectedIndex
    let id = selectClass[index].value
    if (id === "0"){
        link("/admin/article")
    } else {
        link("/admin/article?class=" + id)
    }
}