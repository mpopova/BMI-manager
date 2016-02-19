// When the document is ready
$(document).ready(function () {
    $.post("/getPersonalStat",
    function(data, status){
        if(status == 'success'){
            var json = jQuery.parseJSON(data);

            window.personalBMI = json;
        }
    });

    $.post("/getAverageBMI",
    function(data, status){
        if(status == 'success'){
            var json = jQuery.parseJSON(data);

            window.BMIAvg = json;
        }
    });

    $('.ownBtn').click(function(){
        var dataPoints = [];
        var bmiStats = window.personalBMI;
        var j = 0;
        
        if(!bmiStats.BMI){
            alert('No personal statistics, please first save your BMI !');
            return;
        }
        
        for(i = 0; i < bmiStats.BMI.length; i++){
            if(bmiStats.Date[i] && bmiStats.Date[i]==bmiStats.Date[i+1]){
                j++;
                continue;
            }
            dataPoints[i-j] = {x: new Date(bmiStats.Date[i]), y: parseInt(bmiStats.BMI[i])};
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
    $('.allBtn').click(function(){
        var BMI_Males = window.BMIAvg.BMI_Males[0],
            BMI_Females = window.BMIAvg.BMI_Females[0];
        var chart = new CanvasJS.Chart("chartContainer", {
        title:{
            text: "Male/Female Average BMI"              
        },
        axisX:{
                title: "Gender",
                gridThickness: 2
            },
            axisY: {
                title: "BMI"
            },
        data: [              
        {
            type: "column",
            dataPoints: [
                { label: "Male",  y: parseInt(BMI_Males)  },
                { label: "Female", y: parseInt(BMI_Females)   }
            ]
        }
        ]
    });
    chart.render();
    });
});