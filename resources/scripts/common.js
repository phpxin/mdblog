$("#search-docs").click(function () {
    var keywords = $("#keywords").val() ;
    window.location.href= "/?keywords="+keywords ;
}) ;