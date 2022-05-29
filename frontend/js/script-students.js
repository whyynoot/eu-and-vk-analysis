
    var choiceBtn=document.querySelector('#choice-stud');
    var url;

    choiceBtn.addEventListener('click', event=> {
        if ((event.target.name==='bad')||(event.target.name==='three')||(event.target.name==='good')||(event.target.name==='excellent')){
            switch (event.target.name){
                case 'excellent':
                    url='/interests/excellent';
                    break;
                case 'good':
                    url='/interests/good';
                    break;
                case 'three':
                    url='/interests/three';
                    break;
                case 'bad':
                    url='/interests/bad';
                    break;
            }}
        fetch(url, {method: 'GET', mode: 'no-cors'})
            .then((res) => {
                if (res.status >= 200 && res.status < 300) {
                    return res;
                } else {
                    let error = new Error(res.statusText);
                    error.response = res;
                    throw error
                }
            })
            .then(resp => resp.json())
            .then(function(data) {
                if (data.status!="OK"){
                let error = new Error(data.status);
                throw error;
            } else {
                if (data.statistics.length==0){
                    let error = new Error("Empty data");
                    throw error;
                }
                else {
                    var labels = [];
                    var data_chart = [];
                    var data_pie_chart;
                    var total=0;
                    for(var i in data.statistics) {
                        if (i=="total_students")
                            total=data.statistics[i];
                        else 
                            if (total<20){
                                if (data.statistics[i]>2){
                                    data_chart.push(data.statistics[i]);
                                    labels.push(i);
                                }
                            }
                        else  if (data.statistics[i]>total/10){
                                    data_chart.push(data.statistics[i]);
                                    labels.push(i);
                                }
                        }
                    
                    
                    /*for(var i in data.statistics) {
                        data_pie_chart.push(data.statistics[i]/total)    ;
                    }*/
                    

                    // Vertical bar chart
                    var canvas=document.getElementById('bar-chart-categories');
                    var ctx = canvas.getContext('2d');
                    ctx.clearRect(0, 0, canvas.width, canvas.height);         
                    document.getElementById("result-error").innerHTML = ' ';

                    var myChart = new Chart(ctx, {
                        type: 'bar',
                        data: {
                            labels: labels,
                            datasets: [{
                                label: 'Количество подписчиков',
                                data: data_chart,
                                backgroundColor: 'blue'
                                
                            }]
                        },
                        options: {
                            legend: {
                                display: false
                            },
                            title: {
                                display: true,
                                text: 'Количество подписчиков по категориям сообществ',
                                position: 'top',
                                fontSize: 16,
                                padding: 20
                            },
                            scales: {
                                yAxes: [{
                                    ticks: {
                                        min: 0,
                                        precision:0
                                    }
                                }]
                            }
                        }
                        
                    })
                    
              /*       // Doughnut chart
                var ctx = document.getElementById('sector-chart-categories').getContext('2d');
                var myChart = new Chart(ctx, {
                    type: 'pie',
                    data: {
                        labels: labels,
                        datasets: [{
                            data: data_pie_chart,
                            backgroundColor: ['#e91e63', '#00e676', '#ff5722', '#1e88e5'],
                            borderWidth: 0.5 ,
                            borderColor: '#ddd'
                        }]
                    },
                    options: {
                        title: {
                            display: true,
                            text: 'Процентное соотношение подписчиков по категориям сообщнств ВК',
                            position: 'top',
                            fontSize: 16,
                            fontColor: '#111',
                            padding: 20
                        },
                        legend: {
                            display: true,
                            position: 'bottom',
                            labels: {
                                boxWidth: 20,
                                fontColor: '#111',
                                padding: 15
                            }
                        },
                        tooltips: {
                            enabled: false
                        },
                        plugins: {
                            datalabels: {
                                color: '#111',
                                textAlign: 'center',
                                font: {
                                    lineHeight: 1.6
                                },
                                formatter: function(value, ctx) {
                                    return ctx.chart.data.labels[ctx.dataIndex] + '\n' + value + '%';
                                }
                            }
                        }
                    }
                });*/
         
                    }
            }
                    
            }).catch(function(error) {
                console.log(error);
                document.getElementById("result-error-stud").innerHTML = error.message;

            }
        )})
