import { createContext } from "react";

export interface WeatherData {
  temp_C: number;
  temp_F: number;
  temp_K: number;
}

export interface IWeatherContext {
  error: string | null;
  result: WeatherData | null;
  searchWeather: (cep: string) => Promise<void>;
}

const defaultContext: IWeatherContext = {
  error: null,
  result: null,
  searchWeather: async () => {}
};

const WeatherContext = createContext<IWeatherContext>(defaultContext);

export default WeatherContext;
