// When the document is ready
$(document).ready(function () {
    $.post("/getPersonalStat",
    function(data, status){
        if(status == 'success'){
            var json = jQuery.parseJSON(data);

            window.personalBMI = json;
        }
    });
    $('.ownBtn').click(function(){
        var dataPoints = [];
        var bmiStats = window.personalBMI
        
        for(i = 0; i < bmiStats.BMI.length; i++){
            dataPoints[i] = {x: new Date(bmiStats.Date[i]), y: parseInt(bmiStats.BMI[i])};
        }

        var chart = new CanvasJS.Chart("chartContainer",{
              title:{
                text: "Personal Statistic"
            },
            axisX:{
                title: "Date",
                gridThickness: 2
            },
            axisY: {
                title: "BMI"
            },
            data: [
            {        
                type: "area",
                dataPoints: dataPoints
            }]
        });

    chart.render();
    });
});