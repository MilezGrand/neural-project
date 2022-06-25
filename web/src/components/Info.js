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

  const navigate = useNavigate();

  const handleClick = () => {
    navigate("/lib");
  };

  const [options, setOptions] = React.useState({
    series: [
      {
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
      offsetY: 20,
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
      position: "top",
      style: {
        colors: ["white"],
      },

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
        enabled: false,
      },
    },
    yaxis: {
      labels: {
        show: false,
      },
    },
    title: {
      text: "Средний спектр эмоций человека",
      floating: true,
      offsetY: 380,
      align: "center",
      style: {
        color: "#FFF",
      },
    },
  });

  return (
    <div className={s.info}>
      <span onClick={handleClick} className={s.back}>
        Назад
      </span>
      <div className={s.chart}>
        <Chart
          options={options}
          series={options.series}
          type="bar"
          width={790}
          height={400}
        />
      </div>
    </div>
  );
};

export default Info;
