// When the document is ready
$(document).ready(function () {
    document.mMetric = 0;
    document.kgMetric = 0;
    document.ageMetric = 0;

    $("input.slider1").slider('setValue', 5).on('slide', function(me){
        $(".mMetric").val(me.value);
        document.mMetric = me.value/100;
        var bmiIndex = document.kgMetric/(document.mMetric*document.mMetric);
        $(".bmi").text(parseFloat(bmiIndex).toFixed(2));
        if(bmiIndex<18.5){
            $(".risk").text("Underweight");
        }
        if(bmiIndex<18.5){
            $(".risk").text("Underweight");
        }
        if(bmiIndex>18.5 && bmiIndex<25){
            $(".risk").text("Normal weight");
        }
        if(bmiIndex>25){
            $(".risk").text("Overweight");
        }
    });

    $("input.slider2").slider('setValue', 5).on('slide', function(me){
        $(".kgMetric").val(me.value);
        document.kgMetric = me.value;
        var bmiIndex = document.kgMetric/(document.mMetric*document.mMetric);
        $(".bmi").text(parseFloat(bmiIndex).toFixed(2));
        if(bmiIndex < 18.5){
            $(".risk").text("Underweight");
        }
        if(bmiIndex > 18.5 && bmiIndex < 25){
            $(".risk").text("Normal weight");
        }
        if(bmiIndex > 25){
            $(".risk").text("Overweight");
        }
    });

    $('#calculateBtn button').click(function(){
        var bmiIndex = document.kgMetric/(document.mMetric*document.mMetric);
        if(bmiIndex && bmiIndex > 0){
            $.ajax({
                type: "POST",
                url: "/calculateBMI",
                data: { 
                    BMI: bmiIndex
                },
                success:function(result){
                }
            }); 
        }
        else{
            alert('You must insert kilos and height!');
        }
    });
});