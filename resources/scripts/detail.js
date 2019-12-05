hljs.configure({
    tabReplace: '    '
}) ;

document.addEventListener('DOMContentLoaded', (event) => {
    document.querySelectorAll('code').forEach((block) => {
        console.log(block);
        hljs.highlightBlock(block);
    });
});