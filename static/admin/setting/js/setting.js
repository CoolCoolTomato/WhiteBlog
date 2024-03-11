function changeConfig(){
    let user = document.querySelector('input[name="user"]').value
    let gravatar = document.querySelector('input[name="gravatar"]').value
    let url = document.querySelector('input[name="url"]').value
    let description = document.querySelector('input[name="description"]').value
    let username = document.querySelector('input[name="username"]').value
    let password = document.querySelector('input[name="password"]').value
    let path = document.querySelector('input[name="path"]').value
    let title = document.querySelector('input[name="title"]').value
    let icon = document.querySelector('input[name="icon"]').value
    let bili = document.querySelector('input[name="bili"]').value
    let github = document.querySelector('input[name="github"]').value
    let twitter = document.querySelector('input[name="twitter"]').value
    let mail = document.querySelector('input[name="mail"]').value
    let formData = JSON.stringify({
        authConfig: {
            username: username,
            password: password
        },
        user: {
            username: user,
            gravatar: gravatar,
            url: url,
            description: description
        },
        site: {
            path: path,
            title: title,
            icon: icon,
            bili: bili,
            github: github,
            twitter: twitter,
            mail: mail
        }
    })
    let xhr= new XMLHttpRequest()
    xhr.open('POST', '/admin/setting')
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(formData)
    xhr.onreadystatechange = function (){
        if (xhr.readyState === XMLHttpRequest.DONE){
            if (xhr.status === 200){
                window.location.reload()
            }
            else {
                setTimeout(() => {
                }, 3000)
            }
        }
    }
}