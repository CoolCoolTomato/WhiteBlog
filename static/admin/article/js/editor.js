var textEditor;

$(function() {
    textEditor = editormd("editor", {
        width           : "100%",
        height          : "85%",
        path            : "/static/editor/lib/",
        htmlDecode      : "style,script,iframe",
        tex             : true,
        emoji           : true,
        taskList        : true,
        flowChart       : true,
        sequenceDiagram : true,
        watch           : false,
        toolbarIcons : function() {
            return ["undo", "redo", "|",
                "bold", "italic", "del", "quote", "ucwords", "uppercase", "lowercase", "|",
                "h1", "h2", "h3", "h4", "h5", "h6", "|",
                "list-ul", "list-ol", "hr", "|",
                "link", "reference-link", "image", "code", "preformatted-text", "code-block", "table", "datetime", "html-entities", "pagebreak", "|",
                "watch", "preview", "clear", "search", "full", "|",
                "help"]
        },
        toolbarIconsClass : {
            full : "fa-arrows-alt"
        },
        toolbarHandlers : {
            full : function (cm, icon, cursor, selection){
                let article = document.getElementById("article")
                if (article.className !== "full"){
                    article.className = "full"
                }
                else {
                    article.className = ""
                }
            }
        }
    });
});
// 编辑器自适应
const element = document.querySelector('#article');
let width = element.offsetWidth;
new ResizeObserver(() => {
    if (width !== element.offsetWidth) {
        width = element.offsetWidth;
        let code = document.querySelector('.CodeMirror')
        let view = document.querySelector('.editormd-preview')
        let scroll = document.querySelector('.CodeMirror-scroll')
        if (code.style.display === "none"){
            if (view.style.display === "none"){
                code.style.width = '100%';
            }
            else {
                code.style.width = '50%';
            }
            view.style.width = '100%';
        }
        else {
            if (view.style.display === "none"){
                code.style.width = '100%';
            }
            else {
                code.style.width = '50%';
            }
            view.style.width = '50%';
        }
        code.style.height = '100%';
        view.style.height = '100%';
        scroll.style.height = '100%';
        let toolbar = document.querySelector('.editormd-toolbar')
        let height = toolbar.offsetHeight;
        if (toolbar.style.display === 'none'){
            toolbar.style.display = 'block'
            height = toolbar.offsetHeight;
            toolbar.style.display = 'none'
        }
        code.style.marginTop = height + 'px';
        view.style.top = height + 'px';
    }
}).observe(element);