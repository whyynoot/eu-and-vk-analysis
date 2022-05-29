 var btn=document.querySelector('#enter-btn');
             
    btn.addEventListener('click', event=>{
       var url = '/students/' + document.getElementById('groupID-input').value;
       fetch(url, {
            method: 'GET'
        })
        .then((res) => {
            if (res.status >= 200 && res.status < 300) {
                return res;
            } else {
                let error = new Error(res.statusText);
                error.response = res;
                throw error
            }
        })
           .then(resp => resp.json()).then(function(data) {
            if (data.status!="OK"){
                let error = new Error(data.status);
                throw error;
            } else {
                if (data.statistics.length==0){
                    let error = new Error("Empty data");
                    throw error;
                }
                else {
                       var data_chart=[];
                        var labels=[];
                        var sum=0;
                        for(var i in data.statistics) {
                            data_chart.push(data.statistics[i]);
                            sum+=data.statistics[i];

                        }

                        var data_percent=[];
                        for(var i in data.statistics) {
                                data_percent.push(data.statistics[i]/sum);
                            }


                // Vertical bar chart
                var canvas=document.getElementById('bar-chart-student');
                var ctx = canvas.getContext('2d');
                ctx.clearRect(0, 0, canvas.width, canvas.height); 
                document.getElementById("result-error").innerHTML = ' ';

                    var myChart = new Chart(ctx, {
                    type: 'bar',
                    data: {
                        labels: ['Неуспевающие', 'Отличники', 'Хорошисты', 'Троечники' ],
                        datasets: [{
                                    label: 'Количество подписчиков',
                                    data: data_chart,
                                    backgroundColor: [
                                        'rgba(216, 27, 96, 0.6)',
                                        'rgba(3, 169, 244, 0.6)',
                                        'rgba(255, 152, 0, 0.6)',
                                        'rgba(29, 233, 182, 0.6)',

                                    ],
                                    borderColor: [
                                        'rgba(216, 27, 96, 1)',
                                        'rgba(3, 169, 244, 1)',
                                        'rgba(255, 152, 0, 1)',
                                        'rgba(29, 233, 182, 1)',

                                    ],
                                    borderWidth: 1
                                }]
                            },
                            options: {
                                legend: {
                                    display: false
                                },
                                title: {
                                    display: true,
                                    text: 'Количество подписчиков по категориям успеваемости студентов',
                                    position: 'top',
                                    fontSize: 16,
                                    padding: 20
                                },
                                scales: {
                                yAxes: [{
                                    ticks: {
                                        min: 0,
                                        precision: 0
                                        }
                                    }]
                                }

                            }
                        });


                        // Doughnut chartvar 
                    var canvas=document.getElementById('sector-chart-categories');
                    var ctx = canvas.getContext('2d');
                    ctx.clearRect(0, 0, canvas.width, canvas.height);
                        var myChart = new Chart(ctx, {
                            type: 'pie',
                            data: {
                                labels: ['Отличники', 'Хорошисты', 'Троечники', 'Неуспевающие'],
                                datasets: [{
                                    data: data_percent,
                                    backgroundColor: ['#e91e63', '#00e676', '#ff5722', '#1e88e5'],
                                    borderWidth: 0.5 ,
                                    borderColor: '#ddd'
                                }]
                            },
                            options: {
                                
                                title: {
                                    display: true,
                                    text: 'Процентное соотношение подписчиков по категориям успеваемости студентов',
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
                        });


                        }
                    }
                
           
                    }).catch(function(error) {
                            console.log(error); 
                            document.getElementById("result-error").innerHTML = error.message;
                        });
    })