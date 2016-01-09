// When the document is ready
$(document).ready(function () {
    $("#loginForm").validate({
        rules: {
            "name": {
                required: true
            },  
            "password": {
                required: true
            } 
        }
    });
});

// $( ".login button" ).click(function() {
//   alert( "Handler for .click() called." );
// });
