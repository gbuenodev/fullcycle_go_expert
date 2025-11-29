import { createContext } from "react";

export interface WeatherData {
  city: string;
  temp_C: number;
  temp_F: number;
  temp_K: number;
}

export type Theme = 'day' | 'night';

export interface IWeatherContext {
  error: string | null;
  result: WeatherData | null;
  searchWeather: (cep: string) => Promise<void>;
  theme: Theme;
  toggleTheme: () => void;
}

const defaultContext: IWeatherContext = {
  error: null,
  result: null,
  searchWeather: async () => {},
  theme: 'day',
  toggleTheme: () => {}
};

const WeatherContext = createContext<IWeatherContext>(defaultContext);

export default WeatherContext;
