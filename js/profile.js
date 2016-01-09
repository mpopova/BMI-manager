// When the document is ready
$(document).ready(function () {
    $.get("/userinfo", function(data, status){
        alert("Data: " + data + "\nStatus: " + status);
    });
});