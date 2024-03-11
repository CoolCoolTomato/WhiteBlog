 /*控制rem*/
 let c = () =>{
     let width = document.documentElement.clientWidth;
     let height = document.documentElement.clientHeight;
     let mobile_css = document.getElementById("mobile_css")
     if (width > height){
         mobile_css.href = ""
         adapt(width);
     } else {
         mobile_css.href = "/static/mobile/mobileAdmin.css"
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
 function list(){
     let list = document.getElementById("list")
     if (list.className === "hide"){
         showList()
     }
     else {
         hideList()
         let subLists = document.getElementById("main_ul").children
         for (let s of subLists){
             hideSublist(s)
         }
     }
 }
 function showList() {
     let width = document.documentElement.clientWidth;
     let height = document.documentElement.clientHeight;
     let list = document.getElementById("list")
     list.className = "show"
     if (width > height){
         list.style.width = "10rem"
     } else {
         list.style.width = "20rem"
     }
     appS()
 }
 function hideList() {
     let width = document.documentElement.clientWidth;
     let height = document.documentElement.clientHeight;
     let list = document.getElementById("list")
     list.className = "hide"
     if (width > height){
         list.style.width = "4rem"
     } else {
         list.style.width = "5rem"
     }
     appL()
 }
 function subList(self){
     if (self.className === "li hide"){
         showSublist(self)
     }
     else {
         hideSublist(self)
     }
 }

 function showSublist(self){
     showList()
     let ul = self.children[1]
     self.className = "li show";
     let width = document.documentElement.clientWidth;
     let height = document.documentElement.clientHeight;
     if (width > height){
         ul.style.height = ul.children.length * 2 + "rem"
     } else {
         ul.style.height = ul.children.length * 4 + "rem"
     }
 }
 function hideSublist(self) {
     let ul = self.children[1]
     self.className = "li hide";
     ul.style.height = "0";
 }

 function appL(){
     let width = document.documentElement.clientWidth;
     let height = document.documentElement.clientHeight;
     let a = document.getElementById("app")
     if (width > height){
         a.style.width = "95%"
     }
 }
 function appS(){
     let width = document.documentElement.clientWidth;
     let height = document.documentElement.clientHeight;
     let a = document.getElementById("app")
     if (width > height){
         a.style.width = "88%"
     }
 }

 function link(href) {
     window.location.href = href;
 }
