import React from "react";
import Chart from "react-apexcharts";
import s from "./Results.module.css";
import { useNavigate, useLocation } from "react-router-dom";

const Results = () => {
    const location = useLocation();

    console.log(location)
    const [options, setOptions] = React.useState({
        series: [
          {
            name: "Уровень",
            data: [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0],
    
          },
        ],
        chart: {
          height: 350,
          type: "bar",
        },
        plotOptions: {
          bar: {
            borderRadius: 10,
            dataLabels: {
              position: "top",
            },
          },
        },
        dataLabels: {
          enabled: true,
          formatter: function (val) {
            return Math.round(val) + "%";
          },
          offsetY: -20,
          style: {
            fontSize: "12px",
            colors: ["#304758"],
          },
        },
    
        xaxis: {
          categories: [
            "Счастье",
            "Злость",
            "Отвращение",
            "Страх",
            "Грусть",
            "Удивление",
            "Нейтральный",
          ],
          position: "top",
          axisBorder: {
            show: false,
          },
          axisTicks: {
            show: false,
          },
          crosshairs: {
            fill: {
              type: "gradient",
              gradient: {
                colorFrom: "#D8E3F0",
                colorTo: "#BED1E6",
                stops: [0, 100],
                opacityFrom: 0.4,
                opacityTo: 0.5,
              },
            },
          },
          tooltip: {
            enabled: true,
          },
        },
        yaxis: {
          labels: {
            show: false,
          },
        },
        title: {
          text: "Спектр эмоций человека",
          floating: true,
          offsetY: 330,
          align: "center",
          style: {
            color: "#444",
          },
        },
      });
    
      return (
        <div className={s.results}>
          <div className={s.chart}>
            {location.state.personsResponseJson.map((o) => <div>
            
            <p>{o.emotions.person}</p>
            <Chart
            
              options={options}
              series={[
                {
                  data: [o.emotions.emotions.happy, o.emotions.emotions.angry, o.emotions.emotions.disgust, o.emotions.emotions.fear, o.emotions.emotions.sad, o.emotions.emotions.surprise, o.emotions.emotions.neutral],
          
                },
              ]}
              type="bar"
              width={650}
              height={350}
            />
            </div>)
            }
            
          </div>
        </div>
      );
    };

export default Results;