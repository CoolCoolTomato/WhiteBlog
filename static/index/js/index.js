function choose(key){
    let items_warp = document.getElementById("items_warp")
    let items = ["home", "code", "contact", "poet"]
    switch (key){
        case items[0]:
            items_warp.style.transform = "translateY(0)";
            break
        case items[1]:
            items_warp.style.transform = "translateY(-100%)";
            break
        case items[2]:
            items_warp.style.transform = "translateY(-200%)";
            break
        case items[3]:
            items_warp.style.transform = "translateY(-300%)";
            break
    }
}
