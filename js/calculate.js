// When the document is ready
$(document).ready(function () {
    
    $("input.slider1").slider('setValue', 5).on('slide', function(me){
        $(".mMetric").val(me.value);
    });

    $("input.slider2").slider('setValue', 5).on('slide', function(me){
        $(".kgMetric").val(me.value);
    });

    $("input.slider3").slider('setValue', 5).on('slide', function(me){
        $(".ageMetric").val(me.value);
    });

    $('input:radio').click(function() {
        $("#genRef").text($(this).val());
    });

    $('#calculateBtn button').click(function(){
        debugger
        $.ajax(
            {
                type: "POST",
                url: "/calculateBMI",
                data: { 
                    BMI: 'testtest',
                    HealthRisk: 'testtesttest'
                },
                success:function(result){
                    var json = jQuery.parseJSON(result);
                    console.log(json);
                }});
    })
});