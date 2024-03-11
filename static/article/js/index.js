function drawer(){
    let catalog = document.getElementById("catalog")
    if (catalog.className === "catalog_open"){
        catalog.className = "catalog_close"
    }
    else {
        catalog.className = "catalog_open"
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

    let sub_class_li = document.querySelector(".sub_class_li")
    let h = sub_class_li.clientHeight
    ul.style.height = ul.children.length * h + 2 + "px"
}
function hideClassList(self){
    let ul = self.children[1]
    self.className = "class_li hide";
    ul.style.height = "0";
}