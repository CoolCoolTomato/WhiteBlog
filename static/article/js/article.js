var testEditor;
$(function() {
    testEditor = editormd.markdownToHTML("content", {
        width: "100%",
        preview: true,
        watch: true,
        editor: false,
    });
    const content = document.getElementById('content');
    const toc = document.getElementById('toc');
    const ul = toc.querySelector('ul');
    const headings = content.querySelectorAll('h1, h2, h3, h4, h5, h6');
    if (headings.length > 0){
        const title = toc.querySelector('h2');
        title.innerText = "目录"
    }
    headings.forEach(function(heading, index) {
        const anchor = `heading-${index}`;
        heading.id = anchor;

        const listItem = document.createElement('li');
        listItem.innerHTML = `<a class="toc-link">${heading.textContent}</a>`;
        listItem.style.textIndent = `${(heading.tagName[1] - 1)}rem`;
        listItem.querySelector('.toc-link').addEventListener('click', function() {
            // 动态获取1rem的像素值
            const remInPixels = parseFloat(getComputedStyle(document.documentElement).fontSize);
            // 使用动态计算的rem值来设置偏移量（例如5rem）
            const offset = 4 * remInPixels; // 5rem 的动态像素值
            const headingTop = document.getElementById(anchor).getBoundingClientRect().top;
            const offsetPosition = headingTop + window.scrollY - offset;

            window.scrollTo({
                top: offsetPosition,
                behavior: 'smooth'
            });
        });
        ul.appendChild(listItem);
    });

});
