import { useContext } from "react";
import WeatherContext, { type IWeatherContext } from "../context/WeatherContext";

export const useWeather = (): IWeatherContext => {
  const context = useContext(WeatherContext);

  if (!context) {
    throw new Error("useWeather must be used within a WeatherProvider");
  }

  return context;
};
