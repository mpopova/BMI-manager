// When the document is ready
$(document).ready(function () {
    $.post("/getProfileInfo",
    function(data, status){
    	if(status == 'success'){
    		var json = jQuery.parseJSON(data),
    			userSpan = jQuery("#usertxt"),
    			ageSpan = jQuery("#agetxt"),
    			genderSpan = jQuery("#gendertxt");

    		userSpan.text(json.Name);
    		ageSpan.text(json.Age);
    		genderSpan.text(json.Gender);
    	}
    });
});