hljs.configure({
    tabReplace: '    '
}) ;

document.addEventListener('DOMContentLoaded', (event) => {
    document.querySelectorAll('code').forEach((block) => {
        console.log(block);
        hljs.highlightBlock(block);

        // $("code").each(function(){
        //     console.log(1);
        //     $(this).html("<ol><li>" + $(this).html().replace(/\n/g,"\n</li><li>") +"\n</li></ol>");
        // });
    });
});