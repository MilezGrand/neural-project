import React from "react";
import Chart from "react-apexcharts";
import s from "./Info.module.css";
import { useNavigate, useLocation } from "react-router-dom";

const Info = () => {
  const location = useLocation();

  let personsResponseJson;

  React.useEffect(() => {
    const fetchData = async () => {
      let personsResponse = await fetch(
        `http://localhost:49812/database/emotions/get?id=${location.state.index}`
      );

      personsResponseJson = await personsResponse.json();

      setOptions({
        series: [
          {
            data: [
              personsResponseJson.Happy,
              personsResponseJson.Angry,
              personsResponseJson.Disgust,
              personsResponseJson.Fear,
              personsResponseJson.Sad,
              personsResponseJson.Surprise,
              personsResponseJson.Neutral,
            ],
          },
        ],
      });
    };

    fetchData();
  }, []);



  const [options, setOptions] = React.useState({
    series: [
      {
        data: [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0],
      },
    ],
    chart: {
      height: 350,
      type: "bar",
      toolbar: {
        show: false,
      },
      foreColor: '#FFF'
    },
    plotOptions: {
      bar: {
        borderRadius: 10,
        horizontal: true,
        dataLabels: {
          position: "bottom",
        },
      },
    },
    fill: {
      colors: ["#FF1779"],
    },
    dataLabels: {
      enabled: true,
      formatter: function (val) {
        return Math.round(val) + "%";
      },
      style: {
        fontSize: "12px",
        colors: ["white"],
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
      labels: {
        show: true,
        style: {
          fontSize: "16px",
          colors: ["lightgrey"],
        },
      },
      position: "top",
      axisBorder: {
        show: false,
      },
      axisTicks: {
        show: false,
      },
    },
    yaxis: {
      max: 100,
      labels: {
        show: true,
        offsetY: 5,
        style: {
          colors: ["#FFF"],
          fontSize: "20px",
          cssClass: 'apexcharts-yaxis-label',
        },
      },
    },
    title: {
      text: "Средний спектр эмоций человека",
      offsetX: 0,
      offsetY: 0,
      align: "center",
      style: {
        color: "#FFF",
        fontSize: "16px",
        fontFamily: "Inter",
      },
    },
    tooltip: {
      enabled: false,
    },
  });

  return (
    <div className={s.info}>

      <div className={s.chart}>
        <Chart
          options={options}
          series={options.series}
          type="bar"
          width={780}
          height={400}
        />
      </div>
    </div>
  );
};

export default Info;
