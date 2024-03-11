/*控制rem*/
let c = () =>{
    let width = document.documentElement.clientWidth;
    let height = document.documentElement.clientHeight;
    let mobile = document.getElementById("mobile")
    if (width > height){
        mobile.href = ""
        adapt(width);
    } else {
        mobile.href = "/static/mobile/mobile.css"
        adapt(height)
    }
}
window.addEventListener("load", c);
window.addEventListener("resize", c);

/*rem调整*/
function adapt(l) {
    let html = document.documentElement;
    html.style.fontSize = l / 100 + 'px'
}
